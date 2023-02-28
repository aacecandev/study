package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"regexp"
	"time"

	"cloud.google.com/go/compute/metadata"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/gorilla/mux"
	"github.com/mtslzr/pokeapi-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/yfuruyama/crzerolog"
	"go.opencensus.io/plugin/ochttp"
	"google.golang.org/api/idtoken"

	"net/http"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	mux := mux.NewRouter()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/pokemon/{id:[0-9]+}", pokemonHandler)

	rootLogger := zerolog.New(os.Stdout)
	middleware := crzerolog.InjectLogger(&rootLogger)
	handler := middleware(mux)

	httpHandler := &ochttp.Handler{
		Propagation: &propagation.HTTPFormat{},
		Handler:     handler,
	}
	if err := http.ListenAndServe(":8080", httpHandler); err != nil {
		log.Fatal().Err(err).Msg("Can’t start service")
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	logger.Info().Msg("Serving random pokemon")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ps, err := pokeapi.Resource("pokemon", 1, 500)
	if err != nil {
		log.Ctx(r.Context()).Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	randomSelection := rand.Perm(len(ps.Results))
	re := regexp.MustCompile("/pokemon/([0-9]+)*")

	ids := []string{}
	for _, v := range randomSelection[:3] {
		match := re.FindStringSubmatch(ps.Results[v].URL)
		ids = append(ids, match[1])
	}
	defaultBaseURL := "http://localhost:8080"
	client := NewClient(r, defaultBaseURL)

	var result []Pokemon
	for _, id := range ids {
		logger.Info().Str("pokemonId", id).Msg("Requesting pokemon")
		nextR, err := client.NewRequest()
		if err != nil {
			log.Ctx(r.Context()).Error().Err(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nextR.URL.Path = "/pokemon/" + id
		response, err := client.client.Do(nextR)
		if err != nil {
			logger.Err(err).Msg("Failed to request pokemon")
			continue
		}

		var p Pokemon
		err = json.NewDecoder(response.Body).Decode(&p)

		if err != nil {
			logger.Err(err).Msg("Failed to parse pokemon json")
			continue
		}
		result = append(result, p)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(result)
}
func pokemonHandler(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	vars := mux.Vars(r)
	id := vars["id"]
	p, err := pokeapi.Pokemon(id)
	logger.Info().
		Int("weight", p.Weight).
		Int("height", p.Height).
		Str("name", p.Species.Name).
		Str("pokemonId", id).
		Msg("Spawning pokemon")
	if err != nil {
		log.Ctx(r.Context()).Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(p)
}

type httpClient struct {
	defaultBaseURL string
	baseRequest    *http.Request
	client         *http.Client
}

func NewClient(r *http.Request, defaultBaseURL string) httpClient {
	logger := log.Ctx(r.Context())
	var client *http.Client
	if metadata.OnGCE() {
		var err error
		client, err = idtoken.NewClient(context.Background(), r.Host)
		if err != nil {
			logger.Err(err).Msg("failed to create idtoken client")
		}
	} else {
		client = http.DefaultClient
	}
	return httpClient{
		defaultBaseURL: defaultBaseURL,
		baseRequest:    r,
		client:         client,
	}
}

func (c *httpClient) NewRequest() (*http.Request, error) {
	baseURL := c.defaultBaseURL
	if metadata.OnGCE() {
		baseURL = "https://" + c.baseRequest.Host
	}
	req, err := http.NewRequestWithContext(c.baseRequest.Context(), "GET", baseURL, nil)
	return req, err
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Species        struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Weight int `json:"weight"`
}

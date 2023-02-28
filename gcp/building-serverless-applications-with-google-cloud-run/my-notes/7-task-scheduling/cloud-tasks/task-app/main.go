package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"net/http"

	tasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/compute/metadata"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/yfuruyama/crzerolog"
	"go.opencensus.io/plugin/ochttp"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Ctx(r.Context()).Info().Msg("Task Scheduled")
		queue := os.Getenv("QUEUE")
		workerURL := os.Getenv("WORKER_URL")
		sendTask(workerURL, queue)
		fmt.Fprint(w, "Hello")
	})

	rootLogger := zerolog.New(os.Stdout)
	middleware := crzerolog.InjectLogger(&rootLogger)
	handler := middleware(mux)

	httpHandler := &ochttp.Handler{
		Propagation: &propagation.HTTPFormat{},
		Handler:     handler,
	}
	if err := http.ListenAndServe(":8080", httpHandler); err != nil {
		log.Fatal().Msg("Can't start service")
	}
}

func sendTask(URL string, queueName string) {
	path := ""
	// tasks "cloud.google.com/go/cloudtasks/apiv2"
	client, err := tasks.NewClient(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create client")
	}

	serviceAccount, err := metadata.Email("default")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create client")
	}

	// taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	req := &taskspb.CreateTaskRequest{
		Parent: queueName,
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					Url:        URL + "/" + path,       // Destination
					HttpMethod: taskspb.HttpMethod_GET, // or POST
					// Add ID Token to Request
					AuthorizationHeader: &taskspb.HttpRequest_OidcToken{
						OidcToken: &taskspb.OidcToken{
							ServiceAccountEmail: serviceAccount,
							Audience:            URL, // Scope token to URL
						},
					},
				},
			},
			// Match Cloud Run timeout:
			DispatchDeadline: durationpb.New(15 * time.Minute),
		},
	}
	_, err = client.CreateTask(context.Background(), req)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create task")
	}

}

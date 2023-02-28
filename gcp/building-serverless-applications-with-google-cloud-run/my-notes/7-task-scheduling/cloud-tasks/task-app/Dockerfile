# A named stage: "build"
FROM golang:1.15 AS build
 
WORKDIR /src
COPY go.* /src/
RUN go mod download
 
COPY . /src
RUN go build -o /go/bin/app

# A new stage: "run"
FROM gcr.io/distroless/base-debian10:nonroot AS run

# Copy the binary from stage build
COPY --from=build /go/bin/app /
 
ENTRYPOINT ["/app"]
FROM golang:1.22-alpine as builder

# Set the Go environment
ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOARCH amd64

# Install dependencies
RUN apk --update --no-cache add ca-certificates make bash

WORKDIR /go-expert-lab-o11y-otel/internal/

COPY ./pkg ./pkg

WORKDIR /go-expert-lab-o11y-otel/internal/services/weatherservice

# Build Go binary
COPY ./services/weatherservice/go.mod ./
COPY ./services/weatherservice/go.sum ./
RUN go mod download
COPY ./services/weatherservice .
RUN go mod tidy

# Set the App environment
ARG app_env
ENV APP_ENV $app_env

# Run the app
ARG project
ENV PROJECT $project

# CMD go run ${PROJECT}
CMD go run ${PROJECT}
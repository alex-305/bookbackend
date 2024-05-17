# syntax=docker/dockerfile:1

FROM golang:1.22.2

WORKDIR /app

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

COPY go.mod go.sum ./

RUN go mod download

COPY internal/ ./internal
COPY cmd/bookdom/ ./cmd/bookdom
COPY .env ./

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build ./cmd/bookdom/main.go" -command="./main"

# RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/bookdom/main.go

# CMD [ "./main" ]
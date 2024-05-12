# syntax=docker/dockerfile:1

FROM golang:1.22.2

ENV PROJECT_DIR=/src \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /src

COPY go.mod go.sum ./

COPY . ./
COPY .env ./

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -build="go build -o /bookbackend" -command="/bookbackend"

# RUN CGO_ENABLED=0 GOOS=linux go build -o /bookbackend

# CMD ["/bookbackend"]
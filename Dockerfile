# syntax=docker/dockerfile:1

FROM golang:1.22.2

WORKDIR /src

COPY go.mod go.sum ./

COPY . ./
COPY .env ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bookbackend

CMD ["/bookbackend"]
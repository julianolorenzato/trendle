# syntax=docker/dockerfile:1
FROM golang:1.20-alpine as base

# development
FROM base as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

# production
FROM base as prod

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./  ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server

EXPOSE 8080

# default command (can be overriden)
CMD [ "./server" ]
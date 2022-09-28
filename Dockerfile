# syntax=docker/dockerfile:1

## Build
FROM golang:1.18 AS build

WORKDIR /usr/src/app

# go.sum
COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o site

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /usr/src/app/site /site
COPY static /static

EXPOSE 8080

USER nonroot:nonroot

CMD ["/site"]

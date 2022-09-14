# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.* ./
RUN go mod download

#COPY static /static

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o site


## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build app/site /site
COPY static /static

EXPOSE 8080

USER nonroot:nonroot

CMD ["/site"]

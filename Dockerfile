FROM golang:1.15 AS build

WORKDIR /app

COPY . .
# COPY go.mod go.sum ./

RUN CGO_ENABLED=0 go build -o bin/banner-rotator ./cmd/rotator/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/banner-rotator .
COPY ./configs/config.toml ./configs/config.toml
RUN mkdir logs
CMD [ "./banner-rotator" ]

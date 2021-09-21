FROM golang:1.16 as build

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/price-generator /app

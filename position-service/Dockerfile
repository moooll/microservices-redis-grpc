FROM golang:1.16 as build

WORKDIR /app 

COPY . /app
RUN go mod tidy
RUN go build

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/position-service /app

FROM golang:1.16 as build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build

FROM alpine
COPY --from=build /app/console /app
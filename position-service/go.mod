module github.com/moooll/microservices-redis-grpc/position-service

go 1.16

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/moooll/microservices-redis-grpc/price-generator v0.0.0-20210921170320-4875628925d9 // indirect
	github.com/moooll/microservices-redis-grpc/price-service v0.0.0-20210921170320-4875628925d9 // indirect
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

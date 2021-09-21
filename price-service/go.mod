module github.com/moooll/microservices-redis-grpc/price-service

go 1.16

require (
	github.com/caarlos0/env/v6 v6.7.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/moooll/microservices-redis-grpc/price-generator v0.0.0-20210921170320-4875628925d9
	github.com/onsi/gomega v1.16.0 // indirect
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

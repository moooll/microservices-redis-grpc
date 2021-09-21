module github.com/moooll/microservices-redis-grpc/console

go 1.16

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/moooll/microservices-redis-grpc/position-service v0.0.0-20210920161338-43a830f9f529
	github.com/moooll/microservices-redis-grpc/price-generator v0.0.0-20210921134556-e445799406a3
	github.com/moooll/microservices-redis-grpc/price-service v0.0.0-20210921170320-4875628925d9
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.40.0
)

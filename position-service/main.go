package main

import (
	"context"
	"net"
	"sync"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4"
	rpc "github.com/moooll/microservices-redis-grpc/position-service/communication/client"
	rpcserver "github.com/moooll/microservices-redis-grpc/position-service/communication/server"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/config"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
	pbserver "github.com/moooll/microservices-redis-grpc/position-service/protocol"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
	"github.com/pquerna/ffjson/ffjson"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("error parsing config: ", err.Error())
	}

	conn, client, er := grpcClientConnect(cfg.GRPCClientAddr)
	if er != nil {
		log.Error("error connecting to grpc as a client: ", er.Error()) 
	}

	defer func() {
		if e := conn.Close(); e != nil {
			log.Error("error parsing config: ", e.Error())
		}
	}()

	c := make(chan models.Price)
	latestPrice := make(map[string]models.Price)

	dbConn, e := postgresConnect(context.Background(), cfg.PgURI)
	if e != nil {
		log.Error("error connecting to postgres: ", e.Error())
	}

	cl := *rpc.NewGetPriceService(c, latestPrice, &sync.Mutex{})
	server := rpcserver.NewProfitAndLoss(cl, dbConn)

	go func() {
		if er := launchGRPCServer(cfg.GRPCServerPort, server); er != nil {
			log.Error("error launching gRPC server: ", er.Error())
		}
	}()

	go func() {
		er := cl.GetPrice(context.Background(), client)
		if er != nil {
			log.Error("error getting price: ", er.Error())
		}
	}()

	go func() {
		eer := postgresListen(context.Background(), dbConn, "notification", cfg.ServerID, server)
	if eer != nil {
		log.Error("error listening to postgres: ", eer.Error())
	}
	}()
	var wait chan struct{}
	<-wait
}

func grpcClientConnect(grpcAddr string) (*grpc.ClientConn, pb.PriceServiceClient, error) {
	conn, er := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if er != nil {
		return nil, nil, er
	}

	return conn, pb.NewPriceServiceClient(conn), nil
}

func launchGRPCServer(port string, server rpcserver.ProfitAndLoss) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pbserver.RegisterProfitAndLossServer(s, &server)
	if er := s.Serve(lis); er != nil {
		return er
	}

	return nil
}

func postgresConnect(ctx context.Context, addr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func postgresListen(ctx context.Context, conn *pgx.Conn, channame string, serverID int, server rpcserver.ProfitAndLoss) error {
	_, err := conn.Exec(ctx, "listen notification")
	if err != nil {
		return err
	}

	notif, er := conn.WaitForNotification(ctx)
	if er != nil {
		return er
	}

	p := models.Position{}
	e := ffjson.Unmarshal([]byte(notif.Payload), &p)
	if e != nil {
		return e
	}

	if p.ServerID != serverID {
		var positionStatus pbserver.ProfitAndLossRequest_Position
		if !p.Open {
			positionStatus = pbserver.ProfitAndLossRequest_Position(pbserver.ProfitAndLossRequest_Position_value["CLOSE"])
		} else {
			positionStatus = pbserver.ProfitAndLossRequest_Position(pbserver.ProfitAndLossRequest_Position_value["OPEN"])
		}

		server.GetProfitAndLoss(ctx, &pbserver.ProfitAndLossRequest{
			Id:          p.ID.String(),
			CompanyName: p.CompanyName,
			BuyPrice:    p.BuyPrice,
			SellPrice:   p.SellPrice,
			Position:    positionStatus,
		})
	}

	return nil
}

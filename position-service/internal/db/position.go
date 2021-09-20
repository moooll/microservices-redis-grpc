package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
)

func AddPosition(ctx context.Context, conn *pgx.Conn, position models.Position) error {
	r, err := conn.Query(ctx, `insert into positions(
		id uuid primary key,
		server_id smallint,
		company_name varchar,
		open boolean,
		buy_price real,
		sell_price real,
		pnl real)
		values $1, $2, $3, $4, $5, $6`, 
		position.ID, 
		position.ServerID,
		position.CompanyName,
		position.Open, 
		position.BuyPrice,
		position.SellPrice,
		position.ProfitAndLoss)
	if err != nil {
		return err
	}

	defer r.Close()

	return nil
}

func UpdPosition(ctx context.Context, conn *pgx.Conn, serverID int, sellPrice, pnl float32) error {
	r, err := conn.Query(ctx, "update positions set server_id = $1 , open = false, sell_price = $2, pnl = $3", serverID, sellPrice, pnl)
	if err != nil  {
		return err
	}

	defer r.Close()

	return nil
}
package db

import (
	"context"
	"errors"
	"time"
	"wb_task1/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Conn struct {
	pgConn *pgxpool.Pool
}

func New(path string) (*Conn, error) {
	pgConn, err := pgxpool.New(context.Background(), path)
	if err != nil {
		return nil, errors.New("Can't connect to db")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pgConn.Ping(ctx); err != nil {
		return nil, errors.New("Connection not responding to ping")
	}
	return &Conn{pgConn}, nil
}

func (c *Conn) Write(data entity.Order) error {
	ctx, cancFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancFunc()

	d := &data.Delivery
	var personId int64
	err := c.pgConn.QueryRow(ctx, insertPerson, d.Name, d.Phone, d.Zip, d.City,
		d.Address, d.Region, d.Email).Scan(&personId)
	if err != nil {
		return err
	}
	p := &data.Payment
	var transact string
	err = c.pgConn.QueryRow(ctx, insertPayment, p.Transaction, p.RequestID, p.Currency,
		p.Provider, p.Amount, p.PaymentDt, p.Bank, p.DeliveryCost, p.Total, p.CustomFee).Scan(&transact)
	if err != nil {
		return err
	}
	var orderId string
	err = c.pgConn.QueryRow(ctx, insertOrder, data.ID, data.TrackNum, data.Entry, data.Locale, data.Sign,
		data.CustometID, data.DeliveryService, data.ShardKey, data.SmID, data.DateCreated,
		data.OofShard, personId, transact).Scan(&orderId)
	if err != nil {
		return err
	}
	var itemId int64
	for _, item := range data.Items {
		c.pgConn.QueryRow(ctx, insertItem, item.ID, item.TrackNum, item.Price, item.RID, item.Name,
			item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status).Scan(&itemId)
		if err != nil {
			return err
		}
		_, err = c.pgConn.Exec(ctx, insertItemOrderRel, orderId, itemId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Conn) LoadAll() (*map[string]entity.Order, error) {
	ctx, cancFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancFunc()

	orders := make(map[string]entity.Order)

	rows, err := c.pgConn.Query(ctx, getAllOrders)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var order entity.Order
		err = rows.Scan(&order)
		orders[order.ID] = order
	}

	return &orders, nil
}

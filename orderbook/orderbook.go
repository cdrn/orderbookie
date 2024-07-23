package orderbook

import (
	"container/heap"
	"database/sql"
)

type OrderBook struct {
	Bids *OrderQueue
	Asks *OrderQueue
	DB   *sql.DB
}

func NewOrderBook(db *sql.DB) *OrderBook {
	return &OrderBook{
		Bids: NewOrderQueue(),
		Asks: NewOrderQueue(),
		DB:   db,
	}
}

func (ob *OrderBook) AddOrder(order *Order) error {
	_, err := ob.DB.Exec("INSERT INTO orders (price, quantity, is_buy) VALUES ($1, $2, $3)", order.Price, order.Quantity, order.Side)
	if err != nil {
		return err
	}
	if order.Side == "buy" {
		heap.Push(ob.Bids, order)
		// Sell side
	} else {
		heap.Push(ob.Asks, order)
	}
	return nil
}

func (ob *OrderBook) LoadOrders() error {
	rows, err := ob.DB.Query("SELECT id, price, quantity, is_buy FROM orders ORDER BY id")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Price, &order.Quantity, &order.Side); err != nil {
			return err
		}
		if order.Side == "buy" {
			heap.Push(ob.Bids, &order)
			// Sell side
		} else {
			heap.Push(ob.Asks, &order)
		}
	}
	return rows.Err()
}

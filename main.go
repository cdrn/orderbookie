package main

import (
	"log"
	"orderbookie/db"
	"orderbookie/orderbook"
)

func main() {
	connStr := "user=youruser dbname=orderbookie sslmode=disable"
	database := db.InitDB(connStr)
	defer database.Close()

	ob := orderbook.NewOrderBook(database)
	err := ob.LoadOrders()
	if err != nil {
		log.Fatal(err)
	}

	order := &orderbook.Order{Price: 100.0, Quantity: 10, Side: "buy"}
	err = ob.AddOrder(order)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Order book loaded and updated")
}

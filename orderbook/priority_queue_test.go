package orderbook

import (
	"container/heap"
	"testing"
)

func TestOrderQueue(t *testing.T) {
	pq := NewOrderQueue()

	// Add 5 orders
	orders := []*Order{
		{Price: 5},
		{Price: 3},
		{Price: 8},
		{Price: 1},
		{Price: 7},
	}

	for _, order := range orders {
		heap.Push(pq, order)
	}

	// Check the order of elements in the priority queue
	expectedOrder := []int{1, 3, 5, 7, 8}
	for i, expectedPrice := range expectedOrder {
		order := heap.Pop(pq).(*Order)
		if order.Price != float64(expectedPrice) {
			t.Errorf("Expected order at index %d to have price %f, but got %f", i, float64(expectedPrice), order.Price)
		}
	}
}

func TestOrderQueueWithAdditionalOrders(t *testing.T) {
	pq := NewOrderQueue()

	// Add 5 orders
	orders := []*Order{
		{Price: 10},
		{Price: 2},
		{Price: 6},
		{Price: 4},
		{Price: 9},
	}

	for _, order := range orders {
		heap.Push(pq, order)
	}

	// Check the order of elements in the priority queue
	expectedOrder := []int{2, 4, 6, 9, 10}
	for i, expectedPrice := range expectedOrder {
		order := heap.Pop(pq).(*Order)
		if order.Price != float64(expectedPrice) {
			t.Errorf("Expected order at index %d to have price %f, but got %f", i, float64(expectedPrice), order.Price)
		}
	}
}

func TestOrderQueueMixedOrders(t *testing.T) {
	pq := NewOrderQueue()

	// Add 5 orders
	orders := []*Order{
		{Price: 15},
		{Price: 3},
		{Price: 12},
		{Price: 1},
		{Price: 5},
	}

	for _, order := range orders {
		heap.Push(pq, order)
	}

	// Check the order of elements in the priority queue
	expectedOrder := []int{1, 3, 5, 12, 15}
	for i, expectedPrice := range expectedOrder {
		order := heap.Pop(pq).(*Order)
		if order.Price != float64(expectedPrice) {
			t.Errorf("Expected order at index %d to have price %f, but got %f", i, float64(expectedPrice), order.Price)
			return
		}
	}
}

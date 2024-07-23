package orderbook

import (
	"container/heap"
	"testing"
)

func TestOrderQueue(t *testing.T) {
	// Create a new order queue
	pq := NewOrderQueue()

	// Define some test orders
	order1 := &Order{Price: 100}
	order2 := &Order{Price: 200}
	order3 := &Order{Price: 150}

	// Push orders onto the queue
	heap.Push(pq, order1)
	heap.Push(pq, order2)
	heap.Push(pq, order3)

	// Test the length of the queue
	if pq.Len() != 3 {
		t.Errorf("expected length 3, got %d", pq.Len())
	}

	// Test the order of elements (highest price first)
	expectedOrder := []*Order{order2, order3, order1}
	for i, expected := range expectedOrder {
		if (*pq)[i] != expected {
			t.Errorf("expected order %v at index %d, got %v", expected, i, (*pq)[i])
		}
	}

	// Pop elements and test the order
	poppedOrder := heap.Pop(pq).(*Order)
	if poppedOrder != order2 {
		t.Errorf("expected order2, got %v", poppedOrder)
	}

	poppedOrder = heap.Pop(pq).(*Order)
	if poppedOrder != order3 {
		t.Errorf("expected order3, got %v", poppedOrder)
	}

	poppedOrder = heap.Pop(pq).(*Order)
	if poppedOrder != order1 {
		t.Errorf("expected order1, got %v", poppedOrder)
	}

	// Test the length of the queue after popping all elements
	if pq.Len() != 0 {
		t.Errorf("expected length 0, got %d", pq.Len())
	}
}

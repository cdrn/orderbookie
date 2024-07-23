package orderbook

import "container/heap"

type OrderQueue []*Order

func (pq OrderQueue) Len() int { return len(pq) }

func (pq OrderQueue) Less(i, j int) bool {
	return pq[i].Price < pq[j].Price
}

func (pq OrderQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *OrderQueue) Push(x interface{}) {
	order := x.(*Order)
	*pq = append(*pq, order)
}

func (pq *OrderQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	order := old[n-1]
	*pq = old[0 : n-1]
	return order
}

func NewOrderQueue() *OrderQueue {
	pq := &OrderQueue{}
	heap.Init(pq)
	return pq
}

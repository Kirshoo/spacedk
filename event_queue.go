package spacedk

import (
	"time"
	"sync"
	"container/heap"
)

type EventElement struct {
	value *Event
	TransmitAt time.Time
	index int
}

type EventPriorityQueue struct {
	mutex sync.RWMutex

	queue []*EventElement
}

func (pq *EventPriorityQueue) Len() int { 
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	return len(pq.queue) 
}

func (pq *EventPriorityQueue) Less(i, j int) bool {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	return pq.queue[i].TransmitAt.Before(pq.queue[i].TransmitAt)
}

func (pq *EventPriorityQueue) Swap(i, j int) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
	pq.queue[i].index = i
	pq.queue[j].index = j
}

func (pq *EventPriorityQueue) Push(x any) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	elem := x.(*EventElement)
	elem.index = len(pq.queue)
	pq.queue = append(pq.queue, elem)
}

func (pq *EventPriorityQueue) Pop() any {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	old := pq.queue
	n := len(old)
	elem := pq.queue[n - 1]
	old[n - 1] = nil
	elem.index = -1
	pq.queue = old[:n -1]

	return elem
}

func (pq *EventPriorityQueue) Update(elem *EventElement) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	updateIndex := -1
	for i, element := range pq.queue {
		if !element.value.IsEqualTo(elem.value) {
			continue	
		}

		element.TransmitAt = elem.TransmitAt
		updateIndex = i
		break
	}

	if updateIndex == -1 {
		return
	}

	heap.Fix(pq, updateIndex)
}

func (pq *EventPriorityQueue) Delete(elem *EventElement) *EventElement {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	deleteIndex := -1
	for i, element := range pq.queue {
		if !element.value.IsEqualTo(elem.value) {
			continue	
		}

		deleteIndex = i
		break
	}

	if deleteIndex == -1 {
		return nil
	}

	return heap.Remove(pq, deleteIndex).(*EventElement)
}

func (pq *EventPriorityQueue) Peak() *EventElement {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	if len(pq.queue) == 0 {
		return nil
	}

	return pq.queue[0]
}

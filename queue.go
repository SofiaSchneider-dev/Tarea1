package main

//hay varias formas de implementar un queue en Go
// con slices, doubly linked list, channels.
//en este caso implementaré un queue con slices

type MyQueue struct {
	elements []int
}

func NewQueue() *MyQueue {
	return &MyQueue{elements: []int{}}
}

func (queue *MyQueue) isEmpty() bool {
	if len(queue.elements) == 0 {
		return true
	}
	return false
}

func (queue *MyQueue) enqueue(item int) {
	queue.elements = append(queue.elements, item)
}

func (queue *MyQueue) dequeue() (int, bool) {
	if queue.isEmpty() {
		return 0, false
	}
	item := queue.elements[0]
	queue.elements = queue.elements[1:]
	return item, true
}

func (queue *MyQueue) front() (int, bool) {
	if queue.isEmpty() {
		return 0, false
	}
	return queue.elements[0], true
}

func (queue *MyQueue) back() (int, bool) {
	if queue.isEmpty() {
		return 0, false
	}
	return queue.elements[len(queue.elements)-1], true
}

func (queue *MyQueue) size() int {
	return len(queue.elements)
}

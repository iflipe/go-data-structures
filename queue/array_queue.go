package queue

import "errors"

type ArrayQueue struct {
	v     []int
	front int
	rear  int
}

func (queue *ArrayQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New("empty queue")
	}
	return queue.v[queue.front], nil
}

func (queue *ArrayQueue) Enqueue(val int) {
	if queue.Full() {
		queue.DoubleSize()
	}
	if queue.IsEmpty() {
		queue.front = 0
	}
	queue.rear = (queue.rear + 1) % len(queue.v)
	queue.v[queue.rear] = val
}

func (queue *ArrayQueue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("the queue is empty")
	}
	val := queue.v[queue.front]
	if queue.Size() == 1 {
		queue.front = -1
		queue.rear = -1
	} else {
		queue.front = (queue.front + 1) % len(queue.v)
	}
	return val, nil
}

func (queue *ArrayQueue) Full() bool {
	return queue.front == (queue.rear+1)%len(queue.v)
}

func (queue *ArrayQueue) DoubleSize() {
	doubleQueue := make([]int, len(queue.v)*2)
	for i := 0; i < len(queue.v); i++ {
		x, _ := queue.Dequeue()
		doubleQueue[i] = x
	}
	queue.front = 0
	queue.rear = len(queue.v) - 1
	queue.v = doubleQueue
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *ArrayQueue) Init(size int) {
	queue.v = make([]int, size)
	queue.front = -1
	queue.rear = -1
}

func (queue *ArrayQueue) Size() int {
	if queue.front == -1 && queue.rear == -1 {
		return 0
	}
	size := queue.rear - queue.front + 1
	if queue.front <= queue.rear {
		return size
	}
	return len(queue.v) + size
}

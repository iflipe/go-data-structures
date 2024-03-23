package queue

import (
	"testing"
)

var size int

//var queues [2]IQueue

var queues [1]IQueue

func createQueues(size int) {
	array_queue := &ArrayQueue{}
	//linkedlist_queue := &LinkedListQueue{}
	(*array_queue).Init(size)
	//queues = [2]IQueue{array_queue, linkedlist_queue}
	queues = [1]IQueue{array_queue}
}

func deleteQueues() {
	//in this case, createQueues alone solves the problem
	//however, I let the template here to be used in other tests
	queues[0] = nil
	//queues[1] = nil
}

func setupTest() func() {
	//before each test
	size = 10
	createQueues(size)

	//after each test
	return func() {
		deleteQueues()
	}
}

func TestEnqueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		for i := 0; i < 2*size; i++ {
			queue.Enqueue(i)
			if queue.Size() != i+1 {
				t.Errorf("%T size is %d, but we expected it to be %d", queue, queue.Size(), i+1)
			}
		}
	}
}

func TestDequeue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		for i := 0; i < size; i++ {
			queue.Enqueue(i)
		}
		for i := 0; i < size; i++ {
			val, err := queue.Dequeue()
			//front ==> 0, 1, 2, 3, ..., 9 ==> rear
			if val != i {
				t.Errorf("Value dequeued from %T is %d, but we expected it to be %d", queue, val, i)
			}
			if err != nil {
				t.Errorf(err.Error())
			}
			if queue.Size() != size-i-1 {
				t.Errorf("%T size is %d, but we expected it to be %d", queue, queue.Size(), size-i-1)
			}
		}
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		_, err := queue.Dequeue()
		if err == nil {
			t.Errorf("Dequeued an empty queue")
		}
	}
}

func TestFront(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		for i := 0; i < size; i++ {
			queue.Enqueue(i)
			val, err := queue.Front()
			if val != 0 {
				t.Errorf("Value in front of from %T is %d, but we expected it to be %d", queue, val, 0)
			}
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	}
}

func TestFrontEmptyQueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		_, err := queue.Front()
		if err == nil {
			t.Errorf("Empty %T should return error when asked the front element", queue)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		queue.Enqueue(0)
		empty := queue.IsEmpty()
		if empty {
			t.Errorf("Non Empty %T should return false for IsEmpty operation", queue)
		}
	}
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		empty := queue.IsEmpty()
		if !empty {
			t.Errorf("Empty %T should return true for IsEmpty operation", queue)
		}
	}
}

func TestSize(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		queue.Enqueue(0)
		size := queue.Size()
		if size != 1 {
			t.Errorf("%T size is %d, but we expected it to be %d", queue, size, 1)
		}
	}
}

func TestSizeEmptyQueue(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		size := queue.Size()
		if size != 0 {
			t.Errorf("%T size is %d, but we expected it to be %d", queue, size, 0)
		}
	}
}

func TestEnqueueCircularRight(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		for i := 0; i < size; i++ {
			queue.Enqueue(i)
		}
		for i := 0; i < size-2; i++ {
			queue.Dequeue()
		}
		//front ==> 8, 9 <== rear
		for i := 10; i < 16; i++ {
			queue.Enqueue(i)
		}

		//front ==> 8, 9, 10, 11, 12, 13, 14, 15 <== rear
		//se for vetor, 10 a 15 foram inseridos de forma circular, no início

		for i := 8; i < 16; i++ {
			val, err := queue.Front()
			if val != i {
				t.Errorf("Value dequeued from %T is %d, but we expected it to be %d", queue, val, i)
			}
			if err != nil {
				t.Errorf(err.Error())
			}
			queue.Dequeue()
		}

		size := queue.Size()
		if size != 0 {
			t.Errorf("%T size is %d, but we expected it to be %d", queue, size, 0)
		}
	}
}

func TestEnqueueCircularLeft(t *testing.T) {
	defer setupTest()()
	for _, queue := range queues {
		for i := 0; i < size; i++ {
			queue.Enqueue(i)
		}
		for i := 0; i < size-2; i++ {
			queue.Dequeue()
		}
		//front ==> 8, 9 <== rear
		for i := 10; i < 16; i++ {
			queue.Enqueue(i)
		}

		//front ==> 8, 9, 10, 11, 12, 13, 14, 15 <== rear
		//se for vetor, 10 a 15 foram inseridos de forma circular, no início

		for i := 8; i < 16; i++ {
			val, err := queue.Front()
			if val != i {
				t.Errorf("Value dequeued from %T is %d, but we expected it to be %d", queue, val, i)
			}
			if err != nil {
				t.Errorf(err.Error())
			}
			queue.Dequeue()
		}

		size := queue.Size()
		if size != 0 {
			t.Errorf("%T size is %d, but we expected it to be %d", queue, size, 0)
		}
	}
}

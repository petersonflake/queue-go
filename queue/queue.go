package queue

import (
	"errors"
)

type Queue[T any] interface {
	Enqueue(items ...T)
	TryEnqueue(items ...T) error
	Dequeue() (T, error)
	Peek() (T, error)
	Empty() bool
	Count() int
}

var ErrNoSpaceInQueue = errors.New("No space left in queue")

var ErrQueueEmpty = errors.New("Queue is empty")

package queue

type Fifo[T any] struct {
	Data []T
	head int
	tail int
}

/*
 * Creates a new First-in-First-out queue
 */
func NewFifo[T any](minSize int) *Fifo[T] {
	//Find smallest power of two larger than the requested size
	//so we can use bitwise operations instead of modulo
	i, j := 1, 0
	for ; i < minSize; i <<= 1 {
		//Check for overflow
		if i < j {
			panic("Unable to find power of two larger than that minimum size")
		}
		j = i
	}
	q := new(Fifo[T])
	q.Data = make([]T, 0, i)
	return q
}

func (q *Fifo[T]) enqueue(item T) {
	// next = head + 1 modulo the length of the backing array
	// h & (n - 1) == h % n when n is a power of two
	nextIndex := (q.head + 1) & (cap(q.Data) - 1)
	if nextIndex == q.tail {
		panic("No space left in queue")
	}
	q.Data[q.head] = item
	q.head = nextIndex
}

func (q *Fifo[T]) tryEnqueue(item T) error {
	nextIndex := (q.head + 1) & (len(q.Data) - 1)
	if nextIndex == q.head {
		return ErrNoSpaceInQueue
	}
	q.Data[q.head] = item
	q.head = nextIndex
	return nil
}

/*
 * Enqueue an item.  Panics if there is no space available.
 */
func (q *Fifo[T]) Enqueue(items ...T) {
	for _, i := range items {
		q.enqueue(i)
	}
}

/*
 * Enqueue an item, or return an error if there is not sufficient space
 */
func (q *Fifo[T]) TryEnqueue(items ...T) error {
	for _, i := range items {
		err := q.tryEnqueue(i)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
 * If there are items in queue, return the next one
 * otherwise, return an error
 */
func (q *Fifo[T]) Dequeue() (T, error) {
	var item T
	if q.tail == q.head {
		return item, ErrQueueEmpty
	}
	item = q.Data[q.tail]
	q.tail = (q.tail + 1) & (len(q.Data) - 1)
	return item, nil
}

/*
 * Return the next item in queue without removing it
 */
func (q *Fifo[T]) Peek() (T, error) {
	var item T
	if q.tail == q.head {
		return item, ErrQueueEmpty
	}
	item = q.Data[q.tail]
	return item, nil
}

/*
 * Check if the queue is empty
 */
func (q *Fifo[T]) Empty() bool {
	return q.head == q.tail
}

/*
 * Return the number of items in the queue
 */
func (q *Fifo[T]) Count() int {
	if q.head >= q.tail {
		return q.head - q.tail
	} else {
		return len(q.Data) - q.tail + q.head
	}
}

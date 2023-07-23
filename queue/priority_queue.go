package queue

/*
 * Priority queue.
 */
type PriorityQueue[T any] struct {
	Data    []T
	Compare func(left, right T) bool
	count   int
}

func NewPriorityQueue[T any](compare func(left, right T) bool) *PriorityQueue[T] {
	q := new(PriorityQueue[T])
	q.Compare = compare
	return q
}

func (pq *PriorityQueue[T]) Enqueue(items ...T) {
	for _, item := range items {
		pq.enqueue(item)
	}
}

func (pq *PriorityQueue[T]) TryEnqueue(items ...T) error {
	pq.Enqueue(items...)
	return nil
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	var item T
	if pq.count == 0 {
		return item, ErrQueueEmpty
	}

	if pq.count == 1 {
		pq.count--
		return pq.Data[0], nil
	}

	item = pq.Data[0]
	pq.Data[0] = pq.Data[pq.count-1]
	pq.count--
	pq.heap(0)
	return item, nil
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	var item T
	if pq.count == 0 {
		return item, ErrQueueEmpty
	}

	return pq.Data[0], nil
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.count == 0
}

func (pq *PriorityQueue[T]) Count() int {
	return pq.count
}

// Helper functions
func (pq *PriorityQueue[T]) swap(x, y int) {
	swp := pq.Data[x]
	pq.Data[x] = pq.Data[y]
	pq.Data[y] = swp
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) >> 1
}

func (pq *PriorityQueue[T]) heap(index int) {
	l := left(index)
	r := right(index)
	max := index

	if l < pq.count && pq.Compare(pq.Data[l], pq.Data[index]) {
		max = l
	}

	if r < pq.count && pq.Compare(pq.Data[r], pq.Data[max]) {
		max = r
	}

	if max != index {
		pq.swap(index, max)
		pq.heap(max)
	}
}

func (pq *PriorityQueue[T]) enqueue(item T) {
	if pq.count < len(pq.Data) {
		pq.Data[pq.count] = item
	} else {
		pq.Data = append(pq.Data, item)
	}

	i := pq.count
	for i != 0 && pq.Compare(pq.Data[i], pq.Data[parent(i)]) {
		pq.swap(i, parent(i))
		i = parent(i)
	}
	pq.count++
}

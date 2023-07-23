package queue

import "testing"

// FIFO tests
func TestFifoSize(t *testing.T) {
	q := NewFifo[int](13)
	if cap(q.Data) != 16 {
		t.Fatalf("Expected allocation for 16 elements, got %d\n", cap(q.Data))
	}
}

// Priority Queue tests
func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int](func (a, b int) bool {
		return a > b
	})

	q.Enqueue(3, 2, 1)
	if i, err := q.Dequeue(); err != nil || i != 3 {
		t.Fatal("Expected highest priority item to be 3")
	}

	if i, err := q.Dequeue(); err != nil || i != 2 {
		t.Fatal("Expected highest priority item to be 2")
	}

	if i, err := q.Dequeue(); err != nil || i != 1 {
		t.Fatal("Expected highest priority item to be 1")
	}

	if _, err := q.Dequeue(); err != ErrQueueEmpty {
		t.Fatalf("Queue should be empty, have %d items\n", q.count)
	}
}

func TestPriorityQueueOrdering(t *testing.T) {
	q := NewPriorityQueue[int](func (a, b int) bool { return a > b })
	q.Enqueue(1, 4, 2, 5, -1, 43, 0)
	if q.Count() != 7 {
		t.Fatalf("Queue should have 7 items, have %d\n", q.Count())
	}

	prev, err := q.Dequeue()
	if err != nil {
		t.Fatal("Queue should have items, got error")
	}
	expected_count := 6
	for !q.Empty() {
		if expected_count != q.Count() {
			t.Fatalf("Expected count %d, actual %d\n", expected_count, q.Count())
		}
		nxt, err := q.Dequeue()
		expected_count--
		if err != nil {
			t.Fatal("Queue should still have items")
		}
		if prev < nxt {
			t.Fatal("Items should be sorted in descending order")
		}
		if expected_count != q.Count() {
			t.Fatalf("Expected count %d, actual %d\n", expected_count, q.Count())
		}
	}
}

func TestPeekPop(t *testing.T) {
	q := NewPriorityQueue[int](func (a, b int) bool { return a > b })
	q.Enqueue(1, 4, 5)
	for i := 0; i < 3; i++ {
		peek, err := q.Peek()
		if err != nil {
			t.Fatal("Expected queue to have more items")
		}

		pop, err := q.Dequeue()
		if err != nil {
			t.Fatal("Expected queue to have more items (dequeue)")
		}

		if peek != pop {
			t.Fatalf("%d vs %d, should be the same\n", peek, pop)
		}
	}

	if !q.Empty() {
		t.Fatal("Queue should be empty")
	}
}

package queue

import "errors"

type Queue[T any] struct {
	Elements []T
	Size     int
}

func (q *Queue[T]) GetLength() int {
	return len(q.Elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.GetLength() == 0
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var nilT T
		return nilT, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

func (q *Queue[T]) Enqueue(elem T) error {
	if q.GetLength() == q.Size {
		return errors.New("queue overflow")
	}
	q.Elements = append(q.Elements, elem)
	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var nilT T
		return nilT, errors.New("queue mepty")
	} else {
		res := q.Elements[0]
		if q.GetLength() == 1 {
			q.Elements = nil
		} else {
			q.Elements = q.Elements[1:]
		}
		return res, nil
	}
}

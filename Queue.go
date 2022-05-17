package data_structures

import (
    "errors"
    "reflect"
)

var queueType reflect.Type

// represents a queue structure via an array with any type
type Queue struct {
	items []any
}

// add a value at the end; panic if try add another type
func (q *Queue) Enqueue(value any) {
	/* check valye type */
	if len(q.items) == 0 {
		queueType = reflect.TypeOf(value)
	}
	if reflect.TypeOf(value) != queueType {
		panic("stack type error")
	}
	q.items = append(q.items, value)
}

// remove and return a first value
func (q *Queue) Dequeue() (any, error) {
	if q.Empty() {
		return '0', errors.New("empty stack")
	}

	firstValue := q.items[0]
	q.items = q.items[1:]
	return firstValue, nil
}

// return true if queue is empty and false if not
func (q *Queue) Empty() bool {
	if len(q.items) > 0 {
		return false
	}
	return true
}

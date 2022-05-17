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
func (s *Queue) Enqueue(value any) {
	/* check valye type */
	if len(s.items) == 0 {
		queueType = reflect.TypeOf(value)
	}
	if reflect.TypeOf(value) != queueType {
		panic("stack type error")
	}
	s.items = append(s.items, value)
}

// remove and return a value at the end
func (s *Queue) Dequeue() (any, error) {
	if len(s.items) < 1 {
		return '0', errors.New("empty stack")
	}

	firstValue := s.items[0]
	s.items = s.items[1:]
	return firstValue, nil
}

// return true if queue is empty and false if not
func (s *Queue) Empty() bool {
	if len(s.items) > 0 {
		return false
	}
	return true
}

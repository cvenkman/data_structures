package data_structures

import (
    "errors"
    "reflect"
)

var stackType reflect.Type

// represents a stack structure via an array with any type
type Stack struct {
	items []any
}

// add a value at the end; panic if try add another type
func (s *Stack) Push(value any) {
    /* check valye type */
    if len(s.items) == 0 {
        stackType = reflect.TypeOf(value)
    }
    if reflect.TypeOf(value) != stackType {
        panic("stack type error")
    }
	s.items = append(s.items, value)
}

// remove and return a value at the end
func (s *Stack) Pop() (any, error) {
	lastIndex := len(s.items) - 1
	if s.Empty() {
		return '0', errors.New("empty stack")
	}

	lastValue := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastValue, nil
}

// return true if stack is empty and false if not
func (s *Stack) Empty() bool {
	if len(s.items) > 0 {
		return false
	}
	return true
}

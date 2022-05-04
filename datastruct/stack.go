package datastruct

import "errors"

type Stack []rune

func (s Stack) Push(v rune) Stack {
    return append(s, v)
}

func (s Stack) Pop() (Stack, rune, error) {
    l := len(s)
    if l < 1 {
        err := errors.New("empty stack")
        return nil, '0', err
    }
    return  s[:l-1], s[l-1], nil
}

func (s Stack) Empty() bool {
    if len(s) > 0 {
        return false
    }
    return true
}
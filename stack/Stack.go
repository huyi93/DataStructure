package stack

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func InitStack() *Stack {
	return new(Stack)
}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Push(value interface{}) {
	//fmt.Printf("%T, %T\n", s, *s)
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("Out of index, len is 0. ")
	}
	tempStack := *s
	a := tempStack[len(*s)-1]
	*s = tempStack[:len(*s)-1]
	return a, nil
}

func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("Out of index, len is 0. ")
	}
	return s[len(s)-1], nil
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Clear() {
	if len(*s) != 0 {
		*s = Stack{}
	}
}

// N = (N/d)*d + N%d
func Conversion(num int, binary int) {
	s := InitStack()
	for num > 0 {
		s.Push(num % binary)
		num = num / binary
	}
	for !s.IsEmpty() {
		num, _ := s.Pop()
		fmt.Print(num)
	}
}

func LineEdit(line string) *Stack {
	s := InitStack()
	for _, v := range line {
		switch string(v) {
		case "#":
			if !s.IsEmpty() {
				s.Pop()
			}
		case "@":
			s.Clear()
		default:
			s.Push(string(v))
		}

	}
	return s
}

package data_structures

import (
	"fmt"
	"sync"
)

type Stack struct {
	items []int
	lock  sync.RWMutex
}

func (s *Stack) New() *Stack {
	s.items = []int{}
	return s
}

func (s *Stack) Push(i int) {
	s.lock.Lock()
	s.items = append(s.items, i)
	s.lock.Unlock()
}

func (s *Stack) Pop() *int {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Print() {
	if !s.IsEmpty() {
		pop := s.Pop()
		fmt.Println(*pop, " ", pop)
		s.Print()
		s.Push(*pop)
	}
}

func (s *Stack) Length() *int {

	cont := new(int)

	if !s.IsEmpty() {
		pop := s.Pop()
		*cont = 1 + *s.Length()
		s.Push(*pop)
	}

	return cont
}

func (s *Stack) Copy(r *Stack) {

	if !s.IsEmpty() {
		pop := s.Pop()
		s.Copy(r)
		s.Push(*pop)
		r.Push(*pop)
	}
}

func (s *Stack) Reverse(r *Stack) {

	if !s.IsEmpty() {
		pop := s.Pop()
		r.Push(*pop)
		s.Reverse(r)
		s.Push(*pop)
	}
}

func (s *Stack) SetLastElement(v int) {

	if !s.IsEmpty() {
		pop := s.Pop()
		s.SetLastElement(v)
		s.Push(*pop)
	} else {
		s.Push(v)
	}
}

func (s *Stack) GetLastElement() *int{

	var result int

	if !s.IsEmpty() {
		pop := s.Pop()
		if !s.IsEmpty() {
			s.GetLastElement()
			s.Push(*pop)
		} else {
			result = *pop
		}
	}
	return &result
}


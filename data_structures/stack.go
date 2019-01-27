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

func (s *Stack) GetLastElement() *int {

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

func (s *Stack) FindElement(v int) int {
	var r int
	if !s.IsEmpty() {
		pop := s.Pop()
		if *pop == v {
			s.Push(*pop)
			return *pop
		}
		r = s.FindElement(v)
		s.Push(*pop)
	}
	return r
}

func (s *Stack) RemoveElement(v int) bool {
	var r bool
	if !s.IsEmpty() {
		pop := s.Pop()
		if *pop == v {
			return true
		}
		r = s.RemoveElement(v)
		s.Push(*pop)
	}
	return r
}

func (s *Stack) FindElementAndCount(v int, r int, c int) (int, int) {
	if !s.IsEmpty() {
		pop := s.Pop()
		if *pop == v {
			r = *pop
			c++
		}
		r, c = s.FindElementAndCount(v, r, c)
		s.Push(*pop)
	}
	return r, c
}

func (s *Stack) RemoveElementResultAfterSubtract(v int) bool {
	lastElement := s.GetLastElement()
	result := *lastElement - v
	return s.RemoveElement(result)
}

func (s *Stack) RemoveElementResultAfterSubtract2(v *int) (r bool) {

	if !s.IsEmpty() {
		pop := s.Pop()
		if !s.IsEmpty(){
			r = s.RemoveElementResultAfterSubtract2(v)
		} else {
			*v = *pop - *v
		}
		if *v == *pop {
			r = true
		} else {
			s.Push(*pop)
		}
	}
	return
}

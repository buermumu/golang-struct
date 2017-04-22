package stack

import _ "fmt"

type Element struct {
	next *Element
	data interface{}
}

func (e *Element) Value() interface{} {
	return e.data
}

type Stack struct {
	root Element
	len  int
}

func (s *Stack) Init() *Stack {
	s.root.next = &Element{}
	s.len = 0
	return s
}

func New() *Stack {
	return new(Stack).Init()
}

func (s *Stack) GetTop() *Element {
	if s.len == 0 {
		return nil
	}
	return s.root.next
}

func (s *Stack) Pop() *Element {
	if s.len == 0 {
		return nil
	}
	e := s.root.next
	s.root.next = e.next
	s.len--
	return e
}

func (s *Stack) Push(v interface{}) bool {
	if v == nil {
		return false
	}
	h := s.root.next
	e := Element{h, v}
	s.root.next = &e
	s.len++
	return true
}

func (s *Stack) Clear() int {
	if s.len == 0 {
		return s.len
	}
	h := s.root.next
	e := h.next
	h = nil
	s.len--
	t := 1
	for {
		if e.next == nil {
			break
		}
		h = e.next
		e = nil
		e = h
		s.len--
		t++
	}
	return t
}

func (s *Stack) Destroy() {

}

func (s *Stack) Length() int {
	return s.len
}

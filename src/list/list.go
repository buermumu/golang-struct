package list

import (
	"fmt"
)

type Element struct {
	next  *Element
	value interface{}
}

func (e *Element) Value() interface{} { return e.value }

func (e *Element) Next() *Element { return e.next }

type List struct {
	root Element
	len  int
}

func New() *List { return new(List).Init() }

func (l *List) Init() *List {
	l.root.next = &Element{}
	l.len = 0
	return l
}

func (l *List) Each() {
	e := l.root.next
	for {
		if e.next == nil {
			break
		}
		fmt.Println(e)
		e = e.next
	}
}

func (l *List) InsertHead(v interface{}) {
	f := l.root.next
	l.root.next = &Element{f, v}
	l.len++
}

func (l *List) InsertByIdx(idx int, v interface{}) {
	if idx > l.len {
		panic("idx out of upper limit")
	}
	if idx < 0 {
		panic("idx not less than 0")
	}
	e := &l.root
	for i := 0; i < idx; i++ {
		e = e.next
	}
	e.next = &Element{e.next, v}
	l.len++
}

func (l *List) InsertAfter(ele *Element, v interface{}) bool {
	e := l.root.next
	for i := 1; i <= l.len; i++ {
		if e == ele {
			e.next = &Element{e.next, v}
			l.len++
			return true
		}
		e = e.next
	}
	return false
}

func (l *List) GetEleByIdx(idx int) *Element {
	if idx < 0 {
		panic("idx not less than 0")
	}
	if idx >= l.len {
		panic("idx out of upper limit")
	}
	e := l.root.next
	for i := 0; i < idx; i++ {
		e = e.next
	}
	return e
}

func (l *List) Delete(ele *Element) bool {
	e := l.root.next
	for {
		if e.next == nil {
			return false
		}
		if e.next == ele {
			e.next = ele.next
			ele = nil
			l.len--
			return true
		}
		e = e.next
	}
}

func (l *List) Len() int { return l.len }

package list

import (
	"fmt"
	"strings"

	"github.com/aktoro-lang/enum"
)

// List is a basic linked list implementation
type List interface {
	FirstRest() (interface{}, enum.Enum, bool)
}

type node struct {
	value interface{}
	next  List
}

func (n node) FirstRest() (interface{}, enum.Enum, bool) {
	return n.value, enum.Enum(n.next), true
}

type empty struct{}

func (empty) FirstRest() (interface{}, enum.Enum, bool) {
	return nil, enum.Enum(empty{}), false
}

func (n node) String() string {
	return toString(n)
}

func (e empty) String() string {
	return toString(e)
}

// New constructs a List
func New(elems ...interface{}) List {
	var l List = empty{}
	return Cons(l, elems...)
}

// Cons prepends an element to a list
func Cons(l List, elems ...interface{}) List {
	if len(elems) == 0 {
		return l
	}
	k := len(elems)
	for i := range elems {
		i = k - 1 - i
		el := elems[i]
		l = node{el, l}
	}
	return l
}

func toString(l List) string {
	buffer := []string{}
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			buffer = append(buffer, fmt.Sprintf("%v", el))
		} else {
			break
		}
	}
	str := strings.Join(buffer, ",")
	return "[" + str + "]"
}

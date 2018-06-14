package list

import (
	"fmt"
	"strings"

	"github.com/aktoro-lang/types"
)

// List is a basic linked list implementation
type List struct {
	value interface{}
	next  *List
	ok    bool
}

// FirstRest returns the first element, the rest of the list, and a boolean whether the list is completed
func (l *List) FirstRest() (interface{}, *List, bool) {
	if !l.ok {
		return l.value, l, l.ok
	}
	return l.value, l.next, l.ok
}

func (l *List) String() string {
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
	str := strings.Join(buffer, ", ")
	return "[" + str + "]"
}

// New constructs a List
func New(elems ...interface{}) *List {
	l := &List{nil, nil, false}
	return Cons(l, elems...)
}

// Cons prepends an element to a list
func Cons(l *List, elems ...interface{}) *List {
	if len(elems) == 0 {
		return l
	}
	k := len(elems)
	for i := range elems {
		i = k - 1 - i
		el := elems[i]
		l = &List{el, l, true}
	}
	return l
}

//At is an indexing function on list
func At(l *List, n types.AkInt) interface{} {
	listStart, _ := nth(l, n)
	el, _, _ := listStart.FirstRest()
	return el
}

func nth(l *List, n types.AkInt) (*List, bool) {
	var ok bool
	var rest *List
	for i := types.AkInt(0); ; i++ {
		_, rest, ok = l.FirstRest()
		if !ok {
			return nil, false
		} else if i == n {
			return l, true
		}
		l = rest
	}
}

// GetRange returns a list from index low to high, makes a copy of the list
func GetRange(l *List, low types.AkInt, high types.AkInt) *List {
	var first, cur, prev *List
	orig, ok := nth(l, low)
	if !ok {
		return nil
	}
	for i := low; i <= high; i++ {
		cur = &List{orig.value, nil, false}
		if first == nil {
			first = cur
			first.ok = true
		}
		if prev != nil {
			prev.next = cur
			prev.ok = true
		}
		prev = cur
		orig = orig.next
	}
	prev.next = &List{nil, nil, false}
	return first
}

// Drop returns a list with the first n elements removed
func Drop(l *List, n types.AkInt) *List {
	listStart, _ := nth(l, n)
	return listStart
}

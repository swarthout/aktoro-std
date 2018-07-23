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
			return l, false
		} else if i == n {
			return l, true
		}
		l = rest
	}
}

func emptyList() *List {
	return &List{nil, nil, false}
}

// GetRange returns a list from index low to high, makes a copy of the list
func GetRange(l *List, low types.AkInt, high types.AkInt) *List {
	var first, curr, prev *List
	orig, ok := nth(l, low)
	if !ok {
		return emptyList()
	}
	for i := low; i < high; i++ {
		if !orig.ok {
			break
		}
		curr = &List{orig.value, nil, false}
		if first == nil {
			first = curr
			first.ok = true
		}
		if prev != nil {
			prev.next = curr
			prev.ok = true
		}
		prev = curr
		orig = orig.next
	}
	prev.next = emptyList()
	prev.ok = true
	return first
}

func Length(l *List) interface{} {
	i := 0
	var ok bool
	for {
		if _, l, ok = l.FirstRest(); ok {
			i++
		} else {
			return i
		}
	}
}

func Empty(l *List) interface{} {
	return types.AkBool(!l.ok)
}

func Sort(l *List) interface{} {
	if ! l.ok || !l.next.ok {
		return l
	}
	var less func(interface{}, interface{}) interface{}
	switch l.value.(type) {
	case types.AkInt:
		less = func(p0 interface{}, p1 interface{}) interface{} {
			a := p0.(types.AkInt)
			b := p1.(types.AkInt)
			return a < b
		}
	case types.AkFloat:
		less = func(p0 interface{}, p1 interface{}) interface{} {
			a := p0.(types.AkFloat)
			b := p1.(types.AkFloat)
			return a < b
		}
	case types.AkString:
		less = func(p0 interface{}, p1 interface{}) interface{} {
			a := p0.(types.AkString)
			b := p1.(types.AkString)
			return a < b
		}
	}
	return SortBy(l, less)
}

func SortBy(l *List, less func(interface{}, interface{}) interface{}) interface{} {
	if !l.ok || !l.next.ok {
		return l
	}
	left, right := split(l)

	left = SortBy(left, less).(*List)
	right = SortBy(right, less).(*List)
	return merge(left, right, less)
}
func merge(left *List, right *List, less func(interface{}, interface{}) interface{}) *List {
	if !left.ok {
		return right
	}
	if !right.ok {
		return left
	}
	if less(left.value, right.value).(bool) {
		return Cons(merge(left.next, right, less), left.value)
	} else {
		return Cons(merge(left, right.next, less), right.value)
	}
}
func split(source *List) (*List, *List) {
	if !source.ok || !source.next.ok {
		return source, emptyList()
	}
	first := source.value
	second := source.next.value
	left, right := split(source.next.next)
	return Cons(left, first), Cons(right, second)
}

func Take(l *List, n types.AkInt) interface{} {
	var first, curr, prev *List
	orig := l
	if !l.ok || n == 0 {
		return emptyList()
	}
	for i := types.AkInt(0); i < n; i++ {
		if !orig.ok {
			break
		}
		curr = &List{orig.value, nil, false}
		if first == nil {
			first = curr
			first.ok = true
		}
		if prev != nil {
			prev.next = curr
			prev.ok = true
		}
		prev = curr
		orig = orig.next
	}
	prev.next = emptyList()
	prev.ok = true
	return first
}

func TakeWhile(l *List, f func(interface{}) interface{}) interface{} {
	var first, curr, prev *List
	orig := l
	for {
		if orig.ok {
			if !f(orig.value).(types.AkBool) {
				break
			}
			curr = &List{orig.value, nil, false}
			if first == nil {
				first = curr
				first.ok = true
			}
			if prev != nil {
				prev.next = curr
				prev.ok = true
			}
			prev = curr
			orig = orig.next
		} else {
			if first == nil {
				return emptyList()
			}
			break
		}

	}
	prev.next = emptyList()
	prev.ok = true
	return first
}

// Drop returns a list with the first n elements removed
func Drop(l *List, n types.AkInt) interface{} {
	//listStart, _ := nth(l, n)
	//return listStart
	var ok bool
	orig := l
	if !l.ok || n == 0 {
		return l
	}
	for i := types.AkInt(0); i < n; i++ {
		if !orig.ok {
			break
		}
		if _, l, ok = orig.FirstRest(); ok {
			orig = l
		} else {
			break
		}
	}
	return l
}

func DropWhile(l *List, f func(interface{}) interface{}) interface{} {
	orig := l
	var el interface{}
	var ok bool
	for {
		if el, l, ok = orig.FirstRest(); ok {
			if !f(el).(types.AkBool) {
				return orig
			}
			orig = l
		} else {
			break
		}

	}
	return l
}

func All(l *List, f func(interface{}) interface{}) interface{} {
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			if !f(el).(types.AkBool) {
				return false
			}
		} else {
			break
		}
	}
	return true
}

func Any(l *List, f func(interface{}) interface{}) interface{} {
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			if f(el).(types.AkBool) {
				return true
			}
		} else {
			break
		}
	}
	return false
}

func Find(l *List, f func(interface{}) interface{}) interface{} {
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			if f(el).(types.AkBool) {
				return el
			}
		} else {
			break
		}
	}
	return nil
}

func FindIndex(l *List, f func(interface{}) interface{}) interface{} {
	var el interface{}
	var ok bool
	index := 0
	for {
		if el, l, ok = l.FirstRest(); ok {
			if f(el).(types.AkBool) {
				return index
			}
		} else {
			break
		}
		index++
	}
	return nil
}

func First(l *List) interface{} {
	el, _, _ := l.FirstRest()
	return el
}

func Rest(l *List) interface{} {
	_, rest, _ := l.FirstRest()
	return rest
}

func Reverse(l *List) interface{} {
	newList := New()
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			newList = Cons(newList, el)
		} else {
			return newList
		}
	}
}

func Map(l *List, f func(interface{}) interface{}) interface{} {
	var head, curr, prev *List
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			curr = &List{f(el), nil, false}
			if head == nil {
				head = curr
				head.ok = true
			}
			if prev != nil {
				prev.next = curr
				prev.ok = true
			}
			prev = curr
		} else {
			if head == nil {
				return emptyList()
			}
			break
		}

	}
	prev.next = emptyList()
	prev.ok = true
	return head
}

func Reduce(l *List, acc interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			acc = f(acc, el)
		} else {
			break
		}
	}
	return acc
}

func Filter(l *List, f func(interface{}) interface{}) interface{} {
	var head, curr, prev *List
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			if f(el).(types.AkBool) {
				curr = &List{el, nil, false}
			}
			if head == nil {
				head = curr
				head.ok = true
			}
			if prev != nil {
				prev.next = curr
				prev.ok = true
			}
			prev = curr
		} else {
			if head == nil {
				return emptyList()
			}
			break
		}

	}
	prev.next = emptyList()
	prev.ok = true
	return head
}

func Each(l *List, f func(interface{}) interface{}) interface{} {
	var el interface{}
	var ok bool
	for {
		if el, l, ok = l.FirstRest(); ok {
			f(el)
		} else {
			break
		}
	}
	return nil
}

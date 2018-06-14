package dict

import (
	"fmt"
	"strings"

	"github.com/aktoro-lang/container/set"
	"github.com/aktoro-lang/enum"
)

// Dict represents a map (associative array).
type Dict struct {
	set set.Set
}

// New creates a new dict.
func New(kvs ...keyValue) Dict {
	d := Dict{set.New()}
	if len(kvs) > 0 {
		for _, kv := range kvs {
			d = Put(d, kv.key, kv.value)
		}
	}
	return d
}

// Put inserts a key-value pair into a map.
func Put(d Dict, k set.Entry, v interface{}) Dict {
	return Dict{d.set.Put(NewKeyValue(k, v))}
}

// Delete deletes a pair of a key and a value from a map.
func Delete(d Dict, k set.Entry) Dict {
	return Dict{d.set.Delete(k)}
}

// Get finds a value corresponding to a given key from a map.
// It returns nil if no value is found.
func Get(d Dict, k set.Entry) interface{} {
	e := d.set.Get(k)

	if e == nil {
		return nil
	}

	return e.(keyValue).value
}

// Member returns true if a key-value pair corresponding with a given key is
// included in a map, or false otherwise.
func Member(d Dict, k set.Entry) bool {
	return Get(d, k) != nil
}

// FirstRest returns a key-value pair in a map and a rest of the map.
// This method is useful for iteration.
// The key and value would be nil if the map is empty.
func (d Dict) FirstRest() (interface{}, enum.Enum, bool) {
	e, s, ok := d.set.FirstRest()
	d = Dict{s}
	return e, enum.Enum(d), ok
}

// Size returns a size of a map.
func Size(d Dict) int {
	return d.set.Size()
}

func (d Dict) String() string {
	buffer := []string{}
	var el interface{}
	var rest enum.Enum
	var kv keyValue
	var ok bool
	for {
		if el, rest, ok = d.FirstRest(); ok {
			d = rest.(Dict)
			kv = el.(keyValue)
			buffer = append(buffer, fmt.Sprintf("%v => %v", kv.key, kv.value))
		} else {
			break
		}
	}
	str := strings.Join(buffer, ", ")
	return "%{" + str + "}"
}

package dict

import (
	"github.com/aktoro-lang/container/set"
	"github.com/aktoro-lang/enum"
)

// Dict represents a map (associative array).
type Dict struct {
	set set.Set
}

// New creates a new dict.
func New() Dict {
	return Dict{set.New()}
}

// Put inserts a key-value pair into a map.
func (d Dict) Put(k set.Entry, v interface{}) Dict {
	return Dict{d.set.Put(newKeyValue(k, v))}
}

// Delete deletes a pair of a key and a value from a map.
func (d Dict) Delete(k set.Entry) Dict {
	return Dict{d.set.Delete(k)}
}

// Get finds a value corresponding to a given key from a map.
// It returns nil if no value is found.
func (d Dict) Get(k set.Entry) interface{} {
	e := d.set.Get(k)

	if e == nil {
		return nil
	}

	return e.(keyValue).value
}

// Include returns true if a key-value pair corresponding with a given key is
// included in a map, or false otherwise.
func (d Dict) Include(k set.Entry) bool {
	return d.Get(k) != nil
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
func (d Dict) Size() int {
	return d.set.Size()
}

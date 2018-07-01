package list

import (
	"fmt"
	"testing"
	"github.com/aktoro-lang/types"
)

func double(p0 interface{}) interface{} {
	x := p0.(int)
	return x * 2
}

func TestGetRange(t *testing.T) {
	l := New(1, 2, 3)
	r := GetRange(l, 0, 3)
	fmt.Println(r)
	r = GetRange(l, 0, 2)
	fmt.Println(r)

}

func TestMap(t *testing.T) {
	xs := New(0, 1, 2, 3, 4, 5, 6)
	twice := Map(xs, double).(*List)
	fmt.Println(twice)
}

type Car struct {
	name types.AkString
	year types.AkInt
}

func TestSort(t *testing.T) {
	empty := New()
	sorted := Sort(empty)
	fmt.Println(sorted)
	x := New(types.AkInt(8))
	sorted = Sort(x)
	fmt.Println(x)
	xs := New(types.AkInt(2011), types.AkInt(2009), types.AkInt(2018))
	sorted = Sort(xs)
	fmt.Println(sorted)
}

func TestSortByCars(t *testing.T) {
	cars := New(Car{types.AkString("mustang"), types.AkInt(2011)},
		Car{types.AkString("camaro"), types.AkInt(2009)},
		Car{types.AkString("corvette"), types.AkInt(2018)})

	byName := func(p0 interface{}, p1 interface{}) interface{} {
		a := p0.(Car)
		b := p1.(Car)
		return a.name < b.name
	}

	byAge := func(p0 interface{}, p1 interface{}) interface{} {
		a := p0.(Car)
		b := p1.(Car)
		return a.year < b.year
	}
	sorted := SortBy(cars, byName)
	fmt.Println("by name:", sorted)

	sorted = SortBy(cars, byAge)
	fmt.Println("by age: ", sorted)

}

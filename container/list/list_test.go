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

func TestDropWhile(t *testing.T) {
	xs := New(0, 1, 2, 3, 4, 5, 6)
	lessThanFour := func(p0 interface{}) interface{} {
		return p0.(int) < 4
	}
	small := DropWhile(xs, lessThanFour)
	fmt.Println(small)
}

func TestTake(t *testing.T) {
	empty := New()
	var leftover interface{}
	leftover = Take(empty, types.AkInt(0))
	fmt.Println("take 0 []", leftover)
	x := New(types.AkInt(8))
	leftover = Take(x, types.AkInt(0))
	fmt.Println("take 0 [8]", leftover)
	leftover = Take(x, types.AkInt(1))
	fmt.Println("take 1 [8]", leftover)
	leftover = Take(x, types.AkInt(2))
	fmt.Println("take 2 [8]", leftover)
	xs := New(types.AkInt(2011), types.AkInt(2009), types.AkInt(2018))
	leftover = Take( xs, types.AkInt(0))
	fmt.Println("take 0 [2011, 2009, 2018]",leftover)
	leftover = Take(xs, types.AkInt(1))
	fmt.Println("take 1 [2011, 2009, 2018]",leftover)
	leftover = Take(xs, types.AkInt(2))
	fmt.Println("take 2 [2011, 2009, 2018]",leftover)
	leftover = Take(xs, types.AkInt(4))
	fmt.Println("take 4 [2011, 2009, 2018]",leftover)
}

func TestDrop(t *testing.T) {
	empty := New()
	var leftover interface{}
	leftover = Drop(empty, types.AkInt(0))
	fmt.Println("drop 0 []", leftover)
	x := New(types.AkInt(8))
	leftover = Drop(x, types.AkInt(0))
	fmt.Println("drop 0 [8]", leftover)
	leftover = Drop(x, types.AkInt(1))
	fmt.Println("drop 1 [8]", leftover)
	leftover = Drop(x, types.AkInt(2))
	fmt.Println("drop 2 [8]", leftover)
	xs := New(types.AkInt(2011), types.AkInt(2009), types.AkInt(2018))
	leftover = Drop( xs, types.AkInt(0))
	fmt.Println("drop 0 [2011, 2009, 2018]",leftover)
	leftover = Drop(xs, types.AkInt(1))
	fmt.Println("drop 1 [2011, 2009, 2018]",leftover)
	leftover = Drop(xs, types.AkInt(2))
	fmt.Println("drop 2 [2011, 2009, 2018]",leftover)
	leftover = Drop(xs, types.AkInt(4))
	fmt.Println("drop 4 [2011, 2009, 2018]",leftover)
}

func TestLength(t *testing.T) {
	empty := New()
	x := New(types.AkInt(8))
	xs := New(types.AkInt(2011), types.AkInt(2009), types.AkInt(2018))
	fmt.Println(Length(empty))
	fmt.Println(Length(x))
	fmt.Println(Length(xs))
}
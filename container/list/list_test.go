package list

import (
	"fmt"
	"testing"
)

func TestGetRange(t *testing.T) {
	l := New(1, 2, 3)
	r := GetRange(l, 0, 3)
	fmt.Println(r)
	r = GetRange(l, 0, 2)
	fmt.Println(r)
}

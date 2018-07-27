package types

import (
	"github.com/aktoro-lang/container/set"
)

// AkString is the base string type for aktoro
type AkString string

// Hash returns a hash of the string
func (s AkString) Hash() uint32 {
	return uint32(len(s))
}

// Equal tests equality of two strings
func (s AkString) Equal(e set.Entry) bool {
	otherString, ok := e.(AkString)

	if !ok {
		return false
	}

	return s == otherString
}

// AkInt is the base int type for aktoro
type AkInt int

// Hash hashes the int to a uint32
func (i AkInt) Hash() uint32 {
	return uint32(i)
}

// Equal tests equality of two ints
func (i AkInt) Equal(e set.Entry) bool {
	j, ok := e.(AkInt)

	if !ok {
		return false
	}

	return i == j
}

// AkFloat is the base float type for aktoro
type AkFloat float64

// Hash hashes the float to a uint32
func (i AkFloat) Hash() uint32 {
	return uint32(i)
}

// Equal tests equality of two floats
func (i AkFloat) Equal(e set.Entry) bool {
	j, ok := e.(AkFloat)

	if !ok {
		return false
	}

	return i == j
}

// AkBool is the base bool type for aktoro
type AkBool bool

// Hash hashes the float to a uint32
func (i AkBool) Hash() uint32 {
	if i {
		return 1
	}
	return 0
}

// Equal tests equality of two floats
func (i AkBool) Equal(e set.Entry) bool {
	j, ok := e.(AkBool)

	if !ok {
		return false
	}

	return i == j
}

type AkVariant interface {
	AkVariantConstructor()
}

type AkOption interface {
	AkOption()
}

type Some struct {
	P0 interface{}
}

func (Some) AkOption() {}

func (Some) AkVariantConstructor() {}

type None struct {
}

func (None) AkOption() {}

func (None) AkVariantConstructor() {}

type AkResult interface {
	AkResult()
}

type Ok struct {
	P0 interface{}
}

func (Ok) AkResult() {}

func (Ok) AkVariantConstructor() {}

type Err struct {
	P0 AkString
}

func (Err) AkResult() {}

func (Err) AkVariantConstructor() {}

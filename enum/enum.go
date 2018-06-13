package enum

// Enum interface is base interface for List, Dict, Array, Stream, and Set
type Enum interface {
	// FirstRest returns the first element, an Enum with the first element remove,
	// and a boolean if there are more elements left
	FirstRest() (interface{}, Enum, bool)
}

// At is an indexing operation
func At(e Enum, index int) interface{} {
	return nil
}

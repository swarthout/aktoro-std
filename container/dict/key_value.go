package dict

import "github.com/aktoro-lang/container/set"

type keyValue struct {
	key   set.Entry
	value interface{}
}

func newKeyValue(k set.Entry, v interface{}) keyValue {
	return keyValue{k, v}
}

func (kv keyValue) Hash() uint32 {
	return kv.key.Hash()
}

func (kv keyValue) Equal(e set.Entry) bool {
	if k, ok := e.(keyValue); ok {
		e = k.key
	}

	return kv.key.Equal(e)
}

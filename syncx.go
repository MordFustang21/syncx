package syncx

import (
	"sync"
)

// NewMap initializes a map.
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m: sync.Map{},
	}
}

// Map is a generic wrapper around sync.Map to give better type hinting.
type Map[K comparable, V any] struct {
	m sync.Map
}

// Store will add a key, value to the map.
func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

// Delete will remove and entry from the map.
func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

// Load will get a value from the map.
func (m *Map[K, V]) Load(key K) (V, bool) {
	var zero V
	ret, ok := m.m.Load(key)
	if ret == nil {
		return zero, ok
	}

	return ret.(V), ok
}

// Range will iterate over all entries in the map.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}


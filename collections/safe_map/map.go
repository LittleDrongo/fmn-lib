package safe_map

import "sync"

type storage[K comparable, V any] map[K]V

// SafeMap is a concurrency-safe generic map backed by sync.RWMutex.
type SafeMap[K comparable, V any] struct {
	sync.RWMutex
	storage[K, V]
}

// New creates a new concurrency-safe map and optionally fills it with entries.
func New[K comparable, V any](entries ...map[K]V) *SafeMap[K, V] {
	m := &SafeMap[K, V]{
		storage: make(storage[K, V]),
	}

	if len(entries) == 0 {
		return m
	}

	for key, value := range entries[0] {
		m.storage[key] = value
	}
	return m
}

func (m *SafeMap[K, V]) ensureStorage() {
	if m.storage == nil {
		m.storage = make(storage[K, V])
	}
}

// Set stores a value by key.
func (m *SafeMap[K, V]) Set(key K, value V) {
	m.Lock()
	defer m.Unlock()

	m.ensureStorage()
	m.storage[key] = value
}

// Get returns a value by key and whether it exists.
func (m *SafeMap[K, V]) Get(key K) (V, bool) {
	m.RLock()
	defer m.RUnlock()

	value, ok := m.storage[key]
	return value, ok
}

// Exists reports whether a key exists in the map.
func (m *SafeMap[K, V]) Exists(key K) bool {
	m.RLock()
	defer m.RUnlock()

	_, ok := m.storage[key]
	return ok
}

// Delete removes a value by key.
func (m *SafeMap[K, V]) Delete(key K) {
	m.Lock()
	defer m.Unlock()

	delete(m.storage, key)
}

// Len returns the number of entries in the map.
func (m *SafeMap[K, V]) Len() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.storage)
}

// Clear removes all entries from the map.
func (m *SafeMap[K, V]) Clear() {
	m.Lock()
	defer m.Unlock()

	clear(m.storage)
}

// Keys returns a snapshot of all keys.
func (m *SafeMap[K, V]) Keys() []K {
	m.RLock()
	defer m.RUnlock()

	keys := make([]K, 0, len(m.storage))
	for key := range m.storage {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a snapshot of all values.
func (m *SafeMap[K, V]) Values() []V {
	m.RLock()
	defer m.RUnlock()

	values := make([]V, 0, len(m.storage))
	for _, value := range m.storage {
		values = append(values, value)
	}
	return values
}

// Snapshot returns a copy of the current map state.
func (m *SafeMap[K, V]) Snapshot() map[K]V {
	m.RLock()
	defer m.RUnlock()

	snapshot := make(map[K]V, len(m.storage))
	for key, value := range m.storage {
		snapshot[key] = value
	}
	return snapshot
}

// Range iterates over a snapshot of all entries until fn returns false.
func (m *SafeMap[K, V]) Range(fn func(K, V) bool) {
	type entry struct {
		key   K
		value V
	}

	m.RLock()
	entries := make([]entry, 0, len(m.storage))
	for key, value := range m.storage {
		entries = append(entries, entry{key: key, value: value})
	}
	m.RUnlock()

	for _, entry := range entries {
		if !fn(entry.key, entry.value) {
			return
		}
	}
}

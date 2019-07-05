package dstr;

import "sync"

type SafeMap struct {
	*sync.RWMutex
	Map map[interface{}]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		RWMutex: new(sync.RWMutex),
		Map:     make(map[interface{}]interface{}),
	}
}

func (m *SafeMap) Get(k interface{}) interface{} {

	m.RWMutex.RLock()

	defer m.RWMutex.RUnlock()

	if val, ok := m.Map[k]; ok {

		return val

	}
	return nil

}

func (m *SafeMap) Set(k interface{}, v interface{}) bool {

	m.RWMutex.Lock()

	defer m.RWMutex.Unlock()

	if val, ok := m.Map[k]; !ok {

		m.Map[k] = v

	} else if val != v {

		m.Map[k] = v

	} else {

		return false

	}

	return true

}

// Returns true if k is exist in the map.

func (m *SafeMap) Check(k interface{}) (interface{}, bool) {

	m.RWMutex.RLock()

	defer m.RWMutex.RUnlock()

	if _, ok := m.Map[k]; !ok {

		return nil, false

	}

	return m.Map[k], true

}

func (m *SafeMap) Delete(k interface{}) {

	m.RWMutex.Lock()

	defer m.RWMutex.Unlock()

	delete(m.Map, k)

}

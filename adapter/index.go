package main

var AdapterMap AdapterMap = &AdapterMap {
	Map: make(map[string]interface{})
}

type AdapterMap struct {
    Map    map[string]interface{}
}

// Set ...
func (m *M) Set(key, value interface{}) {
	if m.Map[key] != nil {
		return
	}
    m.Map[key] = value
}

// Get ...
func (m *M) Get(key string) string {
    return m.Map[key]
}

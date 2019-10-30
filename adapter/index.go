package adapter

import (
	"github.com/wujiyu115/yuqueg"
)

//AdapterMap of map
type AdapterMap struct {
	Map map[string]interface{}
}

// Set ...
func (m *AdapterMap) Set(key string, value interface{}) {
	if m.Map[key] != nil {
		return
	}
	m.Map[key] = value
}

// Get ...
func (m *AdapterMap) Get(key string) interface{} {
	return m.Map[key]
}

var AMap = &AdapterMap{
	Map: make(map[string]interface{}),
}

func init() {
	AMap.Set("markdown", GenMarkDown)
	yuqueg.L.Info(AMap)
}

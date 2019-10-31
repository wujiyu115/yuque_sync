package adapter

import (
	"fmt"
	"io/ioutil"

	"github.com/wujiyu115/yuqueg"
	"github.com/wujiyu115/yuques/util"
)

//AdapterMap of map
type AdapterMap struct {
	Map map[string]interface{}
}

var (
	//AMap of adapter
	AMap = &AdapterMap{
		Map: make(map[string]interface{}),
	}
	// L logger
	L = yuqueg.L
)

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

func init() {
	AMap.Set("markdown", GenMarkDown)
	AMap.Set("hexo", GenHexo)
	// yuqueg.L.Info(AMap)
}

//GenFromFile from file
func GenFromFile(fileName string, typ string, targetFile string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		L.Error(fmt.Sprintf("empty fileName: fileName:%s, adapter:%s", fileName, typ))
		return
	}
	fun := AMap.Get(typ)
	if fun == nil {
		L.Error(fmt.Sprintf("empty adapter:  fileName:%s, adapter:%s", fileName, typ))
		return
	}
	r, err1 := util.Call(fun, yuqueg.Doc{
		Body: string(data),
	})
	if err1 != nil {
		L.Error(fmt.Sprintf("call adapter.fail: fileName:%s, adapter:%s", fileName, typ))
		return
	}
	d := r[0].String()
	ioutil.WriteFile(targetFile, []byte(d), 0644)
}

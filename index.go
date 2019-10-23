package main

import (
	"encoding/json"
	"fmt"
)

var (
	//Cfg base config
	Cfg SyncConfig
)

//DoAction do
func DoAction() {
	Cfg = LoadConfig()
	yuque := NewYuQueService(Cfg)
	// d, err := yuque.GetArticle("unrhpa")
	d, err := yuque.GetArticles()
	// d, err := yuque.GetToc()
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString, _ := json.Marshal(d)
	fmt.Println(string(jsonString))
	// readYuqueCache(Cfg.CachePath)
}

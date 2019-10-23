package main

import (
	"github.com/wujiyu115/yuqueg"
)

//Dwong get doc
type Dwong struct {
	yuque *YuQueService
}

//NewDwong of yuque
func NewDwong(config SyncConfig) *Dwong {
	s := &Dwong{
		yuque: NewYuQueService(config),
	}
	return s
}

//FetchArticle
func (d Dwong) FetchArticle(item yuqueg.DocBookDetail) {
	// article := d.yuque.GetArticle(item.Slug)
}

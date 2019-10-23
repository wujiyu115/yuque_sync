package main

import (
	"fmt"

	"github.com/wujiyu115/yuqueg"
)

//YuQueService get doc
type YuQueService struct {
	yuqueg.Service
	nameSapce string
}

//NewYuQueService of yuque
func NewYuQueService(config SyncConfig) *YuQueService {
	s := &YuQueService{
		nameSapce: fmt.Sprintf("%s/%s", config.Login, config.Repo),
	}
	s.Init(config.Token)
	return s
}

//GetArticle get detail of doc
func (y YuQueService) GetArticle(slug string) (interface{}, error) {
	return y.Doc.Get(y.nameSapce, slug, &yuqueg.DocGet{Raw: 1})
}

//GetArticles all articles
func (y YuQueService) GetArticles() (interface{}, error) {
	return y.Doc.List(y.nameSapce)
}

//GetToc of repo
func (y YuQueService) GetToc() (interface{}, error) {
	return y.Repo.GetToc(y.nameSapce)
}

package deal

import (
	"github.com/wujiyu115/yuqueg"
)

//YuQueService get doc
type YuQueService struct {
	yuqueg.Service
	nameSpace string
}

//NewYuQueService of yuque
func NewYuQueService(config SyncConfig) *YuQueService {
	s := &YuQueService{
		nameSpace: GenNameSpace(config),
	}
	s.Init(config.Token)
	return s
}

//GetArticle get detail of doc
func (y YuQueService) GetArticle(slug string) (yuqueg.DocDetail, error) {
	return y.Doc.Get(y.nameSpace, slug, &yuqueg.DocGet{Raw: 1})
}

//GetArticles all articles
func (y YuQueService) GetArticles() (yuqueg.BookDetail, error) {
	return y.Doc.List(y.nameSpace)
}

//GetToc of repo
func (y YuQueService) GetToc() (interface{}, error) {
	return y.Repo.GetToc(y.nameSpace)
}

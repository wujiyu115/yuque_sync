package main

import (
	"time"

	"github.com/wujiyu115/yuqueg"
)

var (
	//L  logger
	L = yuqueg.L
)

//Downg get doc
type Downg struct {
	yuque     *YuQueService
	store     *Store
	nameSpace string
	config    SyncConfig
}

//NewDowng of yuque
func NewDowng(config SyncConfig) *Downg {
	s := &Downg{
		yuque:     NewYuQueService(config),
		store:     NewStore(config.CachePath),
		nameSpace: GenNameSpace(config),
		config:    config,
	}
	L.Info("namespace is:", s.nameSpace)
	return s
}

//FetchArticle one of fetch
func (d *Downg) FetchArticle(item yuqueg.DocBookDetail) {
	_, err := d.yuque.GetArticle(item.Slug)
	if err != nil {
		L.Error(err)
		return
	}
	d.store.AddArticle(d.nameSpace, item.Slug, item)
	L.Info("fetch article slug: ", item.Slug)
}

//filter
func (d *Downg) filterArticles(articles *yuqueg.BookDetail) {
	if articles.Data == nil || len(articles.Data) == 0 {
		L.Warn("articles is empty")
		return
	}
	L.Info("article amount: ", len(articles.Data))
	for i := len(articles.Data) - 1; i >= 0; i-- {
		article := articles.Data[i]
		// L.Info("article ", d.config.OnlyPub, time.Time.IsZero(article.PublishedAt))
		if d.config.OnlyPub && time.Time.IsZero(article.PublishedAt) {
			articles.Data = append(articles.Data[:i],
				articles.Data[i+1:]...)
		}
	}
	L.Info("real article amount: ", len(articles.Data))
}

//fetchArticles one of fetch
func (d *Downg) fetchArticles() {
	articles, err := d.yuque.GetArticles()
	if err != nil {
		L.Error(err)
		return
	}
	d.filterArticles(&articles)

	for _, article := range articles.Data {
		item := d.store.findArticle(d.nameSpace, article.Slug)
		if item != nil && article.UpdatedAt.Equal(item.UpdatedAt) {
			continue
		}
		d.FetchArticle(article)
	}
}

//Save cache
func (d *Downg) Save() {
	d.store.WiteYuqueCache()
}

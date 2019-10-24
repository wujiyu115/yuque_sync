package main

import (
	"fmt"
	"path/filepath"
	"time"
	"os"

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
	doc, err := d.yuque.GetArticle(item.Slug)
	if err != nil {
		L.Error(err)
		return
	}
	d.genPost(doc)
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

func (d *Downg) genPost(post yuqueg.DocDetail) {
	data := post.Data
	if len(data.Body) == 0 || len(data.Title) == 0 {
		L.Error("invalid post:", data.Slug)
		return
	}
	postPath, err := filepath.Abs(d.config.PostPath)
	if err != nil {
		L.Error(fmt.Sprintf("abs path err: slug:%s, postPath:%s", data.Slug, d.config.PostPath))
		return
	}
	title, errStr := ReflectStrVal(data, d.config.MdFormat)
	if errStr != nil || len(title) == 0 {
		L.Error(fmt.Sprintf("empty title: slug:%s,mdFormat:%s", data.Slug, d.config.MdFormat))
		return
	}
	fun := AdapterMap[d.config.Adapter]
	if fun == nil {
		L.Error(fmt.Sprintf("empty adapter: slug:%s, adapter:%s", data.Slug, d.config.Adapter))
		return
	}
	r, err1 := Call(fun, data)
	if err1 != nil {
		L.Error(fmt.Sprintf("call adapter.fail: slug:%s, adapter:%s", data.Slug, d.config.Adapter))
		return
	}
	text := r[0].String()
	postPath = filepath.Join(postPath, fmt.Sprintf("%s.md", title))
	fp, err := os.OpenFile(postPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		L.Error(err)
	}
	defer fp.Close()
	_, err := fp.Write(text)
	if err != nil {
		L.Error(err)
	}
	L.Info(postPath, title, text)
}

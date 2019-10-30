package deal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/wujiyu115/yuqueg"
	"github.com/wujiyu115/yuques/util"
)

//Store of doc
type Store struct {
	CachePath      string
	CachedArticles map[string]map[string]yuqueg.DocBookDetail
}

//NewStore of store
func NewStore(cachePath string) *Store {
	s := &Store{
		CachePath:      cachePath,
		CachedArticles: make(map[string]map[string]yuqueg.DocBookDetail),
	}
	err := s.readYuqueCache()
	if err != nil {
		L.Error(err)
	}
	for namespace, v := range s.CachedArticles {
		L.Info(fmt.Sprintf("namespace:%s, len:%d", namespace, len(v)))
	}
	return s
}

func (s *Store) readYuqueCache() error {
	cerr := util.CreateFile(s.CachePath)
	if cerr != nil {
		return cerr
	}
	jsonFile, err := os.Open(s.CachePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	if len(string(byteValue)) == 0 {
		return nil
	}
	return json.Unmarshal(byteValue, &s.CachedArticles)
}

func (s *Store) findArticle(namespace string, slug string) *yuqueg.DocBookDetail {
	nspace := s.CachedArticles[namespace]
	if nspace != nil {
		d := nspace[slug]
		return &d
	}
	return nil
}

//AddArticle of repo
func (s *Store) AddArticle(namespace string, slug string, item yuqueg.DocBookDetail) {
	// s.CachedArticles = make(map[string]map[string]yuqueg.DocBookDetail)
	nspace := s.CachedArticles[namespace]
	if nspace == nil {
		s.CachedArticles[namespace] = make(map[string]yuqueg.DocBookDetail)
		nspace = s.CachedArticles[namespace]
	}
	nspace[slug] = item
}

//WiteYuqueCache cache
func (s *Store) WiteYuqueCache() error {
	jsonStr, err := json.Marshal(s.CachedArticles)
	if err != nil {
		return err
	}
	fp, err := os.OpenFile(s.CachePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer fp.Close()
	_, err = fp.Write(jsonStr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

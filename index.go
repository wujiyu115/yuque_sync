package main

//DoAction do
func DoAction() {
	cfg := LoadConfig()
	yuque := NewDowng(cfg)
	yuque.fetchArticles()
	yuque.Save()
}

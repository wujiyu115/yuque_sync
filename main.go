package main

func main() {
	cfg := LoadConfig()
	yuque := NewDowng(cfg)
	yuque.fetchArticles()
	yuque.Save()
}

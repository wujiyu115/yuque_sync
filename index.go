package main

var (
	//Cfg base config
	Cfg SyncConfig
)

//DoAction do
func DoAction() {
	Cfg = LoadConfig()
	yuque := NewDowng(Cfg)
	yuque.fetchArticles()
	yuque.Save()
}

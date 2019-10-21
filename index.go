package main

var (
	// base config
	Cfg SyncConfig
)

func DoAction() {
	Cfg = LoadConfig()
	readYuqueCache(Cfg.CachePath)
}

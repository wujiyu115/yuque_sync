package main

import (
	"github.com/wujiyu115/yuques/deal"
)

func main() {
	cfg := deal.LoadConfig()
	y := deal.NewDowng(cfg)
	y.DoSync()
}

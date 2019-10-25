package main

import (
	"github.com/wujiyu115/yuqueg"
)

func GenMarkDown(post yuqueg.Doc) string{
	return FormatRaw(post.Body)
}

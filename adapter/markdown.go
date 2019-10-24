package main

import (
	"github.com/wujiyu115/yuqueg"
)

AdapterMap.Set("markdown", func (post yuqueg.Doc){
	return FormatRaw(post.Body)
})
package adapter

import (
	"github.com/wujiyu115/yuqueg"
	"github.com/wujiyu115/yuques/util"
)

//GenMarkDown of gen
func GenMarkDown(post yuqueg.Doc) string {
	return util.FormatRaw(post.Body)
}

package adapter

import (
	"bytes"
	"fmt"
	"html"
	"regexp"

	"github.com/wujiyu115/yuqueg"
	"github.com/wujiyu115/yuques/parser"
	metadecoders "github.com/wujiyu115/yuques/parser/metadecoders"
)

const (
	//MetaRegStr regex o meta data
	MetaRegStr = "(title:|layout:|tags:|date:|categories:){1}(\\S|\\s)+?---"
	//PostTemplate of post
	// 	PostTemplate = `
	// ---
	// %s
	// ---
	// %s
	// `
	PostTemplate = "\n---\n%s\n---\n%s"
)

func parseMatter(meta string) string {
	d := metadecoders.Default
	m, _ := d.UnmarshalStringTo(meta, make(map[string]interface{}))
	var buf bytes.Buffer

	err := parser.InterfaceToConfig(m, metadecoders.YAML, &buf)
	if err != nil {
		yuqueg.L.Error("unexpected error value: %v", err)
		return ""
	}
	return buf.String()
}

//GenHexo of gen
func GenHexo(post yuqueg.Doc) string {
	body := html.UnescapeString(post.Body)
	metaReg, _ := regexp.Compile(MetaRegStr)
	meta := metaReg.FindString(body)
	if meta != "" {
		metaFormat := parseMatter(meta)
		leftBody := metaReg.ReplaceAllString(body, "")
		return fmt.Sprintf(PostTemplate, metaFormat, leftBody)
	}
	return GenMarkDown(post)
}

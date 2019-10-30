package adapter

import (
	"html"
	"regexp"

	"github.com/wujiyu115/yuqueg"
)

func parseMatter(body string) {

}

//GenHexo of gen
func GenHexo(post yuqueg.Doc) string {
	body := html.UnescapeString(post.Body)
	brTag, _ := regexp.Compile("(title:|layout:|tags:|date:|categories:){1}(\\S|\\s)+?---")
	return brTag.ReplaceAllString(body, "\n")
}

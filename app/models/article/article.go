package article

import (
	"goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	ID    int64
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(a.ID, 10))
}

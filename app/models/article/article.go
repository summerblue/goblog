// Package article 应用的文章模型
package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}

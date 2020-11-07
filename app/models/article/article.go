package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

// GetStringID 获取 ID 的字符串格式
func (a Article) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

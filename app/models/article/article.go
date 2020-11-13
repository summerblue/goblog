package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`

	// UserID uint64 `gorm:"null;index"`
	// User   user.User
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

// GetStringID 获取 ID 的字符串格式
func (a Article) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

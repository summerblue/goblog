// Package category 存放应用的分类数据模型
package category

import (
	"goblog/app/models"
)

// Category 文章分类
type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

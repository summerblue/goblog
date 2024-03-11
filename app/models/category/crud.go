package category

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

// Create 创建分类，通过 category.ID 来判断是否创建成功
func (category *Category) Create() (err error) {
	if err = model.DB.Create(&category).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// All 获取分类数据
func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

package article

import (
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"

	"github.com/vcraescu/go-paginator"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAll 获取全部文章
func GetAll() (pagination.ViewData, error) {

	var articles []Article
	query := model.DB.Model(Article{})
	_paginator := paginator.New(pagination.NewGORMAdapter(query), 2)
	_paginator.SetPage(7)

	url := route.Name2URL("articles.index")
	view := pagination.NewViewData(_paginator, url)

	pages, _ := view.Pages()
	fmt.Printf("pages %v \n", pages) // [2 3 4 5 6 7 8 9 10 11]

	next := view.Next()
	fmt.Printf("next %v \n", next.URL) // 8

	prev, _ := view.Prev()
	fmt.Printf("prev %v \n", prev) // 6

	current, _ := view.Current()
	fmt.Printf("current %v \n", current) // 7

	_paginator.Results(&articles)

	// var articles []Article
	// if err := model.DB.Preload("User").Find(&articles).Error; err != nil {
	// 	return articles, err
	// }
	return view, nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// GetByUserID 获取全部文章
func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

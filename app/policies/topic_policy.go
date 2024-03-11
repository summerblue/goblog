// Package policies 存放应用的授权策略
package policies

import (
	"goblog/app/models/article"
	"goblog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}

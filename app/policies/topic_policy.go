package policies

import (
	"goblog/pkg/auth"

	"goblog/app/models/article"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}

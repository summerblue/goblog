package user

import (
	"goblog/pkg/password"

	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在保存和更新模型前调用
func (user *User) BeforeSave(tx *gorm.DB) (err error) {

	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}

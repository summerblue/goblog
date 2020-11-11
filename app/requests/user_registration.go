package requests

import (
	"goblog/app/models/user"

	"github.com/thedevsaddam/govalidator"
)

// Rules 表单认证规则
var Rules = govalidator.MapData{
	"name":             []string{"required", "alpha_num", "between:3,20"},
	"email":            []string{"required", "min:4", "max:30", "email"},
	"password":         []string{"required", "min:6"},
	"password_comfirm": []string{"required"},
}

// Messages 定制错误消息
var Messages = govalidator.MapData{
	"name": []string{
		"required:用户名为必填项",
		"alpha_num:格式错误，只允许数字和英文",
		"between:用户名长度需在 3~20 之间",
	},
	"email": []string{
		"required:Email 为必填项",
		"min:Email 长度需大于 4",
		"max:Email 长度需小于 30",
		"email:Email 格式不正确，请提供有效的邮箱地址",
	},
	"password": []string{
		"required:密码为必填项",
		"min:长度需大于 6",
	},
	"password_comfirm": []string{
		"required:确认密码框为必填项",
	},
}

// ValidateRegistrationForm 验证表单，返回 errs 长度等于零即通过
func ValidateRegistrationForm(data user.User) map[string][]string {

	// 1. 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         Rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      Messages,
	}

	// 2. 开始验证
	errs := govalidator.New(opts).ValidateStruct()

	// 3. 因 govalidator 不支持 password_comfirm 验证，我们自己写一个
	if data.Password != data.PasswordComfirm {
		errs["password_comfirm"] = append(errs["password_comfirm"], "两次输入密码不匹配！")
	}

	return errs
}

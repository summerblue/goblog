package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/session"
	"net/http"
)

// StartSession 强制标头返回 HTML 内容类型
func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. 启动会话
		session.StartSession(w, r)

		// 2. 初始化认证系统
		auth.InitAuth()

		// 3. 继续处理接下去的请求
		next.ServeHTTP(w, r)
	})
}

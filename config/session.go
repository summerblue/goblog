package config

import "goblog/pkg/config"

func init() {
	config.Add("session", config.StrMap{

		// 目前只支持 Cookie
		"default": config.Env("SESSION_DRIVER", "cookie"),

		// 会话的 Cookie 名称
		"session_name": config.Env("SESSION_NAME", "goblog-session"),
	})
}

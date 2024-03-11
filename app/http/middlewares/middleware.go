package middlewares

import "net/http"

// HTTPHandlerFunc 简写 —— func(http.ResponseWriter, *http.Request)
type HTTPHandlerFunc func(http.ResponseWriter, *http.Request)

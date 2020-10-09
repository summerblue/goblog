package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func MyRouter() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return r
}

func TestMyRouter(t *testing.T) {

	// 1. 布置 —— 把服务跑起来
	server := httptest.NewServer(MyRouter())
	defer server.Close()

	// 2. 请求 —— 模拟用户访问浏览器
	resp, err := http.Get(server.URL + "/health")

	// 3. 检测 —— 是否是我们需要的
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")
}

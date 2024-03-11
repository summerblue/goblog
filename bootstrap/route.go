// Package bootstrap 负责应用初始化相关工作，比如初始化路由。
package bootstrap

import (
	"goblog/pkg/route"
	"goblog/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}

package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"goblog/app/http/controllers"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 文章相关页面
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete")

	// 用户认证
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")
	// r.HandleFunc("/auth/login", auc.Login).Methods("GET").Name("auth.login")
	// r.HandleFunc("/auth/dologin", auc.DoLogin).Methods("POST").Name("auth.dologin")
	// r.HandleFunc("/auth/doregister", auc.DoRegister).Methods("POST").Name("auth.doregister")
	// r.HandleFunc("/auth/forget", auc.Forget).Methods("GET").Name("auth.forget")
	// r.HandleFunc("/auth/forget/verify", auc.Send).Methods("POST").Name("auth.verify")
	// r.HandleFunc("/auth/reset-password", auc.ResetPassword).Methods("GET").Name("auth.reset-password")
	// r.HandleFunc("/auth/reset-password/verify", auc.DoResetPassword).Methods("POST").Name("auth.do-reset-password")

	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 中间件：强制内容类型为 HTML
	// r.Use(middlewares.ForceHTML)
}

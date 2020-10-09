package bootstrap

import (
	"github.com/totoval/framework/database"
	"github.com/totoval/framework/http/middleware"
	"github.com/totoval/framework/logs"
	"github.com/totoval/framework/request"
	"github.com/totoval/framework/validator"

	"totoval/config"

	c "github.com/totoval/framework/config"
)

func Initialize() {
	config.Initialize()
	// sentry.Initialize()
	logs.Initialize()
	// zone.Initialize()
	// cache.Initialize()
	database.Initialize()
	// m.Initialize()
	// queue.Initialize()
	// jobs.Initialize()
	// events.Initialize()
	// listeners.Initialize()

	validator.UpgradeValidatorV8toV9()
}

func Middleware(r *request.Engine) {
	r.Use(middleware.RequestLogger())

	if c.GetString("app.env") == "production" {
		r.Use(middleware.Logger())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Locale())
}

package web

import (
	"crud-lab/internal/services"
	"crud-lab/pkg/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type AppContext struct {
	echo.Context
	svc *services.Services
}

func cast(c echo.Context) *AppContext {
	return c.(*AppContext)
}

type App struct {
	svc *services.Services
	cfg config.Config
}

func NewApp(cfg config.Config, svc *services.Services) *App {
	return &App{
		cfg: cfg,
		svc: svc,
	}
}

func (app *App) ctxMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &AppContext{Context: c, svc: app.svc}
		return next(cc)
	}
}

func (app *App) Start() error {
	e := echo.New()

	e.Use(app.ctxMiddleware)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	app.addRoutes(e)

	return e.Start(app.cfg.WebAddr)
}

func (app *App) addRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}

package http

import (
    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/labstack/gommon/log"
    "net/http"
)

func New() (app *echo.Echo) {
    app = echo.New()
    app.Validator = &CustomValidator{validator: validator.New()}
    app.Use(middleware.SecureWithConfig(middleware.SecureConfig{
        ContentSecurityPolicy: "default-src 'self'",
    }))
    //{
    //    limit := limitMiddleware.New()
    //    app.Use(limit.LimiterWithEcho)
    //}
    //app.Use(middleware.Logger())
    app.Logger.SetLevel(log.ERROR)
    app.GET("/health/check", func(ctx echo.Context) error {
        return ctx.String(http.StatusOK, "ok")
    })

    return
}

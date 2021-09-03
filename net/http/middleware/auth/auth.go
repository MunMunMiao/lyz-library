package auth

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "lyz/library/encode"
)

type Config struct {
    JWTSecret string
    JWTClaims jwt.Claims
}

type Auth struct {
    config *Config
}

func New(c *Config) *Auth {
    return &Auth{
        config: c,
    }
}

func (a *Auth) JWTAuth() echo.MiddlewareFunc {
    return middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: []byte(a.config.JWTSecret),
        Claims: a.config.JWTClaims,
        ErrorHandlerWithContext: func(err error, ctx echo.Context) error {
            return encode.JSON(ctx, encode.Forbidden, nil)
        },
    })
}

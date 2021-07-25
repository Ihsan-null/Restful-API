package web

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	ot "github.com/hb-go/echo-web/middleware/opentracing"
	"github.com/hb-go/echo-web/middleware/session"
	"github.com/hb-go/echo-web/module/auth"
)

var (
	ctxPool = sync.Pool{
		New: func() interface{} {
			return &Context{}
		},
	}
)

func NewContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := ctxPool.Get().(*Context)
			defer func() {
				ctx.reset()
				ctxPool.Put(ctx)
			}()

			ctx.Context = c
			return next(ctx)
		}
	}
}

type Context struct {
	echo.Context
}

func (c *Context) reset() {
	c.Context = nil
}

func (ctx *Context) Session() session.Session {
	return session.Default(ctx)
}

func (ctx *Context) Auth() auth.Auth {
	return auth.Default(ctx)
}

func (ctx *Context) OpenTracingSpan() opentracing.Span {
	return ot.Default(ctx)
}

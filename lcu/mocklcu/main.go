package main

import (
	"context"
	"github.com/LvBay/firebox/lcu/mocklcu/controller"
	"github.com/LvBay/gf/v2/frame/g"
	"github.com/LvBay/gf/v2/net/ghttp"
	"github.com/LvBay/gf/v2/os/gcmd"
	"github.com/LvBay/gf/v2/os/gctx"
)

func main() {
	Main := gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(MiddlewareHandlerResponse)
				group.Bind(
					controller.Service,
				)

			})
			s.Run()
			return nil
		},
	}
	Main.Run(gctx.New())
}

func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	res := r.GetHandlerResponse()
	r.Response.WriteJson(res)
}

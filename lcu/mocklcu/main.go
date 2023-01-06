package main

import (
	"context"
	"github.com/LvBay/firebox/lcu/mocklcu/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"log"
	"time"
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
			s.BindHandler("/ws", func(r *ghttp.Request) {
				ws, err := r.WebSocket()
				if err != nil {
					glog.Error(ctx, err)
					r.Exit()
				}
				list := []string{
					"ReadyCheck",  // 是否接受匹配
					"ChampSelect", // 选择英雄
					"GameStart",   // 选择英雄结束后
					"InProgress",  // 游戏对局开始后
					"PreEndOfGame",
				}
				for _, v := range list {
					log.Println("send:", v)
					if err = ws.WriteMessage(1, []byte(v)); err != nil {
						return
					}
					time.Sleep(5 * time.Second)
				}
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

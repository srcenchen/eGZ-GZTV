package cmd

import (
	"context"
	"eGZ-GZTV/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "start",
		Usage: "start",
		Brief: "start gztv server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			go service.StreamServer() // 启动流媒体直播服务器
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind()
			})
			s.Run()
			return nil
		},
	}
)

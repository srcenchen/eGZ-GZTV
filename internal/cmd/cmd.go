package cmd

import (
	"context"
	"eGZ-GZTV/internal/controller/live"
	"eGZ-GZTV/internal/controller/setting"
	"eGZ-GZTV/internal/controller/upload"
	"eGZ-GZTV/internal/controller/user"
	"eGZ-GZTV/internal/controller/video"
	"eGZ-GZTV/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
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

			api := s.Group("/api")

			api.Group("/user", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					user.NewV1(),
				)
			})

			api.Group("/setting", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					setting.NewV1(),
				)
			})

			api.Group("/video", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					video.NewV1(),
				)
			})

			api.Group("/live", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					live.NewV1(),
				)
			})

			api.Group("/upload", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					upload.NewV1(),
				)
			})

			s.Run()
			return nil
		},
	}
)

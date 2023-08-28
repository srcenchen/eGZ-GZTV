package controller

import (
	"eGZ-GZTV/internal/controller/live"
	"eGZ-GZTV/internal/controller/setting"
	"eGZ-GZTV/internal/controller/upload"
	"eGZ-GZTV/internal/controller/user"
	"eGZ-GZTV/internal/controller/video"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Router(s *ghttp.Server) {
	// 静态资源绑定
	s.SetIndexFolder(true)
	s.AddStaticPath("/resource", "./resource")
	s.SetServerRoot("public")
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.ServeFile("public/index.html")
	})
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

}

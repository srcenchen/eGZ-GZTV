package cmd

import (
	"context"
	"eGZ-GZTV/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"m7s.live/engine/v4"
	_ "m7s.live/plugin/hdl/v4"
	_ "m7s.live/plugin/llhls/v4"
	_ "m7s.live/plugin/record/v4"
	_ "m7s.live/plugin/rtmp/v4"
)

var (
	Main = gcmd.Command{
		Name:  "start",
		Usage: "start",
		Brief: "start gztv server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// go service.StreamServer() // 启动流媒体直播服务器
			go engine.Run(context.Background(), "conf/monibuka_config.yaml")
			s := g.Server()
			controller.Router(s)
			s.Run()
			return nil
		},
	}
)

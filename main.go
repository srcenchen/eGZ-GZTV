package main

import (
	"eGZ-GZTV/internal/cmd"
	_ "eGZ-GZTV/internal/packed"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// 设置默认配置文件目录
	_ = g.Cfg().GetAdapter().(*gcfg.AdapterFile).AddPath("manifest/conf")
	_ = g.Cfg().GetAdapter().(*gcfg.AdapterFile).AddPath("conf")
	cmd.Main.Run(gctx.GetInitCtx())
}

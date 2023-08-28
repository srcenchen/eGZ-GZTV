package main

import (
	_ "eGZ-GZTV/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"eGZ-GZTV/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

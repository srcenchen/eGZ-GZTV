// Copyright 2019, Chef.  All rights reserved.
// https://github.com/srcenchen/gztv
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package service

import (
	"flag"
	"fmt"
	"github.com/q191201771/naza/pkg/nazalog"
	"os"
	"path/filepath"

	"github.com/q191201771/lal/pkg/base"

	"github.com/q191201771/lal/pkg/logic"
	"github.com/q191201771/naza/pkg/bininfo"
)

func StreamServer() {
	defer nazalog.Sync()

	confFilename := parseFlag()
	logic.DefaultConfFilenameList = []string{
		filepath.FromSlash("./conf/lalserver.conf.json"),
		filepath.FromSlash("./manifest/conf/lalserver.conf.json"),
	}
	lals := logic.NewLalServer(func(option *logic.Option) {
		option.ConfFilename = confFilename
	})
	err := lals.RunLoop()
	nazalog.Infof("server manager done. err=%+v", err)
}

func parseFlag() string {
	binInfoFlag := flag.Bool("v", false, "show bin info")
	cf := flag.String("c", "", "specify conf file")
	p := flag.String("p", "", "specify current work directory")
	flag.Parse()

	if *binInfoFlag {
		_, _ = fmt.Fprint(os.Stderr, bininfo.StringifyMultiLine())
		_, _ = fmt.Fprintln(os.Stderr, base.LalFullInfo)
		os.Exit(0)
	}
	if *p != "" {
		os.Chdir(*p)
	}

	return *cf
}

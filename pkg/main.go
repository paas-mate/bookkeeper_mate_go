package main

import (
	"bookkeeper_mate_go/pkg/bk"
	"bookkeeper_mate_go/pkg/config"
	"bookkeeper_mate_go/pkg/path"
	"bookkeeper_mate_go/pkg/util"
	"fmt"
	"github.com/paas-mate/gutil"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	if config.ClusterInit {
		stdout, stderr, err := gutil.CallScript(path.BkInitScript)
		util.Logger().Error("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
		if err != nil {
			util.Logger().Error(fmt.Sprintf("bk server init failed %v", err))
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
	if config.RemoteMode {
		util.Logger().Info("Remote mode")
	} else {
		err := bk.Start()
		if err != nil {
			util.Logger().Error("start bk server failed ", zap.Error(err))
			os.Exit(1)
		}
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}

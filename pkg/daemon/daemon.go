package daemon

import (
	"flag"
	"os"
	"os/exec"

	"go.uber.org/zap"
)

func init() {
	d := flag.Bool("d", false, "run app as a daemon with -d=true.")
	flag.Parse()

	if *d {
		cmd := exec.Command(os.Args[0], flag.Args()...)
		if err := cmd.Start(); err != nil {
			zap.L().Error("start failed", zap.Error(err), zap.Any("data", os.Args[0]))
			os.Exit(1)
		}
		zap.L().Info("webhook running...", zap.Any("data", cmd.Process.Pid))
		os.Exit(0)
	}
}

package daemon

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func init() {
	d := flag.Bool("d", false, "run app as a daemon with -d=true.")
	flag.Parse()

	if *d {
		cmd := exec.Command(os.Args[0], flag.Args()...)
		if err := cmd.Start(); err != nil {
			log.Printf("start %s failed, error: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		log.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		os.Exit(0)
	}
}

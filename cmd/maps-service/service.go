package main

import (
	"git.itzana.me/strafesnet/maps-service/pkg/cmds"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)

	app := cmds.NewApp()
	app.Commands = []*cli.Command{
		cmds.NewServeCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("Failed to run app")
	}
}

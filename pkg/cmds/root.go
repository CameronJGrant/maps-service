package cmds

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

var (
	appName    = filepath.Base(os.Args[0])
	commonFlag = []cli.Flag{
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable debug logging",
			Value: false,
			Action: func(ctx *cli.Context, debug bool) error {
				log.Println("ran")
				if debug {
					log.SetLevel(log.DebugLevel)
					log.SetFormatter(&log.TextFormatter{})
					log.SetReportCaller(false)
				}
				return nil
			},
		},
	}
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "Maps API Service"
	app.Flags = commonFlag

	return app
}

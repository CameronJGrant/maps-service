package cmd

import (
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

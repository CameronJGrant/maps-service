package cmds

import (
	"fmt"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore/gormstore"
	"git.itzana.me/strafesnet/maps-service/pkg/service"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"net/http"
)

func NewServeCommand() *cli.Command {
	return &cli.Command{
		Name:   "serve",
		Usage:  "Run maps service",
		Action: serve,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "pg-host",
				Usage:    "Host of postgres database",
				EnvVars:  []string{"PG_HOST"},
				Required: true,
			},
			&cli.IntFlag{
				Name:     "pg-port",
				Usage:    "Port of postgres database",
				EnvVars:  []string{"PG_PORT"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pg-db",
				Usage:    "Name of database to connect to",
				EnvVars:  []string{"PG_DB"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pg-user",
				Usage:    "User to connect with",
				EnvVars:  []string{"PG_USER"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pg-password",
				Usage:    "Password to connect with",
				EnvVars:  []string{"PG_PASSWORD"},
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "migrate",
				Usage:   "Run database migrations",
				Value:   true,
				EnvVars: []string{"MIGRATE"},
			},
			&cli.IntFlag{
				Name:    "port",
				Usage:   "Port to listen on",
				Value:   8080,
				EnvVars: []string{"PORT"},
			},
		},
	}
}

func serve(ctx *cli.Context) error {
	db, err := gormstore.New(ctx)
	if err != nil {
		log.WithError(err).Fatal("failed to connect database")
	}

	srv, err := api.NewServer(&service.Service{DB: db}, api.WithPathPrefix("/v2"))
	if err != nil {
		log.WithError(err).Fatal("failed to initialize api server")
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", ctx.Int("port")), srv)
}

package cmds

import (
	"fmt"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore/gormstore"
	"git.itzana.me/strafesnet/maps-service/pkg/service"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"net/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"git.itzana.me/strafesnet/go-grpc/auth"
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
			&cli.StringFlag{
				Name:     "auth-rpc-host",
				Usage:    "Host of auth rpc",
				EnvVars:  []string{"AUTH_RPC_HOST"},
				Value: "auth-service:8090",
			},
		},
	}
}

func serve(ctx *cli.Context) error {
	db, err := gormstore.New(ctx)
	if err != nil {
		log.WithError(err).Fatal("failed to connect database")
	}
	svc := &service.Service{
		DB: db,
	}

	conn, err := grpc.Dial(ctx.String("auth-rpc-host"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	sec := service.SecurityHandler{
		Client: auth.NewAuthServiceClient(conn),
	}

	srv, err := api.NewServer(svc, sec, api.WithPathPrefix("/v1"))
	if err != nil {
		log.WithError(err).Fatal("failed to initialize api server")
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", ctx.Int("port")), srv)
}

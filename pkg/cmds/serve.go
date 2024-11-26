package cmd

import "github.com/urfave/cli/v2"

func NewRunCommand() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Run maps service",
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
		},
	}
}

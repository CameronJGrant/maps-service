package main

import (
	"git.itzana.me/strafesnet/data-service/internal/datastore/gormstore"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Logger
	debug := os.Getenv("DEBUG") == "true"
	if debug {
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{})
	} else {
		log.SetLevel(log.InfoLevel)
		log.SetFormatter(&log.JSONFormatter{})
		log.SetReportCaller(true)
	}

	log.Info("Starting server...")

	// Postgres

	if os.Getenv("CRON") == "true" {
		db, err := gormstore.New(false)
		if err != nil {
			log.WithField("error", err).Fatalln("database startup failed")
			return
		}

		cron(db, debug)
	} else {
		db, err := gormstore.New(true)
		if err != nil {
			log.WithField("error", err).Fatalln("database startup failed")
			return
		}

		rpc(db, debug)
	}
}

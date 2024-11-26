package gormstore

import (
	"fmt"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
	"git.itzana.me/strafesnet/utils/logger"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(ctx *cli.Context) (datastore.Datastore, error) {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s",
				ctx.String("pg-host"),
				ctx.String("pg-user"),
				ctx.String("pg-password"),
				ctx.String("pg-db"),
				ctx.Int("pg-port")),
		), &gorm.Config{
			Logger: logger.New(),
		})
	if err != nil {
		log.WithError(err).Errorln("failed to connect to database")
		return nil, err
	}

	if ctx.Bool("migrate") {
		if err := db.AutoMigrate(&model.Submission{}); err != nil {
			log.WithError(err).Errorln("database migration failed")
			return nil, err
		}
	}

	return &Gormstore{db}, nil
}

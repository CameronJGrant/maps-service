package gormstore

import (
	"fmt"
	"os"
	"time"

	"git.itzana.me/strafesnet/maps-service/internal/datastore"
	"git.itzana.me/strafesnet/maps-service/internal/model"
	"git.itzana.me/strafesnet/utils/logger"
	"github.com/eko/gocache/lib/v4/cache"
	gocache_store "github.com/eko/gocache/store/go_cache/v4"
	gocache "github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(migrate bool) (datastore.Datastore, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DB"), os.Getenv("PG_PORT"))), &gorm.Config{
		Logger: logger.New()})
	if err != nil {
		log.WithFields(log.Fields{
			"PG_USER": os.Getenv("PG_USER"),
			"PG_HOST": os.Getenv("PG_HOST"),
			"PG_PORT": os.Getenv("PG_PORT"),
			"PG_DB":   os.Getenv("PG_DB"),
			"error":   err,
		}).Errorln("failed to connect to database")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(25)

	if migrate {
		if err := db.AutoMigrate(&model.Time{}, &model.User{}, &model.Bot{}, &model.Map{}, &model.Event{}, &model.Server{}, &model.Transaction{}, &model.Rank{}, &model.RankCalc{}); err != nil {
			log.WithField("error", err).Errorln("database migration failed")
			return nil, err
		}
	}

	return &Gormstore{db, cache.New[[]byte](gocache_store.NewGoCache(gocache.New(5*time.Minute, 10*time.Minute)))}, nil
}

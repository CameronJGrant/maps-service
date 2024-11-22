package main

import (
	"context"
	"time"

	"git.itzana.me/strafesnet/data-service/internal/datastore"
	"git.itzana.me/strafesnet/data-service/internal/model"
	log "github.com/sirupsen/logrus"
)

func cron(db datastore.Datastore, debug bool) {
	log.Infoln("Starting cron...")

	// Cleanup old events
	if err := db.Events().Clean(context.Background()); err != nil {
		log.WithField("error", err).Errorln("cron: failed to clean events")
	}

	// Cleanup dead servers
	if err := db.Servers().DeleteByLastUpdated(context.Background(), time.Now().Add(-time.Minute*10)); err != nil {
		log.WithField("error", err).Errorln("cron: failed to clean dead servers")
	}

	// Calculate ranks
	pairList, err := db.Times().DistinctStylePairs(context.Background())
	if err != nil {
		log.WithField("error", err).Errorln("cron: failed to update ranks")
	} else {
		for i := 0; i < len(pairList); i++ {
			if err := db.Ranks().UpdateAll(context.Background(), pairList[i].StyleID, pairList[i].GameID, pairList[i].ModeID); err != nil {
				log.WithField("error", err).Errorln("cron: failed to update ranks")
			}
		}

		if err := db.Ranks().UpdateRankCalc(context.Background()); err != nil {
			log.WithError(err).Errorln("cron: failed to update rank calc table")
		}

		db.Events().Create(context.Background(), model.Event{
			Event: "Ranks.Updated",
			Date:  time.Now(),
		})
	}

	log.Infoln("Cron completed!")
}

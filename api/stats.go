// Copyright (C) 2024 Pagefault Games - All Rights Reserved
// https://github.com/pagefaultgames

package api

import (
	"log"
	"time"

	"github.com/pagefaultgames/rogueserver/db"
	"github.com/robfig/cron/v3"
)

var (
	scheduler           = cron.New(cron.WithLocation(time.UTC))
	playerCount         int
	battleCount         int
	classicSessionCount int
)

func scheduleStatRefresh() {
	scheduler.AddFunc("@every 30s", func() {
		err := updateStats()
		if err != nil {
			log.Printf("failed to update stats: %s", err)
		}
	})

	scheduler.Start()
}

func updateStats() error {
	var err error
	playerCount, err = db.FetchPlayerCount()
	if err != nil {
		return err
	}

	battleCount, err = db.FetchBattleCount()
	if err != nil {
		return err
	}

	classicSessionCount, err = db.FetchClassicSessionCount()
	if err != nil {
		return err
	}

	return nil
}

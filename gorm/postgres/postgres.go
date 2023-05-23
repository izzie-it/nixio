package postgres

import (
	"time"

	"github.com/izzie-it/nixio/log"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connectionString string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	for i := 1; i <= 10; i++ {
		log.Infof("gorm postgres connection try %d\n", i)
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}

	log.Info("gorm postgres connected")

	return db, err
}

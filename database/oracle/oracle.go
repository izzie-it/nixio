package oracle

import (
	"database/sql"
	"time"

	"github.com/izzie-it/nixio/log"
	go_ora "github.com/sijms/go-ora/v2"
	"go.elastic.co/apm/module/apmsql/v2"
)

func Connect(databaseString string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	apmsql.Register("oracle", &go_ora.OracleDriver{})

	for i := 1; i <= 10; i++ {
		log.Infof("oracle connection try %d\n", i)
		db, err = apmsql.Open("oracle", databaseString)
		if err != nil {
			return nil, err
		}

		err = db.Ping()

		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		return nil, err
	}

	go func() {
		for {
			if err := db.Ping(); err != nil {
				db, _ = sql.Open("oracle", databaseString)
			}
			time.Sleep(time.Second)
		}
	}()

	log.Info("oracle connected")

	return db, nil
}

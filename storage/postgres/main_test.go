package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/husanmusa/to-do-service/config"
	"github.com/husanmusa/to-do-service/pkg/db"
	"github.com/husanmusa/to-do-service/pkg/logger"
)

var pgRepo *taskRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	pgRepo = NewTaskRepo(connDB)

	os.Exit(m.Run())
}

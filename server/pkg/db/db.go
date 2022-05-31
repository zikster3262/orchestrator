package db

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/zikster3262/orchestrator-server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init() Handler {

	url := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Err(err)
	}

	log.Info().Msg("Database connection established.")
	db.AutoMigrate(&models.Worker{})

	return Handler{db}
}

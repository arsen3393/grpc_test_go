package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"team00_01/internal/config"
	"team00_01/pkg/model"
)

func MustConnectDB(cfg *config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Ошибка к подключению к базе", err)
		return nil
	}
	fmt.Println("Успешное подключение к базе")

	err = db.AutoMigrate(&model.Anomalies{})
	if err != nil {
		log.Panic("Ошибка миграции", err)
		return nil
	}
	return db
}

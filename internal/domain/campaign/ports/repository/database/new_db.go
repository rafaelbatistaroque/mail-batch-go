package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := `host=localhost
			user=batista_dev
			password=postgres.kronos.67
			dbname=campaign_db
			port=5432
			sslmode=disable`

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}

	return db
}

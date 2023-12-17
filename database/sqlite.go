package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresClient struct {
	Client *gorm.DB
}

func New(dsn string) *PostgresClient {
	postgresConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err.Error())
	}
	return &PostgresClient{Client: postgresConn}
}

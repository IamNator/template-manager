package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLiteClient struct {
	Client *gorm.DB
}

func New(user, pass, host, port, dbName string) *SQLiteClient {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbName,
	)

	sqliteConn, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err.Error())
	}
	return &SQLiteClient{Client: sqliteConn}
}

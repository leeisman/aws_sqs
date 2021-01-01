package db

import (
	"gorm.io/gorm"
	configs "sqs/config"
)

type DB struct {
	Writer *gorm.DB
	Reader *gorm.DB
}

func NewDB(config *configs.Config) *DB {
	db := &DB{
		Writer: NewMysqlDB(config),
	}
	return db
}

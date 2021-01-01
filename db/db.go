package db

import (
	configs "aws_sqs/config"
	"gorm.io/gorm"
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

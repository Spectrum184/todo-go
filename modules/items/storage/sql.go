package storage

import "gorm.io/gorm"

type SqlStorage struct {
	db *gorm.DB
}

func NewSqlStorage(db *gorm.DB) *SqlStorage {

	return &SqlStorage{db: db}
}

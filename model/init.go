package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

type Config struct {
	Username string
	Password string
	Addr     string
	DbName   string
}

var DB *Database

func (d *Database) Init(conf *Config) {
	DB = &Database{
		d.openDB(conf),
	}
}

func (d *Database) openDB(conf *Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		conf.Username,
		conf.Password,
		conf.Addr,
		conf.DbName,
		true,
		"Local",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	return db
}

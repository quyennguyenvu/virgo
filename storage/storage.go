package storage

import (
	"sync"
	"virgo/config"
	"virgo/helper"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var once sync.Once

// Connect database
func Connect() {
	once.Do(func() {
		conf := config.GetConnection()
		var err error
		db, err = gorm.Open(conf.Driver, conf.DataSource)
		if err != nil {
			helper.Logging("Storage", "Connect", err.Error())
		}
		db.LogMode(conf.LogMode)
	})
}

// Disconnect ..
func Disconnect() {
	db.Close()
}

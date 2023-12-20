package drivers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitGormDB(dataSourceName string) error {
	var err error
	db, err = gorm.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

func CloseGormDB() error {
	return db.Close()
}

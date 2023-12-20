package drivers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection() (*gorm.DB, error) {
	// Replace with your database connection details
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

package utils

import (
	"time"

	"github.com/anggitrestuu/go-rest-api/internal/config"
	"github.com/anggitrestuu/go-rest-api/internal/constants"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/drivers"
	"gorm.io/gorm"
)

func SetupGORMPostgresConnection() (*gorm.DB, error) {
	var dsn string
	switch config.AppConfig.Environment {
	case constants.EnvironmentDevelopment:
		dsn = config.AppConfig.DBPostgreDsn
	case constants.EnvironmentProduction:
		dsn = config.AppConfig.DBPostgreURL
	}

	// Setup GORM config for PostgreSQL
	gormConfig := drivers.GORMConfig{
		DriverName:     config.AppConfig.DBPostgreDriver,
		DataSourceName: dsn,
		MaxOpenConns:   100,
		MaxIdleConns:   10,
		MaxLifetime:    15 * time.Minute,
	}

	// Initialize PostgreSQL connection with GORM
	conn, err := gormConfig.InitializeGORMDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

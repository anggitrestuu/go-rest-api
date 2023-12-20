package drivers

import (
	"fmt"
	"time"

	"github.com/anggitrestuu/go-rest-api/internal/constants"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/records"
	"github.com/anggitrestuu/go-rest-api/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GORMConfig holds the configuration for the GORM database instance
type GORMConfig struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	MaxLifetime    time.Duration
}

// InitializeGORMDatabase returns a new GORM DB instance
func (config *GORMConfig) InitializeGORMDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database with GORM: %v", err)
	}

	err = db.AutoMigrate(&records.Users{})
	if err != nil {
		return nil, fmt.Errorf("error migrating database with GORM: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting DB from GORM: %v", err)
	}

	// Set maximum number of open connections to database
	logger.Info(fmt.Sprintf("Setting maximum number of open connections to %d", config.MaxOpenConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	// Set maximum number of idle connections in the pool
	logger.Info(fmt.Sprintf("Setting maximum number of idle connections to %d", config.MaxIdleConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)

	// Set maximum time to wait for new connection
	logger.Info(fmt.Sprintf("Setting maximum lifetime for a connection to %s", config.MaxLifetime), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	sqlDB.SetConnMaxLifetime(config.MaxLifetime)

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database with GORM: %v", err)
	}

	return db, nil
}

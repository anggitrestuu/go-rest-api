package drivers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anggitrestuu/go-rest-api/internal/constants"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/records"
	"github.com/anggitrestuu/go-rest-api/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	logGorm "gorm.io/gorm/logger"
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
	newLogger := logGorm.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logGorm.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logGorm.Error, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(config.DataSourceName), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("error opening database with GORM: %v", err)
	}

	err = db.AutoMigrate(&records.Users{}, &records.Roles{}, &records.Authorizations{}, &records.RoleAuthorizations{}, &records.Accounts{}, &records.Products{})
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

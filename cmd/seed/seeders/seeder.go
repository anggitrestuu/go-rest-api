package seeders

import (
	"errors"
	"time"

	"github.com/anggitrestuu/go-rest-api/internal/constants"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/records"
	"github.com/anggitrestuu/go-rest-api/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Seeder interface {
	UserSeeder(userData []records.Users) (err error)
}

type seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) Seeder {
	return &seeder{db: db}
}

func (s *seeder) UserSeeder(userData []records.Users) (err error) {
	if len(userData) == 0 {
		return errors.New("users data is empty")
	}

	logger.Info("inserting users data...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	for _, user := range userData {
		user.CreatedAt = time.Now().In(constants.GMT7)
		// Using GORM's Create method to insert user record
		if err := s.db.Create(&user).Error; err != nil {
			return err
		}
	}
	logger.Info("users data inserted successfully", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})

	return
}

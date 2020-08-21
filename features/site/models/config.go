package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lucasasoaresmar/features-go/adapters/helpers"
)

// Config model
type Config struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ConfigRepository ...
type ConfigRepository struct{}

// Migrate site configuration
func (cr *ConfigRepository) Migrate() (err error) {
	helpers.ExecGORM(func(db *gorm.DB) {
		err = db.AutoMigrate(&Config{}).Error
		if err != nil {
			return
		}

		err = db.Create(&Config{Name: "My", Description: "Desc"}).Error
		if err != nil {
			return
		}
	})

	return err
}

// Get site configuration
func (cr *ConfigRepository) Get() (config Config, err error) {
	helpers.ExecGORM(func(db *gorm.DB) {
		err = db.First(&config).Limit(1).Error
		if err != nil {
			return
		}
	})

	return config, err
}

// Update site configuration
func (cr *ConfigRepository) Update(configChanges *Config) (updatedConfig Config, err error) {

	helpers.ExecGORM(func(db *gorm.DB) {
		err = db.First(&updatedConfig).Limit(1).Error
		if err != nil {
			return
		}

		err = db.Model(&updatedConfig).Updates(configChanges).Error
		if err != nil {
			return
		}
	})

	return updatedConfig, err
}

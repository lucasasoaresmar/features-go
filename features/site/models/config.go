package models

import (
	"github.com/jinzhu/gorm"
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
	execute(func(db *gorm.DB) {
		db.AutoMigrate(&Config{})
		db.Create(&Config{Name: "My", Description: "Desc"})
	})

	return nil
}

// Get site configuration
func (cr *ConfigRepository) Get() (config Config, err error) {
	execute(func(db *gorm.DB) {
		db.First(&config, 1)
	})

	return config, nil
}

// Update site configuration
func (cr *ConfigRepository) Update(configChanges *Config) (updatedConfig Config, err error) {

	execute(func(db *gorm.DB) {
		db.First(&updatedConfig, 1)
		db.Model(&updatedConfig).Updates(configChanges)
	})

	return updatedConfig, nil
}

package services

import (
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"gitlab.com/carriot-team/nominatim-to-elastic/configs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var CarDBProvider DBProvider

type DBProvider struct {
	config configs.DatabaseConfiguration

	mutex sync.Mutex
	DB    *gorm.DB
}

func CreateDBProvider(config configs.DatabaseConfiguration) (DBProvider, error) {
	provider := &DBProvider{
		config: config,
	}

	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s", config.Host, config.Port, config.User, config.Password, config.DB))

	if err != nil {
		log.Info("Error in Create DB: ", err)
		return *provider, err
	}
	db.DB().SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	provider.DB = db
	return *provider, nil
}

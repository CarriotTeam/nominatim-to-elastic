package services

import (
	"fmt"
	"sort"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"gitlab.com/carriot-team/nominatim-to-elastic/configs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DbProvider DBProvider

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

type OSMData struct {
	PlaceId int64  `sql:"column:place_id"`
	OSMId   int64  `sql:"column:osm_id"`
	OSMType string `sql:"column:osm_type"`
}

type GlobalDataType struct {
	Data []OSMData
	M    sync.Mutex
}

var GlobalData GlobalDataType

func Pop() ([]OSMData, bool) {
	GlobalData.M.Lock()
	defer GlobalData.M.Unlock()
	lenData := len(GlobalData.Data)
	if lenData < 1 {
		return []OSMData{}, false
	} else if lenData < configs.Config.System.DataPerRequest {
		element := GlobalData.Data
		GlobalData.Data = nil
		return element, true
	} else {
		index := lenData - (configs.Config.System.DataPerRequest)
		last := index + configs.Config.System.DataPerRequest
		element := (GlobalData.Data)[index:last]
		GlobalData.Data = (GlobalData.Data)[:index]
		return element, true
	}
}

func CreateGlobalData() {
	temp := []OSMData{}
	DbProvider.DB.Table("placex").Select("place_id ,osm_id , osm_type").Limit(300000).Find(&temp)
	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].PlaceId < temp[j].PlaceId
	})
	GlobalData = GlobalDataType{Data: temp}
}

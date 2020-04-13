package services

import (
	"fmt"
	"github.com/CarriotTeam/nominatim-to-elastic/configs"
	"github.com/CarriotTeam/nominatim-to-elastic/src/utlis"
	"log"
	"sort"
	"sync"
	"time"

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
		log.Fatal("Error in Create DB: ", err)
		return *provider, err
	}
	db.DB().SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	provider.DB = db
	return *provider, nil
}






func Pop() ([]utlis.OSMData, bool) {
	utlis.GlobalData.M.Lock()
	defer utlis.GlobalData.M.Unlock()
	lenData := len(utlis.GlobalData.Data)
	if lenData < 1 {
		return []utlis.OSMData{}, false
	} else if lenData < configs.Config.System.DataPerRequest {
		element := utlis.GlobalData.Data
		utlis.GlobalData.Data = nil
		return element, true
	} else {
		index := lenData - (configs.Config.System.DataPerRequest)
		last := index + configs.Config.System.DataPerRequest
		element := (utlis.GlobalData.Data)[index:last]
		utlis.GlobalData.Data = (utlis.GlobalData.Data)[:index]
		return element, true
	}
}

func CreateGlobalData() {
	temp := []utlis.OSMData{}
	log.Println("Start `placex` table query.")
	DbProvider.DB.Table("placex").Select("place_id ,osm_id , osm_type").Find(&temp)
	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].PlaceId < temp[j].PlaceId
	})
	utlis.GlobalData = utlis.GlobalDataType{Data: temp}
}





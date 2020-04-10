package app

import (
	"gitlab.com/carriot-team/datamodel/model"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/services"
	"log"
)

func OnNewLog(deviceLog *model.DeviceLog) {
	services.MQTTDevices.WithLabelValues(deviceLog.DeviceID).Inc()
	insert(*deviceLog)
}

func insert(temp model.DeviceLog) {
	err := services.CarDBProvider.DB.Create(&temp)
	if err.Error != nil {
		log.Println(err.Error)
	}
}

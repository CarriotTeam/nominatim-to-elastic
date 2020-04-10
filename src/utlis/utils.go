package utlis

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"gitlab.com/carriot-team/datamodel/model"
)

type rawData struct {
	Data []float64 `json:"data"`
}

func DecodeToDeviceLog(deviceLog []byte, deviceId string) *model.DeviceLog {
	temp := &rawData{}
	var message = string(deviceLog)
	var isEeprom = false
	if strings.Contains(message, "{eeprom:true}") {
		message = strings.Split(message, "{eeprom:true}")[0]
		isEeprom = true
	}
	if strings.Contains(message, "{ee}") {
		message = strings.Split(message, "{ee}")[0]
		isEeprom = true
	}
	err := json.Unmarshal([]byte(message), temp)
	if err != nil {
		log.Println(err)
		return nil
	}
	result := temp.jsonToLog(deviceId)
	result.Eeprom = isEeprom
	return result
}

func (l *rawData) jsonToLog(deviceId string) *model.DeviceLog {
	layout := "2006-01-02 15:04:05"
	stringTime := fmt.Sprintf("%f", l.Data[4])
	stringTime =
		stringTime[:4] + // year
			"-" +
			stringTime[4:6] + // month
			"-" +
			stringTime[6:8] + // day
			" " +
			stringTime[8:10] + // hour
			":" +
			stringTime[10:12] + // min
			":" +
			stringTime[12:14] // sec
	deviceTime, err := time.Parse(layout, stringTime)
	if err != nil {
		log.Println(err)
		return nil
	}
	var latitude = l.Data[1]
	var longitude = l.Data[2]
	if float64(int(l.Data[1]/100)) > 0 {
		latitude = float64(int(l.Data[1]/100)) + float64(float64(int(l.Data[1])%100)+l.Data[1]-float64(int(l.Data[1])))/60.0
		longitude = float64(int(l.Data[2]/100)) + float64(float64(int(l.Data[2])%100)+l.Data[2]-float64(int(l.Data[2])))/60.0
	}
	response := &model.DeviceLog{
		ServerTime:     time.Now(),
		DeviceID:       deviceId,
		DeviceTime:     deviceTime,
		Latitude:       latitude,
		Longitude:      longitude,
		Altitude:       l.Data[3],
		Satellites:     int(l.Data[6]),
		Course:         l.Data[8],
		SpeedOTG:       float32(l.Data[7]),
		AccelerationX1: float32(l.Data[9]),
		AccelerationY1: float32(l.Data[10]),
		AccelerationZ1: float32(l.Data[11]),
		AccelerationX2: float32(l.Data[12]),
		AccelerationY2: float32(l.Data[13]),
		AccelerationZ2: float32(l.Data[14]),
		Battery:        l.Data[15] != 0,
		Signal:         int(l.Data[16]),
		PowerSupply:    int(l.Data[17]),
	}
	if len(l.Data) > 18 {
		response.CarStatus = (l.Data[18]) != 0
		response.Humidity = float32(l.Data[19])
		response.Temp = float32(l.Data[20])
	}
	return response
}

package dataTest

import (
	"fmt"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/app"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/utlis"
	"time"
)

func Do() {

	i := 1
	s := getData()
	for {
		time.Sleep(time.Second * 5)
		if i%2 == 0 {
			app.OnNewLog(utlis.DecodeToDeviceLog([]byte(s), "865067024752126"))
		} else {
			app.OnNewLog(utlis.DecodeToDeviceLog([]byte(s), "865067024755384"))
		}
		i++
	}

}

func getData() string {
	timeTemp := fmt.Sprintf("%s%s%s%s%s%s", fmt.Sprint(time.Now().Year()), "05", "06", "07", "06", "23")
	data := "[" +
		"0," +
		"29.621465," + //Longitude
		"52.540548," + //Latitude
		"1538.8," + //Altitude
		timeTemp + "," + // Time
		"5," +
		"6," +
		"5678," + // SpeedOTG
		"8," +
		"56788888887," + //AccelerationX1
		"10," + //AccelerationY1
		"11," + //AccelerationZ1
		"12," + //AccelerationX2
		"13," + //AccelerationY2
		"14," + //AccelerationZ2
		"0," + //Battery
		"12," +
		"123," +
		"4," +
		"12," +
		"0.1" +
		"]"
	data = fmt.Sprintf("{\"data\":%s}", data)
	return data
}

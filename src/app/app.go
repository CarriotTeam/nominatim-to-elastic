package app

import (
	"gitlab.com/carriot-team/nominatim-to-elastic/src/services"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/utlis"
	"log"
	"time"
)

func App() {
	start := time.Now()
	services.CreateGlobalData()
	log.Printf("Create Global Data in %f second  \n ", time.Since(start).Seconds())
	go utlis.CalcTime(len(services.GlobalData.Data))
	utlis.ServeWorkers()
	utlis.Timer.Lock.Lock()
	t := utlis.Timer
	utlis.Timer.Lock.Unlock()
	log.Printf(" %d data in %f second ", t.Count, t.TimeStamp.Sub(utlis.StartTime).Seconds())
}

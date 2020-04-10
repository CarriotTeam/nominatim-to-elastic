package services

import (
	"gitlab.com/carriot-team/nominatim-to-elastic/configs"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func ServeLogger() {
	// set location of log file
	var logpath = configs.Config.Logger.Path
	f, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	Log = log.New(f, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}

package services

import (
	"github.com/CarriotTeam/nominatim-to-elastic/configs"
	"github.com/CarriotTeam/nominatim-to-elastic/src/utlis"
	"log"
	"sync"
)

var sW sync.WaitGroup

func ServeWorkers() {
	log.Println("Start Create.")
	for i := 0; i < configs.Config.System.Threads; i++ {
		sW.Add(1)
		go worker()
	}
	sW.Wait()
}

func worker() {
	defer sW.Done()
	for {
		osmData, ok := Pop()
		if !ok {
			break
		}
		data, err := utlis.SendRequest(utlis.StringBuilder(osmData))
		if err != nil {
			for _, x := range osmData {
				utlis.ErrHandel(utlis.Err{
					Err:  err,
					Data: x,
				})
			}
		}
		cnt :=ElasticProvider.Insert(data)
		utlis.UpdateTimer(cnt)
	}
}

package services

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/carriot-team/nominatim-to-elastic/configs"
	"log"
	"net/http"
)

func ServeMonitor() {
	port := configs.Config.Monitor.Port
	fmt.Println("prometheus agent port : " + port)
	prometheus.MustRegister(StartTime)
	prometheus.MustRegister(ConnectionStatus)
	prometheus.MustRegister(MQTTDevices)
	StartTime.SetToCurrentTime()
	http.Handle(configs.Config.Monitor.Url, promhttp.Handler())
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("can not start monitor server, err : %s", err.Error())
	}
}

var (
	StartTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fisher_start_time",
		Help: "app start time in second",
	})
	ConnectionStatus = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fisher_MQTT_connection",
		Help: "fisher to MQTT connection status",
	})
	MQTTDevices = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fisher_MQTT_devices",
			Help: "count logs per device",
		},
		[]string{"device_id"},
	)
)

package utlis

import (
	"encoding/json"
	"fmt"
	"gitlab.com/carriot-team/nominatim-to-elastic/configs"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/services"
	"log"
	"net/http"
	"strings"
	"sync"
)

var sW sync.WaitGroup

func ServeWorkers() {
	log.Println("Start Crawling.")
	for i := 0; i < configs.Config.System.Threads; i++ {
		sW.Add(1)
		go worker()
	}
	sW.Wait()
}

type NOMINATIMData struct {
	PlaceID     string `json:"place_id"`
	OsmType     string `json:"osm_type"`
	OsmID       string `json:"osm_id"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
	Class       string `json:"class"`
	Type        string `json:"type"`
	Importance  string `json:"importance"`
	Address     struct {
		Address29   string `json:"address29"`
		Pedestrian  string `json:"pedestrian"`
		Suburb      string `json:"suburb"`
		Road        string `json:"road"`
		Town        string `json:"town"`
		City        string `json:"city"`
		Village     string `json:"village"`
		Hamlet      string `json:"hamlet"`
		County      string `json:"county"`
		State       string `json:"state"`
		Postcode    string `json:"postcode"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
	} `json:"address"`
}



func worker() {
	defer sW.Done()
	for {
		osmData, ok := services.Pop()
		if !ok {
			break
		}
		data, err := sendRequest(stringBuilder(osmData))
		if err != nil {
			errHandel(err, osmData)
		}
		UpdateTimer(len(data))
	}
}

func errHandel(err error, data []services.OSMData) {
	log.Println(err)
}

func getBaseUrl() string {
	return configs.Config.System.Url + "/lookup?format=json&addressdetails=1&accept-language=" + configs.Config.System.Lng + "&osm_ids="
}

func sendRequest(address string) ([]NOMINATIMData, error) {
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	temp := []NOMINATIMData{}
	// do not handel
	json.NewDecoder(resp.Body).Decode(&temp)
	return temp, nil
}

func stringBuilder(osmData []services.OSMData) string {
	var sb strings.Builder
	for _, str := range osmData {
		sb.WriteString(fmt.Sprintf("%s%d,", str.OSMType, str.OSMId))
	}
	osmIds := sb.String()[:sb.Len()-1]
	return getBaseUrl() + osmIds
}

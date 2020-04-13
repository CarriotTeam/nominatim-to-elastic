package utlis

import (
	"encoding/json"
	"fmt"
	"gitlab.com/carriot-team/nominatim-to-elastic/configs"
	"net/http"
	"strings"
)

func ErrHandel(err Err) {
	ErrLop[err.Data.PlaceId]++
	if ErrLop[err.Data.PlaceId] < 4 {
		// push again
		GlobalData.M.Lock()
		GlobalData.Data = append(GlobalData.Data, err.Data)
		GlobalData.M.Unlock()
	} else {
		fmt.Printf("err : %d \n ", err.Data.PlaceId)
	}
}

func getBaseUrl() string {
	return configs.Config.System.Url + "/lookup?format=json&addressdetails=1&accept-language=" + configs.Config.System.Lng + "&osm_ids="
}

func SendRequest(address string) ([]NOMINATIMData, error) {
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

func StringBuilder(osmData []OSMData) string {
	var sb strings.Builder
	for _, str := range osmData {
		sb.WriteString(fmt.Sprintf("%s%d,", str.OSMType, str.OSMId))
	}
	osmIds := sb.String()[:sb.Len()-1]
	return getBaseUrl() + osmIds
}

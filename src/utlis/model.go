package utlis

import "sync"

type OSMData struct {
	PlaceId int  `sql:"column:place_id"`
	OSMId   int  `sql:"column:osm_id"`
	OSMType string `sql:"column:osm_type"`
}

type GlobalDataType struct {
	Data []OSMData
	M    sync.Mutex
}
var ErrLop = make(map[int]int)
type Err struct {
	Err  error
	Data OSMData
}


var GlobalData GlobalDataType

type NOMINATIMData struct {
	PlaceID     int `json:"place_id"`
	OsmType     string `json:"osm_type"`
	OsmID       int `json:"osm_id"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
	Class       string `json:"class"`
	Type        string `json:"type"`
	Importance  float64 `json:"importance"`
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

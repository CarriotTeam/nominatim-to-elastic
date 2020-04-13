package services

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/CarriotTeam/nominatim-to-elastic/configs"
	"github.com/CarriotTeam/nominatim-to-elastic/src/utlis"
	"log"
	"sync"
)

var ElasticProvider EsProvider

type EsProvider struct {
	config configs.Elastic
	Client *elastic.Client
}

func CreateElasticProvider(config configs.Elastic) (EsProvider, error) {
	provider := &EsProvider{
		config: config,
	}

	es, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%s", config.Host, config.Port)))
	if err != nil {
		log.Println(err)
		return *provider, err
	}
	provider.Client = es
	return *provider, nil
}

type tempCountType struct {
	count int
	mux   sync.Mutex
}

func (s *EsProvider) Insert(data []utlis.NOMINATIMData) int {
	tempCount := tempCountType{}
	bulkRequest := s.Client.Bulk()
	for _, Data := range data {
		index2Req := elastic.NewBulkIndexRequest().Index(s.config.Topic).Id(fmt.Sprintf("%d", Data.OsmID)).Doc(Data)
		bulkRequest = bulkRequest.Add(index2Req)
	}
	bulkResponse, err := bulkRequest.Do(context.TODO())
	if err != nil || bulkResponse == nil {
		for _, x := range data {
			utlis.ErrHandel(utlis.Err{
				Err: err,
				Data: utlis.OSMData{
					PlaceId: x.PlaceID,
					OSMId:   x.OsmID,
					OSMType: x.OsmType,
				},
			})
		}
	} else {
		ss := bulkResponse.Indexed()
		for _, j := range ss {
			if j.Status == 201 {
				tempCount.mux.Lock()
				tempCount.count++
				tempCount.mux.Unlock()
			}
		}
	}
	return tempCount.count
}

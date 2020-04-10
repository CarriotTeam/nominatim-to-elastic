package services

import (
	"github.com/elastic/go-elasticsearch/v7"
	"gitlab.com/carriot-team/nominatim-to-elastic/configs"
	"log"
)

var ElasticProvider EsProvider

type EsProvider struct {
	config configs.Elastic
	Client *elasticsearch.Client
}

func CreateElasticProvider(config configs.Elastic) (EsProvider, error) {
	provider := &EsProvider{
		config: config,
	}
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + provider.config.Host + ":" + provider.config.Port,
		},
	}
	var es, err = elasticsearch.NewClient(cfg)
	provider.Client = es
	if err != nil {
		log.Println(err)
		return *provider, err
	}
	return *provider, nil
}

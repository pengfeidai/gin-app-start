package es

import (
	"context"
	"gin-app-start/app/common"
	"gin-app-start/app/config"

	"github.com/olivere/elastic"
)

var EsClient *elastic.Client
var logger = common.Logger

// es Init
func Init() {
	config := config.Conf.Es
	client, err := elastic.NewClient(elastic.SetURL(config.Host))
	if err != nil {
		logger.Error("es connection error: ", err)
		panic(err)
	}
	//检查健康的状况，ping指定ip
	result, _, err := client.Ping(config.Host).Do(context.Background())
	if err != nil {
		logger.Error("es ping error: ", err)
		panic(err)
	}
	logger.Infof("es connection ping result:", result)
	EsClient = client
}

package noaa_client

import (
	"context"
	"log"
	"net/http"
	victoria "noaa-exporter/internal/victoria_metrics_client"
	"strings"
)

const KINDEX = "noaa-planetary-k-index.json"

func (nc *NoaaClient) GetDailyKIndex() *ResponseResult {
	log.Println("getting daily k index data")
	ctx := context.Background()
	result, err := nc.DoRequest(ctx, http.MethodGet, NOAA_PRODUCTS, KINDEX, nil)
	if err != nil {
		log.Fatal("some errors during request")
	}

	return result
}

func ScrapeKIndex(noaaClient *NoaaClient, vmetriClient *victoria.VMMetricsClient) {
	log.Println("loading daily k index data to vmetrics")
	dailyResponse := noaaClient.GetDailyKIndex()

	var responseBody [][]string
	if err := dailyResponse.ExtractResult(&responseBody); err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(responseBody); i++ {
		ctx := context.Background()
		csvData := strings.Join(responseBody[i], ",")

		if err := vmetriClient.SendCSVMetrics(ctx, victoria.NOAA_KINDEX_FORMAT, csvData); err != nil {
			log.Fatal(err)
		}
	}
}

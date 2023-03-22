package noaa_client

import (
	"context"
	"log"
	"net/http"
	victoria "noaa-exporter/internal/victoria_metrics_client"
	"strings"
)

const NOAA_PLASMA = "plasma-1-day.json"

func (nc *NoaaClient) GetPDailyPlasma() *ResponseResult {
	log.Println("getting daily data")
	ctx := context.Background()
	result, err := nc.DoRequest(ctx, http.MethodGet, NOAA_PLASMA, nil)
	if err != nil {
		log.Fatal("some errors during request")
	}

	return result
}

func ScrapePlasmaData(noaaClient *NoaaClient, vmetriClient *victoria.VMMetricsClient) {
	log.Println("scraping daily plasma data")
	dailyResponse := noaaClient.GetPDailyPlasma()

	var responseBody [][]string
	if err := dailyResponse.ExtractResult(&responseBody); err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(responseBody); i++ {
		ctx := context.Background()
		csvData := strings.Join(responseBody[i], ",")

		if err := vmetriClient.SendCSVMetrics(ctx, victoria.NOAA_PLASMA_FORMAT, csvData); err != nil {
			log.Fatal(err)
		}
	}
}

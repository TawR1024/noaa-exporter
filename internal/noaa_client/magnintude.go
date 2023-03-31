package noaa_client

import (
	"context"
	"log"
	"net/http"
	victoria "noaa-exporter/internal/victoria_metrics_client"
	"strings"
)

const DAILY_MAG = "mag-1-day.json"

func (nc *NoaaClient) GetDailyMagnitude() *ResponseResult {
	log.Println("getting daily data")
	ctx := context.Background()
	result, err := nc.DoRequest(ctx, http.MethodGet, NOAA_SOLAR_WIND, DAILY_MAG, nil)
	if err != nil {
		log.Fatal("some errors during request")
	}

	return result
}

func ScrapeMagnitudeData(noaaClient *NoaaClient, vmetriClient *victoria.VMMetricsClient) {
	log.Println("scraping daily magnitude data")
	dailyResponse := noaaClient.GetDailyMagnitude()

	var responseBody [][]string
	if err := dailyResponse.ExtractResult(&responseBody); err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(responseBody); i++ {
		ctx := context.Background()
		csvData := strings.Join(responseBody[i], ",")

		if err := vmetriClient.SendCSVMetrics(ctx, victoria.NOAA_MAGNITUDE_FORMAT, csvData); err != nil {
			log.Fatal(err)
		}
	}
}

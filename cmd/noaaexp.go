package main

import (
	"context"
	"log"
	"noaa-exporter/config"
	noaa "noaa-exporter/internal/noaa_client"
	victoria "noaa-exporter/internal/victoria_metrics_client"
	"strings"
	"time"
)

func main() {
	// init from config
	if err := config.ReadFromFile("config.yaml"); err != nil {
		log.Fatal(err)
	}

	noaaClient := noaa.NewNOAAClient(config.Config.NOAAURL)
	vmetriClient := victoria.NewVMMetricsClient(config.Config.VMMetricsURL)

	for { // get new data chunk ones a day
		log.Println("scraping daily data")
		dailyResponse := noaaClient.GetDailyMagnitude()

		var responseBody [][]string
		dailyResponse.ExtractResult(&responseBody)

		for i := 1; i < len(responseBody); i++ {
			ctx := context.Background()
			csvData := strings.Join(responseBody[i], ",")

			vmetriClient.SendMetrics(ctx, victoria.NOAA_DATA_FORMAT, csvData)
		}
		time.Sleep(24 * time.Hour)
	}
}

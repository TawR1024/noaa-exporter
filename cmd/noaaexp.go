package main

import (
	"log"
	"noaa-exporter/config"
	noaa "noaa-exporter/internal/noaa_client"
	victoria "noaa-exporter/internal/victoria_metrics_client"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// init from config
	if err := config.ReadFromFile("config.yaml"); err != nil {
		log.Fatal(err)
	}

	// prepare graceful shutdown
	systemInterrupt := make(chan os.Signal, 1)
	signal.Notify(systemInterrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(systemInterrupt)

	// register data scrapers
	scrapers := []func(noaaClient *noaa.NoaaClient, vmetriClient *victoria.VMMetricsClient){
		noaa.ScrapeMagnitudeData,
		noaa.ScrapePlasmaData,
		noaa.ScrapeKIndex,
	}

	stopChan := make(chan struct{}, 1)
	for _, scraper := range scrapers {
		go func(c chan struct{}, scraper func(noaaClient *noaa.NoaaClient, vmetriClient *victoria.VMMetricsClient)) { // run daily data scraper

			noaaClient := noaa.NewNOAAClient(config.Config.NOAAURL)
			vmetriClient := victoria.NewVMMetricsClient(config.Config.VMMetricsURL)

			// first time data scraping
			// used only on service start
			scraper(noaaClient, vmetriClient)
			for {
				select {

				case <-time.Tick(24 * time.Hour):
					scraper(noaaClient, vmetriClient)
				case <-c:
					log.Println("scraper stopped")
					return
				}
			}
		}(stopChan, scraper)
	}

	sig := <-systemInterrupt
	log.Println("got a signal", sig)
	log.Println("stopping goroutines")
	close(stopChan)
}

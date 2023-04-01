package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var Config *AppConfig

type AppConfig struct {
	NOAAURL           string `yaml:"noaa_url"`
	VMMetricsURL      string `yaml:"vm_metrics_url"`
	DataScrapePeriodH int    `yaml:"data_scape_period_hours"`
}

func ReadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := loadString(data); err != nil {
		return err
	}

	log.Printf("Config loaded from: %s", path)

	return nil
}

func loadString(data []byte) error {
	cfg := AppConfig{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}

	Config = &cfg

	return nil
}

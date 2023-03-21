package victoria_metrics_client

import (
	"net/http"
	"noaa-exporter/internal/httpclient"
)

type VMMetricsClient struct {
	HTTPClient *http.Client
	URL        string
}

func NewVMMetricsClient(url string) *VMMetricsClient {
	return &VMMetricsClient{
		HTTPClient: httpclient.NewHTTPClient(),
		URL:        url,
	}
}

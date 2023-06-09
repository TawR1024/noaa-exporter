package victoria_metrics_client

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/url"
)

const VMMETRICS_PATH = "/api/v1/import/csv?format="
const NOAA_MAGNITUDE_FORMAT = "1:time:custom:2006-01-02 15:04:05.000,2:metric:bx_gsm,3:metric:by_gsm,4:metric:bz_gsm,5:metric:lon_gsm,6:metric:lat_gsm,7:metric:bt"
const NOAA_PLASMA_FORMAT = "1:time:custom:2006-01-02 15:04:05.000,2:metric:density,3:metric:speed,4:metric:temperature"
const NOAA_KINDEX_FORMAT = "1:time:custom:2006-01-02 15:04:05.000,2:metric:Kp,3:metric:a_running,4:metric:station_count"

// SendCSVMetrics imports csv data to victoria-metrics.
func (vm *VMMetricsClient) SendCSVMetrics(ctx context.Context, format, data string) error {
	rawQery := url.QueryEscape(format)
	request, err := http.NewRequest(http.MethodPost, vm.URL+VMMETRICS_PATH+rawQery, bytes.NewBufferString(data))
	if err != nil {
		log.Fatal(err)
	}

	request = request.WithContext(ctx)
	_, err = vm.HTTPClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

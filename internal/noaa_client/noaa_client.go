package noaa_client

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"noaa-exporter/internal/httpclient"
)

const (
	NOAA_PATH   = "/products/solar-wind/"
	NOAA_DAILY  = "mag-1-day.json"
	NOAA_HOURLY = "mag-2-hour.json"
)

type NoaaClient struct {
	HTTPClient *http.Client
	URL        string
}

func NewNOAAClient(url string) *NoaaClient {
	return &NoaaClient{
		HTTPClient: httpclient.NewHTTPClient(),
		URL:        url,
	}
}

type ResponseResult struct {
	*http.Response
}

func (result *ResponseResult) ExtractResult(to interface{}) error {
	defer result.Body.Close()
	if err := json.NewDecoder(result.Body).Decode(&to); err != nil {
		log.Fatal("failed to unmarshal json", err)
		return err
	}

	return nil
}

func (nc *NoaaClient) DoRequest(ctx context.Context, method, timeRange string, body io.Reader) (*ResponseResult, error) {
	request, err := http.NewRequest(method, nc.URL+NOAA_PATH+timeRange, body)
	if err != nil {
		return nil, err
	}

	request = request.WithContext(ctx)
	response, _ := nc.HTTPClient.Do(request)
	responseResult := &ResponseResult{
		Response: response,
	}

	return responseResult, nil

}

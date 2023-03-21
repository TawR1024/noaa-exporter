package noaa_client

import (
	"context"
	"log"
	"net/http"
)

func (nc *NoaaClient) GetDailyMagnitude() *ResponseResult {
	log.Println("getting daily data")
	ctx := context.Background()
	result, err := nc.DoRequest(ctx, http.MethodGet, NOAA_DAILY, nil)
	if err != nil {
		log.Fatal("some errors during request")
	}

	return result
}

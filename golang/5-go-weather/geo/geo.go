package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	response, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	fmt.Println(response.StatusCode)
	if response.StatusCode != 200 {
		return nil, errors.New("Bad status code")
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

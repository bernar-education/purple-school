package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "Almaty"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3

	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	if !strings.Contains(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "Zero format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Almaty"
			geoData := geo.GeoData{
				City: expected,
			}
			format := 125

			_, err := weather.GetWeather(geoData, format)
			if err != weather.ErrorWrongFormat {
				t.Errorf("Expected %v, got %v", weather.ErrorWrongFormat, err)
			}
		})
	}
}

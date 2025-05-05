package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - preparing, expected result, data for testing function
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - execute function
	got, err := geo.GetMyLocation(city)

	// Assert - checking result with expected
	if err != nil {
		t.Error(err)
	}
	if got.City == expected.City {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	// Arrange - preparing, expected result, data for testing function
	city := "LondonTest"

	// Act - execute function
	_, err := geo.GetMyLocation(city)

	// Assert - checking result with expected
	if err != geo.ErrorNoCity {
		t.Errorf("Expected %v, got %v", geo.ErrorNoCity, err)
	}
}

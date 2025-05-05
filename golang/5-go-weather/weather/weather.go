package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrorWrongFormat = errors.New("ERROR_WRONG_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_HTTP")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_READBODY")
	}
	return string(body), nil
}

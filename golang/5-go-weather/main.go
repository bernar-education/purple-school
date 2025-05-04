package main

import (
	"demo/weather/geo"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("___ Weather API ___")
	city := flag.String("city", "", "person's city")
	// format := flag.Int("format", 1, "Output weather format")

	flag.Parse()

	fmt.Println(*city)
	// fmt.Println(*format)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(geoData)
}

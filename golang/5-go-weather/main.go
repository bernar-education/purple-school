package main

import (
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "person's city")
	format := flag.Int("format", 1, "Output weather format")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)
}

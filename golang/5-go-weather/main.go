package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	city := flag.String("city", "", "person's city")
	format := flag.Int("format", 1, "Output weather format")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)

	reader := strings.NewReader("Hello, I'm data stream")
	block := make([]byte, 4)
	for {
		_, err := reader.Read(block)
		if err == io.EOF {
			break
		}
		fmt.Printf("%q\n", block)
	}
}

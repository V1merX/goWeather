package main

import (
	"fmt"
	api "goWeather/api"
	"log"
)

func main() {
	city := scanCity()
	api.Start(city)
}

func scanCity() string {
	var city string

	_, err := fmt.Scan(&city)
	if err != nil {
		log.Fatal(err)
	}

	return city
}

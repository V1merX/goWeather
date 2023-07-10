package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type response struct {
	Temperature string          `json:"temperature"`
	Wind        string          `json:"wind"`
	Description string          `json:"description"`
	Forecasts   []forecastsData `json:"forecast"`
}

type forecastsData struct {
	Day         string `json:"day"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}

func Start(city string) {
	printData(getResponse(sendRequest(city)))
}

func sendRequest(city string) *http.Response {
	client := &http.Client{}

	resp, err := client.Get("https://goweather.herokuapp.com/weather/" + city)
	if err != nil {
		fmt.Println("We could`t send request")
		log.Fatal(err)
	}

	return resp
}

func getResponse(resp *http.Response) response {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r response

	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Invalid data request")
		log.Fatal(err)
	}

	return r
}

func printData(r response) {
	fmt.Println("")

	fmt.Println("Weather information:")
	fmt.Println("Temperature:", r.Temperature)
	fmt.Println("Wind:", r.Wind)
	fmt.Println("Description:", r.Description)

	for _, item := range r.Forecasts {
		fmt.Println(item.Day, "information")
		fmt.Println("Temperature:", item.Temperature)
		fmt.Println("Wind:", item.Wind)
	}
}

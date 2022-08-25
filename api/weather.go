package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetWeatherByDate(date time.Time) {
	apiKey := "e3c41a97e2a8d70669790181a34ff3e8"
	dateStr := date.Format("02-01-2006")
	fmt.Println(dateStr)
	url := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s&historical_date=%s&interval=24", apiKey, "Buenos Aires", dateStr)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error getting weather for %s: %v", dateStr, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
}

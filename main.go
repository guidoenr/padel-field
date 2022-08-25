package main

import (
	"github.com/guidoenr/padel-field/api"
	"time"
)

/*
	author: @github.com/guidoenr
	repo: github.com/guidoenr/padelfield
*/

func main() {
	// TODO, change port to heroku later?
	// Listen and Serve on 8080
	//ListenAndServe()
	//InitializeTurnos()
	//UpdateTurnos()
	date := time.Now().AddDate(0, 0, 3)
	api.GetWeatherByDate(date)
}

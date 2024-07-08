package main

import (
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/config"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/adapters/webserver"
)

func main() {

	/*
		Não consegui fazer pegar de forma dinâmica o local do arquivo .env
		Em tese ele sempre está no mesmo local do arquivo main.go
	*/
	cfg, err := config.LoadConfig("./cmd/fullcyclelab/")
	if err != nil {
		panic(err)
	}

	ws := webserver.NewWebServer(cfg.WebServerPort)
	ws.AddHandler("GET /weather/zipcode/{zipcode}", webserver.GetWeatherByZipcode)
	ws.Start()

}

package main

import (
	"os"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/adapters/webserver"
)

func main() {

	ws := webserver.NewWebServer(os.Getenv("WEB_SERVER_PORT"))
	ws.AddHandler("GET /weather/zipcode/{zipcode}", webserver.GetWeatherByZipcode)
	ws.Start()

}

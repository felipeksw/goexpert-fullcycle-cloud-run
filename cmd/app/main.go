package main

import (
	"log/slog"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/infra/webserver"
)

func main() {

	slog.SetLogLoggerLevel(slog.LevelDebug)

	ws := webserver.NewWebServer("8080")
	ws.AddHandler("GET /weather/zipcode/{zipcode}", webserver.GetWeatherByZipcodeHandler)
	err := ws.Start()
	if err != nil {
		slog.Error("could not start the webserver")
	}
}

package service

import (
	"encoding/json"
	"log/slog"
	"os"
	"time"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/adapter/webclient"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/dto"
)

func GetWeatherByCity(ad *dto.AddressResponseDto) (*dto.WeatherResponseDto, error) {

	wc, err := webclient.NewWebclient("GET", "https://api.weatherapi.com/v1/current.json?key="+os.Getenv("WEATHERAPI_KEY")+"&q="+ad.Localidade+"&aqi=no", 60*time.Second)
	if err != nil {
		slog.Error("[weatherapi webserver client]", "error", err.Error())
		return nil, err
	}

	var a dto.WeatherResponseDto

	err = wc.Do(func(p []byte) error {
		err = json.Unmarshal(p, &a)
		return err
	})

	slog.Debug("[struct]", "WeatherResponseDto", a)

	return &a, err
}

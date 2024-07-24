package usecase

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/dto"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/infra/webclient"
)

func NewWeatherByAddress(a dto.AddressDto, client *http.Client) (*dto.WeatherDto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var urlQuery = map[string]string{}
	urlQuery["key"] = os.Getenv("WEATHERAPI_KEY")
	urlQuery["q"] = a.Localidade
	urlQuery["aqi"] = "no"

	wc, err := webclient.NewWebclient(ctx, client, http.MethodGet, "https://api.weatherapi.com/v1/current.json", urlQuery)
	if err != nil {
		slog.Error("[weatherapi webserver client]", "error", err.Error())
		return nil, err
	}

	var w dto.WeatherDto

	err = wc.Do(func(p []byte) error {
		err = json.Unmarshal(p, &w)
		if err != nil {
			slog.Error("[weather body unmarshal]", "error", err.Error())
		}
		return err
	})
	if err != nil {
		slog.Error("[weather do]", "error", err.Error())
		return nil, err

	}

	slog.Debug("[struct]", "WeatherResponseDto", w)

	return &w, nil
}

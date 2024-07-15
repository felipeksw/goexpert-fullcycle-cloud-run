package weather

import (
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/adapters/webclient"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/entity"
)

type weatherRespDTO struct {
	Localidade string  `json:"localidade"`
	TempC      float32 `json:"temp_c"`
	TempF      float32 `json:"temp_f"`
	TempK      float32 `json:"temp_k"`
}

func GetWeatherByZipcode(zipcode string) (*weatherRespDTO, error) {

	z, err := entity.NewZipcode(zipcode)
	if err != nil {
		return nil, err
	}

	reqV, err := webclient.NewViacepRequest(*z)
	if err != nil {
		return nil, err
	}

	address, err := reqV.Do()
	if err != nil {
		return nil, err
	}

	reqW, err := webclient.NewWeatherRequest(address.Localidade)
	if err != nil {
		return nil, err
	}

	weather, err := reqW.Do()
	if err != nil {
		return nil, err
	}

	w := entity.NewCurrentWeather(weather.Current.Temp_c)

	return &weatherRespDTO{
		Localidade: address.Localidade,
		TempC:      w.TempC,
		TempF:      w.TempF,
		TempK:      w.TempK,
	}, nil
}

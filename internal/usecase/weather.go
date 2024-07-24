package usecase

import (
	"log/slog"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/dto"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/entity"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/service"
)

func GetWeatherByZipcode(zipcode string) (*dto.WeatherDto, error) {

	zipcodeEntity, err := entity.NewZipcode(zipcode)
	if err != nil {
		slog.Error("[NewZipcode]", "error", err.Error())
		return nil, err
	}

	var zipcodeDto = &dto.ZipcodeDto{
		Zipcode: zipcodeEntity.Zipcode(),
	}

	addressDto, err := service.GetAddressByZipcode(zipcodeDto)
	if err != nil {
		slog.Error("[GetAddressByZipcode]", "error", err.Error())
		return nil, err
	}

	weatherDto, err := service.GetWeatherByCity(addressDto)
	if err != nil {
		slog.Error("[GetWeatherByCity]", "error", err.Error())
		return nil, err
	}

	slog.Debug("[struct]", "weatherDto", *weatherDto)
	slog.Debug("[string]", "weatherDto.Current.TempC", weatherDto.Current.TempC)

	weatherEntity, err := entity.NewCurrentWeather(weatherDto.Current.TempC)
	if err != nil {
		slog.Error("[NewCurrentWeather]", "error", err.Error())
		return nil, err
	}

	slog.Debug("[struct]", "weatherEntity", *weatherEntity)

	return &dto.WeatherDto{
		Location: addressDto.Localidade,
		TempC:    weatherEntity.TempC(),
		TempF:    weatherEntity.TempF(),
		TempK:    weatherEntity.TempK(),
	}, nil
}

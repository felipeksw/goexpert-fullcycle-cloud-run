package entity

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/dto"
)

type localWeatherEntity struct {
	locale string
	tempC  float32
	tempF  float32
	tempK  float32
}

func (w *localWeatherEntity) TempC() float32 {
	return w.tempC
}

func (w *localWeatherEntity) Locale() string {
	return w.locale
}

func NewLocaleWeather(locale string, tempC float32) (*dto.LocalWeatherDto, error) {

	locale = strings.TrimSpace(locale)

	var tc = &localWeatherEntity{
		locale: locale,
		tempC:  tempC,
		tempF:  0,
		tempK:  0,
	}
	slog.Debug("struct", "localWeatherEntity", tc)

	err := tc.IsValid()
	if err != nil {
		slog.Error("[invalid locale]", "error", err.Error())
		return nil, err
	}

	return &dto.LocalWeatherDto{
		Locale: tc.locale,
		TempC:  tc.tempC,
		TempF:  tc.tempC*1.8 + 32,
		TempK:  tc.tempC + 273,
	}, nil
}

func (z *localWeatherEntity) IsValid() error {

	if len(z.Locale()) < 1 {
		return errors.New("location can not be empty")
	}

	if z.TempC() > 58 || z.tempC < -89 {
		return errors.New("temperature is outside the earth range")
	}
	return nil
}

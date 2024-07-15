package entity

import (
	"errors"
	"log/slog"
)

type weatherEntity struct {
	tempC float32
	tempF float32
	tempK float32
}

func (w *weatherEntity) TempC() float32 {
	return w.tempC
}

func (w *weatherEntity) TempF() float32 {
	return w.tempF
}

func (w *weatherEntity) TempK() float32 {
	return w.tempK
}

func NewCurrentWeather(tempC float32) (*weatherEntity, error) {

	var tc = &weatherEntity{
		tempC: tempC,
	}

	err := tc.IsValid()
	if err != nil {
		slog.Error("[invalid temperature]", "error", err.Error())
		return nil, err
	}

	return &weatherEntity{
		tempC: tempC,
		tempF: tempC*1.8 + 32,
		tempK: tempC + 273,
	}, nil
}

func (z *weatherEntity) IsValid() error {

	if z.TempC() > 58 || z.tempC < -89 {
		return errors.New("temperature is outside the earth range")
	}
	return nil
}

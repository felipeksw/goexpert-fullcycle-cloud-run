package entity_test

import (
	"testing"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/entity"
	"github.com/stretchr/testify/assert"
)

func FuzzNewCurrentWeather(f *testing.F) {

	seed := []float32{-100, -89, -20.05, 0, 15.17, 24.20, 58, 100}
	for _, temps := range seed {
		f.Add(temps)
	}

	f.Fuzz(func(t *testing.T, temp float32) {
		weather, err := entity.NewCurrentWeather(temp)

		if temp > 58 || temp < -89 {
			assert.Error(t, err, "temperature is outside the earth range")
		} else {
			assert.Equal(t, temp, weather.TempC())
			assert.Equal(t, temp*1.8+32, weather.TempF())
			assert.Equal(t, temp+273, weather.TempK())
		}
	})
}

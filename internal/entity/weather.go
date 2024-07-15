package entity

type CurrentWeather struct {
	TempC float32
	TempF float32
	TempK float32
}

func NewCurrentWeather(tempC float32) *CurrentWeather {

	return &CurrentWeather{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273,
	}
}

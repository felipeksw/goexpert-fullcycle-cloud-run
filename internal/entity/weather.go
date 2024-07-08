package entity

type CurrentWeather struct {
	TempC float32
	TempF float32
	TempK float32
}

func NewCurrentWeather(tempC float32) *CurrentWeather {

	var tempF = tempC*1.8 + 32
	var tempK = tempC + 273

	return &CurrentWeather{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}
}

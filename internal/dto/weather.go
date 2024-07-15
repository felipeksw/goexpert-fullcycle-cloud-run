package dto

type WeatherResponseDto struct {
	Current weatherResposeCurrentDto
}

type weatherResposeCurrentDto struct {
	TempC float32 `json:"temp_c"`
}

type WeatherDto struct {
	Location string
	TempC    float32
	TempF    float32
	TempK    float32
}

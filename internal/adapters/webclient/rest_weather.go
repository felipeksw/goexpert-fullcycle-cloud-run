package webclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type WeatherRespDTO struct {
	Current CurrentDTO
}

type CurrentDTO struct {
	Temp_c float32 `json:"temp_c"`
}

type WeatherApiRequest struct {
	Req http.Request
}

func NewWeatherRequest(location string) (*WeatherApiRequest, error) {

	// Não consegui acertar essa parte de contexto com timeout
	/*
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	*/

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://api.weatherapi.com/v1/current.json?key="+os.Getenv("WEATHERAPI_KEY")+"&q="+location+"&aqi=no", nil)
	if err != nil {
		return nil, err
	}

	//log.Println(*req)

	return &WeatherApiRequest{
		Req: *req,
	}, nil
}

func (r *WeatherApiRequest) Do() (*WeatherRespDTO, error) {

	client := http.DefaultClient
	resp, err := client.Do(&r.Req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//log.Println(*resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//log.Println(body)

	var a WeatherRespDTO
	err = json.Unmarshal(body, &a)
	if err != nil {
		return nil, err
	}

	//log.Println(a)

	return &a, nil
}

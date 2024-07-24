package service

import (
	"encoding/json"
	"log/slog"
	"time"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/adapter/webclient"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/dto"
)

func GetAddressByZipcode(zc *dto.ZipcodeDto) (*dto.AddressResponseDto, error) {

	wc, err := webclient.NewWebclient("GET", "https://viacep.com.br/ws/"+zc.Zipcode+"/json/", 60*time.Second)
	if err != nil {
		slog.Error("[viacep NewWebclient failed]", "error", err.Error())
		return nil, err
	}

	var a dto.AddressResponseDto

	err = wc.Do(func(p []byte) error {
		err = json.Unmarshal(p, &a)
		return err
	})

	return &a, err
}

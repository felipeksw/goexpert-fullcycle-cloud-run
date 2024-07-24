package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/dto"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/infra/webclient"
)

func NewAddressByZipcode(z dto.ZipcodeDto, client *http.Client) (*dto.AddressDto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wc, err := webclient.NewWebclient(ctx, client, http.MethodGet, "https://viacep.com.br/ws/"+z.Zipcode+"/json/", nil)
	if err != nil {
		slog.Error("[viacep NewWebclient failed]", "error", err.Error())
		return nil, err
	}

	var a dto.AddressDto

	err = wc.Do(func(p []byte) error {
		err = json.Unmarshal(p, &a)
		if err != nil {
			slog.Error("[zipcode body unmarshal]", "error", err.Error())
		}
		return err
	})
	if err != nil {
		slog.Error("[webclient do]", "error", err.Error())

	}
	slog.Debug("[zipcode body]", "body", a)

	if a.Error != "" {
		return nil, errors.New("zip code not found")
	}

	return &a, err
}

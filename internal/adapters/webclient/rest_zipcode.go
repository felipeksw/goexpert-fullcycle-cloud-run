package webclient

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/entity"
)

type addressRespDTO struct {
	Cep        string `json:"cep"`
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

type ViacepRequest struct {
	Req http.Request
}

func NewViacepRequest(cep entity.Zipcode) (*ViacepRequest, error) {

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://viacep.com.br/ws/"+cep.Zipcode+"/json/", nil)
	if err != nil {
		return nil, err
	}

	return &ViacepRequest{
		Req: *req,
	}, nil
}

func (r *ViacepRequest) Do() (*addressRespDTO, error) {

	client := http.DefaultClient
	resp, err := client.Do(&r.Req)
	if err != nil {
		return nil, err
	}

	//log.Println(*resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var a addressRespDTO
	err = json.Unmarshal(body, &a)
	if err != nil {
		return nil, err
	}

	if a.Erro == "true" {
		return nil, errors.New("zip code not found")
	}

	return &a, nil
}

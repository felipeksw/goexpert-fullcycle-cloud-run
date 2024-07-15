package webserver

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/usecase"
)

type AddressRespDTO struct {
	Cep        string `json:"cep"`
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

func GetWeatherByZipcodeHandler(w http.ResponseWriter, r *http.Request) {

	zipcode := r.PathValue("zipcode")

	slog.Info("[Handler start]", "zipcode", zipcode)

	wt, err := usecase.GetWeatherByZipcode(zipcode)
	if err != nil {
		switch err.Error() {
		case "zip code must be 8 numeric digits":
			w.WriteHeader(http.StatusUnprocessableEntity)
		case "zip code not found":
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}

	slog.Debug("[struct]", "GetWeatherByZipcode", *wt)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wt)

	slog.Info("[Handler end]", "json", *wt)
}

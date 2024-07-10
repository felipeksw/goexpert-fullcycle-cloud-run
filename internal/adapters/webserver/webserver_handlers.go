package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/usecase/weather"
)

func GetWeatherByZipcode(w http.ResponseWriter, r *http.Request) {

	zp := r.PathValue("zipcode")

	//log.Println(zp)

	resp, err := weather.GetWeatherByZipcode(zp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		if err.Error() == "zip code must be 8 numeric digits" {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		if err.Error() == "zip code not found" {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

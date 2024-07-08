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
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

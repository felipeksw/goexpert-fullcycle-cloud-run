package webclient_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/infra/mockup"
	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/infra/webclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ReqresResponse struct {
	data    ReqresResponseData
	support ReqresResponseSupport
}
type ReqresResponseData struct {
	id         string
	email      string
	first_name string
	last_name  string
	avatar     string
}
type ReqresResponseSupport struct {
	url  string
	text string
}

func TestNewWebclient(t *testing.T) {

	mocReqresResponseSuccessBody := `{"data":{"id":3,"email":"emma.wong@reqres.in","first_name":"Emma","last_name":"Wong","avatar":"https://reqres.in/img/faces/3-image.jpg"},"support":{"url":"https://reqres.in/#support-heading","text":"To keep ReqRes free, contributions towards server costs are appreciated!"}}`

	mockRoundTripper := new(mockup.MockRoundTripper)
	mockClient := &http.Client{Transport: mockRoundTripper}

	mockRoundTripper.On("RoundTrip", mock.Anything).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(mocReqresResponseSuccessBody))),
	}, nil)

	var urlQuery = map[string]string{}
	urlQuery["key01"] = "value01"
	urlQuery["key02"] = "value02"
	urlQuery["output"] = "json"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wc, err := webclient.NewWebclient(ctx, mockClient, http.MethodGet, "https://dummy.restapiexample.com/api/v1/employee/1", urlQuery)
	assert.Nil(t, err)

	var w ReqresResponse

	err = wc.Do(func(p []byte) error {
		err = json.Unmarshal(p, &w)
		assert.Nil(t, err)
		return err
	})
	assert.Nil(t, err)
	slog.Info("[test ReqresResponse]", "ReqresResponse", w)
}

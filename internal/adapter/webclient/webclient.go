package webclient

import (
	"io"
	"log/slog"
	"net/http"
	"time"
)

type webClient struct {
	request *http.Request
	client  *http.Client
}

func NewWebclient(method string, url string, timeout time.Duration) (*webClient, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		slog.Error("[http.NewRequest failed]", "error", err.Error())
		return nil, err
	}

	//ctx, cancel := context.WithTimeout(context.Background(), timeout)
	//defer cancel()
	//req = req.WithContext(ctx)

	return &webClient{
		request: req,
		client:  http.DefaultClient,
	}, nil
}

func (w *webClient) Do(ret func([]byte) error) error {

	resp, err := w.client.Do(w.request)
	if err != nil {
		slog.Error("[http.Client.Do failed]", "error", err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	defer func() {
		body = nil
	}()
	if err != nil {
		slog.Error("[io.ReadAll failed]", "error", err.Error())
		return err
	}

	return ret(body)
}

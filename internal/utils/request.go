package utils

import (
	"fmt"
	"io"
	"net/http"
)

var (
	ErrCreateRequest = fmt.Errorf("falha ao criar requisição")
	ErrTimeout       = fmt.Errorf("timeout na requisição")
	ErrRequest       = fmt.Errorf("falha durante a requisição")
	ErrReadBody      = fmt.Errorf("falha ao ler o corpo da resposta")
)

func RequestPageBody(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fail on request page: %w", err)
	}
	// TODO
	return res.Body, nil
}

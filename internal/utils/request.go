package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	ErrCreateRequest = fmt.Errorf("falha ao criar requisição")
	ErrTimeout       = fmt.Errorf("timeout na requisição")
	ErrRequest       = fmt.Errorf("falha durante a requisição")
	ErrReadBody      = fmt.Errorf("falha ao ler o corpo da resposta")
)

func RequestPageBody(url string) (string, error) {
	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrCreateRequest, err)
	}

	req.Header.Add("Accept", "text/html")

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", ErrTimeout
		}
		return "", fmt.Errorf("%w: %v", ErrRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code inesperado: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error on get request body")
		return "", fmt.Errorf("%w: %v", ErrReadBody, err)
	}

	return string(body), nil
}

package client

import (
	"net/http"
	"time"
)

type Client struct {
	apiUrl     string
	httpClient *http.Client
	token      string
}

// Создаем конструктор  вешаем на клиент таймаут
func NewClient(apiUrl string) *Client {
	client := &Client{
		httpClient: &http.Client{Timeout: 5 * time.Second},
		apiUrl:     apiUrl,
	}
	return client
}

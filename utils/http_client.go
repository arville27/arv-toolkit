package utils

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	client := &http.Client{Timeout: 15 * time.Second}
	return client
}

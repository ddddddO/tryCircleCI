package main

import (
	"log"

	"github.com/pkg/errors"
)

func main() {
	var (
		version = "0.0.1"
		mode    = "development"
	)
	client := NewClient(version, mode)

	endpoint := "http://success.com"
	status, err := client.Get(endpoint)
	if err != nil {
		log.Printf("status: %d\n", status)
		log.Fatalf("err: %v", err)
	}

	log.Printf("status: %d\n", status)
}

type Client struct {
	Version string
	Mode    string
}

func NewClient(v, m string) *Client {
	return &Client{
		Version: v,
		Mode:    m,
	}
}

func (c *Client) Get(endpoint string) (int, error) {
	switch endpoint {
	case "http://success.com":
		return 200, nil

	case "http://fail.com":
		return 500, errors.New("server error")

	default:
		return 400, errors.New("client error")

	}
}

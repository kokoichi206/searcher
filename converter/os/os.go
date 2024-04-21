package os

import (
	"crypto/tls"
	"net/http"

	"github.com/opensearch-project/opensearch-go/v2"
)

type client struct {
	osc *opensearch.Client
}

func New() (*client, error) {
	osc, err := opensearch.NewClient(
		opensearch.Config{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Addresses: []string{"https://localhost:9200"},
			Username:  "admin",
			Password:  "90@A1qwe#Pa_q",
		},
	)

	client := &client{
		osc: osc,
	}

	return client, err
}

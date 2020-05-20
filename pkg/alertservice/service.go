package alertservice

import (
	"log"

	"github.com/phamvinhdat/httpclient"
)

type Service interface {
	Alert(message string, uid uint) error
}

type service struct {
	alertURL string
	client   httpclient.Client
}

func (s *service) Alert(message string, uid uint) error {
	log.Println("alert message: ", message)
	return nil
}

func New(alertURL string, client httpclient.Client) Service {
	return &service{
		alertURL: alertURL,
		client:   client,
	}
}

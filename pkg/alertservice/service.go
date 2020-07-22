package alertservice

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/phamvinhdat/httpclient"
	"github.com/phamvinhdat/httpclient/body"
)

type alertStruct struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}

type Service interface {
	Alert(title, message string, uid uint) error
}

type service struct {
	alertURL string
	client   httpclient.Client
}

func (s *service) Alert(title, message string, uid uint) error {
	statusCode, err := s.client.Post(context.Background(), s.alertURL,
		httpclient.WithBodyProvider(body.NewJson(alertStruct{
			Title:   title,
			Message: message,
			UserID:  uid,
		})))
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return errors.New(
			fmt.Sprintf("call alert with status code %d", statusCode))
	}

	return nil
}

func New(alertURL string, client httpclient.Client) Service {
	return &service{
		alertURL: alertURL,
		client:   client,
	}
}

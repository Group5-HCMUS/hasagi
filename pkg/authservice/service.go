package authservice

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/phamvinhdat/httpclient"
	"github.com/phamvinhdat/httpclient/hook"
	"github.com/sirupsen/logrus"
)

const (
	Authorization = "Authorization"
)

type Service interface {
	VerifyToken(token string) (User, error)
}

type service struct {
	authURL string
	client  httpclient.Client
}

func (s *service) VerifyToken(token string) (user User, err error) {

	statusCode, err := s.client.Get(context.Background(), s.authURL,
		httpclient.WithHeader(Authorization, authToken(token)),
		httpclient.WithHookFn(hook.UnmarshalResponse(&user)),
	)
	if err != nil {
		logrus.Error("failed to verify token, error: ", err)
		return
	}

	if statusCode != http.StatusOK {
		msgErr := fmt.Sprintf("status code is %d", statusCode)
		logrus.Error(msgErr)
		err = errors.New(msgErr)
		return
	}

	if user.Parent == nil {
		if user.Child == nil {
			user.Role = Unknown
			return
		}

		user.Role = Child
	}

	user.Role = Parent
	return
}

func New(authURL string, client httpclient.Client) Service {
	return &service{
		authURL: authURL,
		client:  client,
	}
}

func authToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

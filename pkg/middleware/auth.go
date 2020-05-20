package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Group5-HCMUS/hasagi/pkg/authservice"
	"github.com/Group5-HCMUS/hasagi/pkg/model"
	"github.com/gin-gonic/gin"
)

const UserKey = "user"

func VerifyToken(authService authservice.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		strToken := c.GetHeader(authservice.Authorization)
		token := trimToken(strToken)
		user, err := authService.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.HttpResponse{
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set(UserKey, user)
		c.Next()
	}
}

func Role(role authservice.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetUser(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.HttpResponse{
				Message: err.Error(),
			})
			return
		}

		if user.Role != role {
			c.JSON(http.StatusUnauthorized, model.HttpResponse{
				Message: "permission denied",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUser(c *gin.Context) (user authservice.User, err error) {
	u, ok := c.Get(UserKey)
	if !ok {
		err = errors.New("cannot get user info")
		return
	}

	user, ok = u.(authservice.User)
	if !ok {
		err = errors.New("cannot get parse user info")
		return
	}

	return
}

func trimToken(strToken string) string {
	strs := strings.Split(strToken, " ")
	if len(strs) != 2 {
		return ""
	}
	return strs[1]
}

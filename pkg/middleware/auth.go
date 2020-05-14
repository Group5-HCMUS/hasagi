package middleware

import (
	"net/http"
	"strings"

	"github.com/Group5-HCMUS/hasagi/pkg/model"

	"github.com/Group5-HCMUS/hasagi/pkg/authservice"
	"github.com/gin-gonic/gin"
)

const userKey = "user"

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

		c.Set(userKey, user)
		c.Next()
	}
}

func Role(role authservice.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, ok := c.Get(userKey)
		if !ok {
			c.JSON(http.StatusInternalServerError, model.HttpResponse{
				Message: "cannot get user info",
			})
			c.Abort()
			return
		}

		user, ok := u.(authservice.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, model.HttpResponse{
				Message: "cannot get parse user info",
			})
			c.Abort()
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

func trimToken(strToken string) string {
	strs := strings.Split(strToken, " ")
	if len(strs) != 2 {
		return ""
	}
	return strs[1]
}

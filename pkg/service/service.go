package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Register(g gin.IRouter) {
	g.POST("/alert-location/:userID")
}

func (s *service) postAlertLocation(c *gin.Context) {
	userIDStr := c.Param("userID")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, HttpResponse{
			Message: "user id is nil",
		})
		c.Abort()
		return
	}

	userID, err := cast.ToUintE(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			Message: "user id is not a number",
		})
		c.Abort()
		return
	}

	req := CreateAlertLocationRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			Message: "invalid data",
		})
		c.Abort()
		return
	}

	err = s.repo.CreateAlertLocation(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, HttpResponse{
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, HttpResponse{
		Message: "create alert location successfully",
	})
}

package service

import (
	"net/http"

	"github.com/Group5-HCMUS/hasagi/pkg/authservice"
	"github.com/Group5-HCMUS/hasagi/pkg/middleware"

	"github.com/Group5-HCMUS/hasagi/pkg/model"
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
	// child role
	child := g.Use(middleware.Role(authservice.Child))

	// parent role
	parent := g.Use(middleware.Role(authservice.Parent))
	parent.POST("/alert-location/:userID")

}

func (s *service) postAlertLocation(c *gin.Context) {
	userIDStr := c.Param("userID")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, model.HttpResponse{
			Message: "user id is nil",
		})
		c.Abort()
		return
	}

	userID, err := cast.ToUintE(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse{
			Message: "user id is not a number",
		})
		c.Abort()
		return
	}

	req := CreateAlertLocationRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse{
			Message: "invalid data",
		})
		c.Abort()
		return
	}

	err = s.repo.CreateAlertLocation(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse{
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, model.HttpResponse{
		Message: "create alert location successfully",
	})
}

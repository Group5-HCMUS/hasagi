package service

import (
	"net/http"

	"github.com/Group5-HCMUS/hasagi/pkg/authservice"
	"github.com/Group5-HCMUS/hasagi/pkg/middleware"
	"github.com/Group5-HCMUS/hasagi/pkg/model"
	"github.com/gin-gonic/gin"
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
	child.POST("/location/history", s.postLocationHistory)

	// parent role
	parent := g.Use(middleware.Role(authservice.Parent))
	parent.POST("/location/alert", s.postAlertLocation)

}

func (s *service) postLocationHistory(c *gin.Context) {
	//createLcHistoryReq := CreateLocationHistoryRequest{}
	//err := c.BindJSON(&createLcHistoryReq)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, model.HttpResponse{
	//		Message: "invalid data",
	//	})
	//	return
	//}
	//
	//user, err := middleware.GetUser(c)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, model.HttpResponse{
	//		Message: err.Error(),
	//	})
	//	return
	//}
	//
	//createLcHistoryReq.UserID = user.ID
	//err := s.repo.CreateLocationHistoryAndAlert()
}

func (s *service) postAlertLocation(c *gin.Context) {
	alertLocationReq := CreateAlertLocationRequest{}
	err := c.BindJSON(&alertLocationReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse{
			Message: "invalid data",
		})
		c.Abort()
		return
	}

	err = s.repo.CreateAlertLocation(alertLocationReq)
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

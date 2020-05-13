package service

import (
	"github.com/Group5-HCMUS/hasagi/pkg/allocationrepo"
	"github.com/Group5-HCMUS/hasagi/pkg/lchistoryrepo"
)

type Repository interface {
	CreateAlertLocation(userID uint,
		createAlertLocationRequest CreateAlertLocationRequest) error
}

type repository struct {
	alLocationRepo allocationrepo.Repository
	lcHistoryRepo  lchistoryrepo.Repository
}

func (r *repository) CreateAlertLocation(userID uint,
	createAlertLocationRequest CreateAlertLocationRequest) error {
	alLocation := allocationrepo.AlertLocation{
		Longitude: createAlertLocationRequest.Longitude,
		Name:      createAlertLocationRequest.Name,
		Latitude:  createAlertLocationRequest.Latitude,
		UserID:    userID,
	}

	return r.alLocationRepo.Create(&alLocation)
}

func NewRepository(
	alLocationRepo allocationrepo.Repository,
	lcHistoryRepo lchistoryrepo.Repository,
) Repository {
	return &repository{
		alLocationRepo: alLocationRepo,
		lcHistoryRepo:  lcHistoryRepo,
	}
}

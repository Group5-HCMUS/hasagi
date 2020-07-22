package service

import (
	"fmt"
	"time"

	"github.com/Group5-HCMUS/hasagi/pkg/alertservice"
	"github.com/Group5-HCMUS/hasagi/pkg/allocationrepo"
	"github.com/Group5-HCMUS/hasagi/pkg/lchistoryrepo"
	"github.com/Group5-HCMUS/hasagi/pkg/utils"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	CreateAlertLocation(createAlertLocationRequest CreateAlertLocationRequest) error
	CreateLocationHistoryAndAlert(alertToUserID uint,
		createLocationHistoryRequest CreateLocationHistoryRequest) error
}

type repository struct {
	maxDistanceAlert float64       // meters
	maxTimeAlert     time.Duration // minute
	alertService     alertservice.Service
	alLocationRepo   allocationrepo.Repository
	lcHistoryRepo    lchistoryrepo.Repository
}

func (r *repository) CreateAlertLocation(
	createAlertLocationRequest CreateAlertLocationRequest) error {
	alLocation := allocationrepo.AlertLocation{
		Longitude: createAlertLocationRequest.Longitude,
		Name:      createAlertLocationRequest.Name,
		Latitude:  createAlertLocationRequest.Latitude,
		UserID:    createAlertLocationRequest.UserID,
	}

	return r.alLocationRepo.Create(&alLocation)
}

func (r *repository) CreateLocationHistoryAndAlert(alertToUserID uint,
	createLocationHistoryRequest CreateLocationHistoryRequest) error {
	lcHistory := lchistoryrepo.LocationHistory{
		Name:      createLocationHistoryRequest.Name,
		Longitude: createLocationHistoryRequest.Longitude,
		Latitude:  createLocationHistoryRequest.Latitude,
		UserID:    createLocationHistoryRequest.UserID,
		Timestamp: createLocationHistoryRequest.Timestamp,
	}

	err := r.lcHistoryRepo.Create(&lcHistory)
	if err != nil {
		logrus.Error("failed to create location history, error: ", err)
		return err
	}

	return r.alert(lcHistory.Timestamp, lcHistory.UserID, alertToUserID,
		lcHistory.Latitude, lcHistory.Longitude)
}

func (r *repository) getListAlertLocation(parentID, childID uint) ([]allocationrepo.AlertLocation, error) {
	return nil, nil
}

func NewRepository(
	maxDistanceAlert float64,
	maxTimeAlert time.Duration,
	alLocationRepo allocationrepo.Repository,
	lcHistoryRepo lchistoryrepo.Repository,
) Repository {
	return &repository{
		maxDistanceAlert: maxDistanceAlert,
		maxTimeAlert:     maxTimeAlert,
		alLocationRepo:   alLocationRepo,
		lcHistoryRepo:    lcHistoryRepo,
	}
}

func (r *repository) alert(timestamp time.Time, userID, alertToUserID uint,
	latitude, longitude float64) error {
	tFromNow := time.Now().Sub(timestamp)
	if tFromNow > r.maxTimeAlert {
		logrus.Info("timestamp sub now (minutes): ", tFromNow.Minutes())
		return nil
	}

	alLocations, err := r.alLocationRepo.GetByUserID(userID)
	if err != nil {
		logrus.Error("failed to get alert locations, error: ", err)
		return err
	}

	for _, alLocation := range alLocations {
		distance := utils.Distance(alLocation.Latitude, alLocation.Longitude,
			latitude, longitude)
		if distance <= r.maxDistanceAlert {
			msg := fmt.Sprintf("Your child has arrived %s", alLocation.Name)
			err = r.alertService.Alert("Child safe alert", msg, alertToUserID)
			if err != nil {
				logrus.Errorf("failed to alert message: %s, error: %v", msg,
					err)
				return err
			}
		}
	}

	return nil
}

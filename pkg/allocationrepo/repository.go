package allocationrepo

import "github.com/jinzhu/gorm"

type Repository interface {
	Create(alertLocation *AlertLocation) error
	GetByUserID(userID uint) ([]AlertLocation, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(alertLocation *AlertLocation) error {
	if err := r.db.Create(alertLocation).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByUserID(userID uint) ([]AlertLocation, error) {
	var alertLocations []AlertLocation
	if err := r.db.Find(alertLocations, "user_id = ?",
		userID).Error; err != nil {
		return nil, err
	}

	return alertLocations, nil
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

package lchistoryrepo

import "github.com/jinzhu/gorm"

type Repository interface {
	Create(alertLocation *LocationHistory) error
	GetByUserID(userID uint) ([]LocationHistory, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(locationHistory *LocationHistory) error {
	if err := r.db.Create(locationHistory).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByUserID(userID uint) ([]LocationHistory, error) {
	var locationHistories []LocationHistory
	if err := r.db.Order("timestamp desc").
		Find(locationHistories, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return locationHistories, nil
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

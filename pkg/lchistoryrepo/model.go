package lchistoryrepo

import (
	"time"

	"github.com/jinzhu/gorm"
)

type LocationHistory struct {
	gorm.Model
	Name      string    `json:"name"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	UserID    uint      `json:"userid"`
	Timestamp time.Time `json:"timestamp"`
}

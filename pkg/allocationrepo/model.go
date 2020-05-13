package allocationrepo

import "github.com/jinzhu/gorm"

type AlertLocation struct {
	gorm.Model
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	UserID    uint    `json:"userid"`
}

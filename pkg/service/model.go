package service

import "time"

type CreateAlertLocationRequest struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name"`
	UserID    uint    `json:"user_id, omitempty"`
}

type CreateLocationHistoryRequest struct {
	CreateAlertLocationRequest
	Timestamp time.Time `json:"timestamp"`
}

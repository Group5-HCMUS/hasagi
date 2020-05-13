package service

type HttpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateAlertLocationRequest struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name"`
}

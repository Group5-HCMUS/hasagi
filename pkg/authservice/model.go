package authservice

type UserRole uint8

const (
	Unknown UserRole = 0
	Parent  UserRole = 1
	Child   UserRole = 2
)

type User struct {
	ID     uint          `json:"id"`
	Role   UserRole      `json:"-"`
	Parent *AdditionInfo `json:"parent"`
	Child  *AdditionInfo `json:"child"`
}

type AdditionInfo struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
}

package model

import "time"

type User struct {
	ID       int64     `gorm:"primary_key" json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	AddTime  time.Time `json:"add_time"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	Gender   string    `json:"gender"`
}
type UserAndInfo struct {
	User    User    `json:"user"`
	Student Student `json:"student"`
	Teacher Teacher `json:"teacher"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

package model

import "time"

type Teacher struct {
	ID         int64     `gorm:"primary_key" json:"id"`
	UserID     int64     `json:"user_id"`
	AddTime    time.Time `json:"add_time"`
	EmployeeID string    `json:"employee_id"`
	Department string    `json:"department"`
	IDNumber   string    `json:"id_number"`
}

func (Teacher) TableName() string {
	return "teacher"
}

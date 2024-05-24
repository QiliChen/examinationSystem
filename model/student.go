package model

import "time"

type Student struct {
	ID        int64     `gorm:"primary_key" json:"id"`
	UserID    int64     `json:"user_id"`
	AddTime   time.Time `json:"add_time"`
	StudentID string    `json:"student_id"`
	Phone     string    `json:"phone"`
	IDNumber  string    `json:"id_number"`
	Address   string    `json:"address"`
}

func (Student) TableName() string {
	return "student"
}

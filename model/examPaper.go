package model

type ExamPaper struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	TeacherId   int64  `json:"teacher_id"`
	AddTime     string `json:"add_time"`
	Name        string `json:"name"`
	Duration    int    `json:"duration"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}

func (ExamPaper) TableName() string {
	return "exam_paper"
}

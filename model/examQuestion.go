package model

type ExamQuestion struct {
	ID            int64  `json:"id" gorm:"primaryKey"`
	AddTime       string `json:"add_time"`
	PaperID       int64  `json:"paper_id"`
	Question      string `json:"question"`
	Options       string `json:"options"`
	Score         int64  `json:"score"`
	CorrectAnswer string `json:"correct_answer"`
	Analysis      string `json:"analysis"`
	Type          int    `json:"type"`
}

func (ExamQuestion) TableName() string {
	return "exam_question"
}

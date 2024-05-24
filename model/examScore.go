package model

import "time"

type ExamScore struct {
	ID               int64     `json:"id" gorm:"primary_key"`
	StudentID        int64     `json:"student_id"`
	PaperID          int64     `json:"paper_id"`
	GradingTeacherID int64     `json:"grading_teacher_id"`
	Score            int       `json:"score"`
	Comments         string    `json:"comments"`
	AddTime          time.Time `json:"add_time"`
}

// GradeExamRequest 批改试卷请求结构体
type GradeExamRequest struct {
	StudentID        int64  `json:"student_id"`
	PaperID          int64  `json:"paper_id"`
	GradingTeacherID int64  `json:"grading_teacher_id"`
	Comments         string `json:"comments"`
	Score            int64  `json:"answers"`
	ScoreId          int64  `json:"score_id"`
}

// GradeExamResponse 批改试卷响应结构体
type GradeExamResponse struct {
	Score   int    `json:"score"`
	Message string `json:"message"`
}

// GetExamScoresByStudentIDResponse 根据学生ID获取成绩响应结构体
type GetExamScoresByStudentIDResponse struct {
	ID               int64     `json:"id" gorm:"primary_key"`
	StudentID        int64     `json:"student_id"`
	StudentName      string    `json:"student_name"`
	PaperID          int64     `json:"paper_id"`
	PaperName        string    `json:"paper_name"`
	GradingTeacherID int64     `json:"grading_teacher_id"`
	Score            int       `json:"score"`
	Comments         string    `json:"comments"`
	AddTime          time.Time `json:"add_time"`
}

func (ExamScore) TableName() string {
	return "exam_score"
}

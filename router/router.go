package router

import (
	"examination_system/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/exampaper", controller.GetExamPapers)
	r.GET("/exampaper/:id", controller.GetExamPaperByID)
	r.POST("/exampaper", controller.CreateExamPaper)
	r.PUT("/exampaper/:id", controller.UpdateExamPaper)
	r.DELETE("/exampaper/:id", controller.DeleteExamPaper)

	r.GET("/examquestion", controller.GetExamQuestions)
	r.GET("/examquestion/:id", controller.GetExamQuestionByID)
	r.POST("/examquestion", controller.CreateExamQuestion)
	r.PUT("/examquestion/:id", controller.UpdateExamQuestion)
	r.DELETE("/examquestion/:id", controller.DeleteExamquestion)
	r.GET("/examquestionByExamPaperID/:examPaperID", controller.GetExamQuestionByExamPaperID)

	r.GET("/student", controller.GetStudents)
	r.GET("/student/:id", controller.GetStudentByID)
	r.POST("/student", controller.CreateStudent)
	r.PUT("/student/:id", controller.UpdateStudent)
	r.DELETE("/student/:id", controller.DeleteStudent)

	r.GET("/teacher", controller.GetTeachers)
	r.GET("/teacher/:id", controller.GetTeacherByID)
	r.POST("/teacher", controller.CreateTeacher)
	r.PUT("/teacher/:id", controller.UpdateTeacher)
	r.DELETE("/teacher/:id", controller.DeleteTeacher)

	r.GET("/user", controller.GetUsers)
	r.GET("/user/:id", controller.GetUserByID)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)

	r.GET("/examscores", controller.GetExamScores)
	r.GET("/examscores/:id", controller.GetExamScoreByID)
	r.POST("/examscores", controller.CreateExamScore)
	r.PUT("/examscores/:id", controller.UpdateExamScore)
	r.DELETE("/examscores/:id", controller.DeleteExamScore)

	r.POST("/login", controller.Login)
	r.GET("/getExamPaperWithQuestions/:id", controller.GetExamPaperWithQuestions)
	r.POST("/gradeExam", controller.GradeExam)
	r.GET("/getExamScoresByStudentID/:student_id", controller.GetExamScoresByStudentID)
	r.GET("/getAllStudentScores", controller.GetAllStudentScores)
	r.POST("/createExamPaperAQuestion", controller.CreateExamPaperAQuestion)
	r.GET("/getAllUsersAndInfo", controller.GetAllUsersAndInfo)
	r.POST("/createUserWithType", controller.CreateUserWithType)
}

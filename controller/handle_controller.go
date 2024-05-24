package controller

import (
	"errors"
	"examination_system/database"
	"examination_system/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}

	var user model.User
	result := database.DB.Where("username = ?", loginRequest.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法登录: " + result.Error.Error()})
		}
		return
	}

	if user.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	if user.Role == "student" {
		var student model.Student
		result := database.DB.Where("user_id = ?", user.ID).First(&student)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法登录: " + result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user, "student": student})
	} else if user.Role == "teacher" {
		var teacher model.Teacher
		result := database.DB.Where("user_id = ?", user.ID).First(&teacher)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法登录: " + result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user, "teacher": teacher})
	} else if user.Role == "admin" {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无效的用户角色,请联系管理员!"})
	}
}

// GetExamPaperWithQuestions 获取试卷和相关问题
func GetExamPaperWithQuestions(c *gin.Context) {
	id := c.Param("id")
	var examPaper model.ExamPaper
	var examQuestions []model.ExamQuestion

	// 获取试卷信息
	if err := database.DB.First(&examPaper, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到试卷"})
		return
	}

	// 获取试卷相关的问题
	if err := database.DB.Where("paper_id = ?", id).Find(&examQuestions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试卷问题"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exam_paper": examPaper,
		"questions":  examQuestions,
	})
}

// GradeExam 批改试卷接口
func GradeExam(c *gin.Context) {
	var req model.GradeExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}

	// 查找现有的成绩记录
	var examScore model.ExamScore
	if err := database.DB.Where("id = ?", req.ScoreId).First(&examScore).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到相关成绩记录"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取成绩记录: " + err.Error()})
		}
		return
	}

	// 更新成绩记录
	examScore.GradingTeacherID = req.GradingTeacherID
	examScore.Comments = req.Comments
	examScore.AddTime = time.Now()

	if err := database.DB.Save(&examScore).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法保存成绩: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.GradeExamResponse{
		Score:   examScore.Score,
		Message: "批改完成",
	})
}

// GetExamScoresByStudentID 根据学生ID获取成绩
func GetExamScoresByStudentID(c *gin.Context) {
	studentIDStr := c.Param("student_id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的学生ID"})
		return
	}

	log.Println("Converted studentID:", studentID) // 输出转换后的 studentID 日志
	var examScores []model.ExamScore
	result := database.DB.Where("student_id = ?", studentID).Find(&examScores)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取成绩信息: " + result.Error.Error()})
		return
	}
	var getExamScoresByStudentIDResponses []model.GetExamScoresByStudentIDResponse
	for _, examScore := range examScores {
		var examPaper model.ExamPaper
		result := database.DB.First(&examPaper, examScore.PaperID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试卷信息: " + result.Error.Error()})
			return
		}
		getExamScoresByStudentIDResponses = append(getExamScoresByStudentIDResponses, model.GetExamScoresByStudentIDResponse{
			ID:               examScore.ID,
			PaperID:          examScore.PaperID,
			PaperName:        examPaper.Name,
			GradingTeacherID: examScore.GradingTeacherID,
			Comments:         examScore.Comments,
			Score:            examScore.Score,
			AddTime:          examScore.AddTime,
		})
	}
	c.JSON(http.StatusOK, getExamScoresByStudentIDResponses)
}

// GetAllStudentScores 获取所有学生的成绩
func GetAllStudentScores(c *gin.Context) {
	var examScores []model.ExamScore
	result := database.DB.Find(&examScores)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取成绩信息: " + result.Error.Error()})
		return
	}

	var getExamScoresByStudentIDResponses []model.GetExamScoresByStudentIDResponse
	for _, examScore := range examScores {
		var examPaper model.ExamPaper
		result := database.DB.First(&examPaper, examScore.PaperID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试卷信息: " + result.Error.Error()})
			return
		}

		var student model.Student
		result = database.DB.Where("id = ?", examScore.StudentID).First(&student)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取学生信息: " + result.Error.Error()})
			return
		}

		// 查询user表
		var user model.User
		result = database.DB.Where("id = ?", student.UserID).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息: " + result.Error.Error()})
			return
		}

		getExamScoresByStudentIDResponses = append(getExamScoresByStudentIDResponses, model.GetExamScoresByStudentIDResponse{
			ID:               examScore.ID,
			PaperID:          examScore.PaperID,
			PaperName:        examPaper.Name,
			StudentID:        student.ID,
			StudentName:      user.Username,
			GradingTeacherID: examScore.GradingTeacherID,
			Comments:         examScore.Comments,
			Score:            examScore.Score,
			AddTime:          examScore.AddTime,
		})
	}
	c.JSON(http.StatusOK, getExamScoresByStudentIDResponses)
}

// CreateExamPaperAQuestion 创建试卷和试题的接口
func CreateExamPaperAQuestion(c *gin.Context) {
	var request struct {
		ExamPaper     model.ExamPaper      `json:"exam_paper"`
		ExamQuestions []model.ExamQuestion `json:"exam_questions"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}

	// 设置添加时间
	request.ExamPaper.AddTime = time.Now().Format("2006-01-02 15:04:05")

	request.ExamPaper.TeacherId = 1

	// 创建试卷
	result := database.DB.Create(&request.ExamPaper)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建试卷: " + result.Error.Error()})
		return
	}

	// 获取创建的试卷 ID
	examPaperID := request.ExamPaper.ID

	// 创建试题
	for i := range request.ExamQuestions {
		request.ExamQuestions[i].AddTime = time.Now().Format("2006-01-02 15:04:05")
		request.ExamQuestions[i].PaperID = examPaperID
		request.ExamQuestions[i].Type = 1
		result := database.DB.Create(&request.ExamQuestions[i])
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建试题: " + result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"exam_paper":     request.ExamPaper,
		"exam_questions": request.ExamQuestions,
	})
}

// GetAllUsersAndInfo 获取所有用户信息
func GetAllUsersAndInfo(c *gin.Context) {
	var users []model.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息: " + result.Error.Error()})
		return
	}

	var usersAndInfo []model.UserAndInfo
	for _, user := range users {
		var userAndInfo model.UserAndInfo
		userAndInfo.User = user

		if user.Role == "student" {
			var student model.Student
			result := database.DB.Where("user_id = ?", user.ID).First(&student)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取学生信息: " + result.Error.Error()})
				return
			}
			userAndInfo.Student = student
		} else if user.Role == "teacher" {
			var teacher model.Teacher
			result := database.DB.Where("user_id = ?", user.ID).First(&teacher)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取教师信息: " + result.Error.Error()})
				return
			}
			userAndInfo.Teacher = teacher
		}

		usersAndInfo = append(usersAndInfo, userAndInfo)
	}

	c.JSON(http.StatusOK, usersAndInfo)
}

// CreateUserWithType 创建用户和相关信息
func CreateUserWithType(c *gin.Context) {
	var request struct {
		User    model.User    `json:"user"`
		Student model.Student `json:"student"`
		Teacher model.Teacher `json:"teacher"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}

	// 设置 add_time 字段为当前时间
	request.User.AddTime = time.Now()

	// 创建用户
	result := database.DB.Create(&request.User)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建用户: " + result.Error.Error()})
		return
	}

	// 获取创建的用户 ID
	userID := request.User.ID

	// 根据角色创建学生或老师信息
	if request.User.Role == "student" {
		request.Student.UserID = userID
		request.Student.AddTime = time.Now()
		result := database.DB.Create(&request.Student)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建学生: " + result.Error.Error()})
			return
		}
	} else if request.User.Role == "teacher" {
		request.Teacher.UserID = userID
		request.Teacher.AddTime = time.Now()
		result := database.DB.Create(&request.Teacher)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建老师: " + result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":    request.User,
		"student": request.Student,
		"teacher": request.Teacher,
	})
}

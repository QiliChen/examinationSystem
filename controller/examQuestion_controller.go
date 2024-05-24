package controller

import (
	"errors"
	"examination_system/database"
	"examination_system/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetExamQuestions(c *gin.Context) {
	var examQuestions []model.ExamQuestion
	result := database.DB.Find(&examQuestions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试题信息: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, examQuestions)
}

func GetExamQuestionByID(c *gin.Context) {
	id := c.Param("id")
	var examQuestion model.ExamQuestion
	result := database.DB.First(&examQuestion, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到试题"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试题信息: " + result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, examQuestion)
}

func GetExamQuestionByExamPaperID(c *gin.Context) {
	examPaperID := c.Param("examPaperID")
	var examQuestions []model.ExamQuestion
	result := database.DB.Where("paper_id = ?", examPaperID).Find(&examQuestions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试题信息: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, examQuestions)

}

func CreateExamQuestion(c *gin.Context) {
	var examQuestion model.ExamQuestion
	if err := c.ShouldBindJSON(&examQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	result := database.DB.Create(&examQuestion)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建试题: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, examQuestion)
}

func UpdateExamQuestion(c *gin.Context) {
	id := c.Param("id")
	var examQuestion model.ExamQuestion
	result := database.DB.First(&examQuestion, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到试题"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试题信息: " + result.Error.Error()})
		}
		return
	}
	if err := c.ShouldBindJSON(&examQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	database.DB.Save(&examQuestion)
	c.JSON(http.StatusOK, examQuestion)
}

func DeleteExamquestion(c *gin.Context) {
	id := c.Param("id")
	var examQuestion model.ExamQuestion
	result := database.DB.First(&examQuestion, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到试题"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试题信息: " + result.Error.Error()})
		}
		return
	}
	database.DB.Delete(&examQuestion)
	c.JSON(http.StatusOK, gin.H{"message": "试题已删除"})
}

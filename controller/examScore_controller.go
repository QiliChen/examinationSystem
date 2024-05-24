package controller

import (
	"errors"
	"examination_system/database"
	"examination_system/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetExamScores 获取所有成绩
func GetExamScores(c *gin.Context) {
	var examScores []model.ExamScore
	result := database.DB.Find(&examScores)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取成绩信息: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, examScores)
}

// GetExamScoreByID 根据ID获取成绩
func GetExamScoreByID(c *gin.Context) {
	id := c.Param("id")
	var examScore model.ExamScore
	result := database.DB.First(&examScore, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到成绩"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取成绩信息: " + result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, examScore)
}

// CreateExamScore 创建成绩
func CreateExamScore(c *gin.Context) {
	var examScore model.ExamScore
	if err := c.ShouldBindJSON(&examScore); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	examScore.AddTime = time.Now()
	result := database.DB.Create(&examScore)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建成绩: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, examScore)
}

// UpdateExamScore 更新成绩
func UpdateExamScore(c *gin.Context) {
	id := c.Param("id")
	var examScore model.ExamScore
	result := database.DB.First(&examScore, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到成绩"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取成绩信息: " + result.Error.Error()})
		}
		return
	}
	if err := c.ShouldBindJSON(&examScore); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	database.DB.Save(&examScore)
	c.JSON(http.StatusOK, examScore)
}

// DeleteExamScore 删除成绩
func DeleteExamScore(c *gin.Context) {
	id := c.Param("id")
	var examScore model.ExamScore
	result := database.DB.First(&examScore, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到成绩"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取成绩信息: " + result.Error.Error()})
		}
		return
	}
	database.DB.Delete(&examScore)
	c.JSON(http.StatusOK, gin.H{"message": "成绩已删除"})
}

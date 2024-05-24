package controller

import (
	"errors"
	"examination_system/database"
	"examination_system/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetExamPapers(c *gin.Context) {
	var examPapers []model.ExamPaper
	result := database.DB.Find(&examPapers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试卷信息: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, examPapers)
}

func GetExamPaperByID(c *gin.Context) {
	id := c.Param("id")
	var examPaper model.ExamPaper
	result := database.DB.First(&examPaper, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到试卷"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试卷信息: " + result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, examPaper)
}

func CreateExamPaper(c *gin.Context) {
	var examPaper model.ExamPaper
	if err := c.ShouldBindJSON(&examPaper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	result := database.DB.Create(&examPaper)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建试卷: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, examPaper)
}

func UpdateExamPaper(c *gin.Context) {
	id := c.Param("id")
	var examPaper model.ExamPaper
	result := database.DB.First(&examPaper, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到试卷"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试卷信息: " + result.Error.Error()})
		}
		return
	}
	if err := c.ShouldBindJSON(&examPaper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	database.DB.Save(&examPaper)
	c.JSON(http.StatusOK, examPaper)
}

func DeleteExamPaper(c *gin.Context) {
	id := c.Param("id")
	var examPaper model.ExamPaper
	result := database.DB.First(&examPaper, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到试卷"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取试卷信息: " + result.Error.Error()})
		}
		return
	}
	database.DB.Delete(&examPaper)
	c.JSON(http.StatusOK, gin.H{"message": "试卷已删除"})
}

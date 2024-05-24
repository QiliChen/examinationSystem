package controller

import (
	"errors"
	"examination_system/database"
	"examination_system/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTeachers(c *gin.Context) {
	var teachers []model.Teacher
	result := database.DB.Find(&teachers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取教师信息: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, teachers)
}

func GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	var teacher model.Teacher
	result := database.DB.First(&teacher, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到教师"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取教师信息: " + result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func CreateTeacher(c *gin.Context) {
	var teacher model.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	result := database.DB.Create(&teacher)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建教师: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, teacher)
}

func UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher model.Teacher
	result := database.DB.First(&teacher, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到教师"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取教师信息: " + result.Error.Error()})
		}
		return
	}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	database.DB.Save(&teacher)
	c.JSON(http.StatusOK, teacher)
}

func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher model.Teacher
	result := database.DB.First(&teacher, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到教师"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取教师信息: " + result.Error.Error()})
		}
		return
	}
	database.DB.Delete(&teacher)
	c.JSON(http.StatusOK, gin.H{"message": "教师已删除"})
}

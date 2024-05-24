package controller

import (
	"errors"
	"examination_system/database"
	"examination_system/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStudents(c *gin.Context) {
	var students []model.Student
	result := database.DB.Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取学生信息: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	var student model.Student
	result := database.DB.First(&student, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到学生"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取学生信息: " + result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	result := database.DB.Create(&student)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建学生: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student model.Student
	result := database.DB.First(&student, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到学生"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取学生信息: " + result.Error.Error()})
		}
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据: " + err.Error()})
		return
	}
	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student model.Student
	result := database.DB.First(&student, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到学生"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取学生信息: " + result.Error.Error()})
		}
		return
	}
	database.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "学生已删除"})
}

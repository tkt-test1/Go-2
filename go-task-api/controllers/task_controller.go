package controllers

import (
	"go-task-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskInput struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func CreateTask(c *gin.Context) {
	email, _ := c.Get("email")
	var user models.User
	models.DB.Where("email = ?", email).First(&user)

	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task := models.Task{
		Title:  input.Title,
		Status: input.Status,
		UserID: user.ID,
	}
	models.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func ListTasks(c *gin.Context) {
	email, _ := c.Get("email")
	var user models.User
	models.DB.Where("email = ?", email).First(&user)

	var tasks []models.Task
	title := c.Query("title")
	status := c.Query("status")

	query := models.DB.Where("user_id = ?", user.ID)
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Preload("Histories").Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := models.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// ステータス変更があれば履歴を追加
	if task.Status != input.Status {
		history := models.TaskHistory{
			TaskID:    task.ID,
			OldStatus: task.Status,
			NewStatus: input.Status,
		}
		models.DB.Create(&history)
	}

	task.Title = input.Title
	task.Status = input.Status
	models.DB.Save(&task)

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := models.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	models.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

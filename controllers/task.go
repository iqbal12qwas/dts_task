package controllers

import (
    "gin-api/models"
    "gin-api/repository"
    "gin-api/utils"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
    IdProfile int `json:"id_profile"`
    AssingedTo string `json:"assignedTo"`
    Task       string `json:"task"`
    Deadline   string `json:"deadline`
}

type UpdateTaskInput struct {
    AssingedTo string `json:"assignedTo"`
    Task       string `json:"task"`
    Deadline   string `json:"deadline`
}

// Pagination
func GetPaginateTasks(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var task models.Task
	taskLists, err := repository.GetAllTasks(&task, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
        "meta": pagination,
		"data": taskLists,
	})

}

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {
    tasks, err := repository.AllTasks()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 400, "message": "Failed", "error": "Record not found!"})
        return
    } 
    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": tasks})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
    // Validate input
    var input CreateTaskInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    date := "2006-01-02"
    deadline, _ := time.Parse(date, input.Deadline)

    // Create task
    task := models.Task{IdProfile: input.IdProfile, AssingedTo: input.AssingedTo, Task: input.Task, Deadline: deadline}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&task)

    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": task})
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
    var task models.Task
    task, err := repository.FindTaskById(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 400, "message": "Failed", "error": "Record not found!"})
        return
    } 
    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": task})
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {

    // Get model if exist
    var task models.Task
    task, err := repository.FindTaskById(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 400, "message": "Failed", "error": "Record not found!"})
        return
    } 

    // Validate input
    var input UpdateTaskInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 200, "message": "Failed", "error": err.Error()})
        return
    }

    date := "2006-01-02"
    deadline, _ := time.Parse(date, input.Deadline)

    var updatedInput models.Task
    updatedInput.Deadline = deadline
    updatedInput.AssingedTo = input.AssingedTo
    updatedInput.Task = input.Task

    _, err = repository.UpdateTask(task, updatedInput)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 200, "message": "Failed", "error": err.Error()})
        return
    } 

    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": &task})
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
    // Get model if exist
    var task models.Task
    task, err := repository.FindTaskById(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 400, "message": "Failed", "error": "Record not found!"})
        return
    } 

    err = repository.DeleteTask(task)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 200, "message": "Failed", "error": err.Error()})
        return
    } 

    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": task})
}
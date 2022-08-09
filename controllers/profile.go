package controllers

import (
    "gin-api/models"
	"gin-api/repository"
	"gin-api/utils"
    // "gin-api/middleware"
    "net/http"
    "time"
    // "fmt"

    "github.com/gin-gonic/gin"
    // "github.com/jinzhu/gorm"
)

type CreateProfileInput struct {
    Name string `json:"name" validate:"required"`
    Birth   string `json:"birth" validate:"required"`
}

type UpdateProfileInput struct {
    Name string `json:"name" validate:"required"`
    Birth   string `json:"birth" validate:"required"`
}

// Pagination
func GetPaginateProfiles(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var profile models.Profile
	profileLists, err := repository.GetAllProfiles(&profile, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
        "meta": pagination,
		"data": profileLists,
	})

}

// GET /profiles
// Get all profiles
func FindProfiles(c *gin.Context) {
    profiles, err := repository.AllProfiles()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 400, "message": "Success", "error": "Record not found!"})
        return
    } 
    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": profiles})
}

// POST /profiles
// Create new profiles
func CreateProfile(c *gin.Context) {
	var input CreateProfileInput

    // Validate input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 200, "message": "Failed", "error": err.Error()})
        return
    }

    date := "2006-01-02"
    birth, _ := time.Parse(date, input.Birth)

    // Create profiles
    profile := models.Profile{Name: input.Name, Birth: birth}
    // db := c.MustGet("db").(*gorm.DB)
    // db.Create(&profile)

    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": profile})
}

// GET /profiles/:id
// Find a profiles
func FindProfile(c *gin.Context) { // Get model if exist
    var profile models.Profile

	profile, err := repository.FindProfileById(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 400, "message": "Failed", "error": "Record not found!"})
        return
    } 

    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": profile})
}

// PATCH /profiles/:id
// Update a profile
func UpdateProfile(c *gin.Context) {

    // Get model if exist
    var profile models.Profile
    profile, err := repository.FindProfileById(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 400, "message": "Failed", "error": "Record not found!"})
        return
    } 

    // Validate input
    var input UpdateProfileInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 200, "message": "Failed", "error": err.Error()})
        return
    }

    date := "2006-01-02"
    birth, _ := time.Parse(date, input.Birth)

    var updatedInput models.Profile
    updatedInput.Name = input.Name
    updatedInput.Birth = birth

    _, err = repository.UpdateProfile(profile, updatedInput)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 200, "message": "Failed", "error": err.Error()})
        return
    } 

    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": &profile})
}

// DELETE /profiles/:id
// Delete a profile
func DeleteProfile(c *gin.Context) {
    // Get model if exist
    var profile models.Profile
    profile, err := repository.FindProfileById(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 400, "message": "Failed", "error": "Record not found!"})
        return
    } 

    // Delete Task
    err = repository.DeleteTaskByIdProfile(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 200, "message": "Failed", "error": err.Error()})
        return
    } 

    // Delete Profile
    err = repository.DeleteProfile(profile)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code_message" : 200, "message": "Failed", "error": err.Error()})
        return
    } 

    c.JSON(http.StatusOK, gin.H{"code_message" : 200, "message": "Success", "data": profile})
}


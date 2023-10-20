package controllers

import (
	"go-crud/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateGatheringInput struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	TypeID   uint   `json:"type_id"`
	// ScheduledAt time.Time `json:"scheduled_at"`
}

type UpdateGatheringInput struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	TypeID   uint   `json:"type_id"`
	// ScheduledAt time.Time `json:"scheduled_at"`
}

// GET /members
// Get all members
func FindGatherings(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data []models.Gathering
	db.Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// POST /members
// Create new member
func CreateGathering(c *gin.Context) {
	// Validate input
	var input CreateGatheringInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	data := models.Gathering{Name: input.Name, Location: input.Location, TypeID: input.TypeID}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GET /members/:id
// Find a member
func FindGathering(c *gin.Context) { // Get model if exist
	var data models.Gathering

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// PATCH /members/:id
// Update a member
func UpdateGathering(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var data models.Gathering
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateGatheringInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Gathering
	updatedInput.Name = input.Name
	updatedInput.Location = input.Location
	updatedInput.TypeID = input.TypeID
	// updatedInput.ScheduledAt = input.ScheduledAt

	db.Model(&data).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// DELETE /members/:id
// Delete a member
func DeleteGathering(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var data models.Gathering
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&data)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

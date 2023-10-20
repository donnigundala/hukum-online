package controllers

import (
	"go-crud/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICreateGatheringType struct {
	Name string `json:"name"`
}

type IUpdateGatheringType struct {
	Name string `json:"name"`
}

// GET /gathering-types
// Get all gathering-types
func FindGatheringTypes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data []models.GatheringType
	db.Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// POST /gathering-type
// Create new gathering-type
func CreateGatheringType(c *gin.Context) {
	// Validate input
	var input ICreateGatheringType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create gathering-type
	data := models.GatheringType{Name: input.Name}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GET /gathering-type/:id
// Find a gathering-type
func FindGatheringType(c *gin.Context) { // Get model if exist
	var data models.GatheringType

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// PATCH /gathering-type/:id
// Update a gathering-type
func UpdateGatheringType(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var data models.GatheringType
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input IUpdateGatheringType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.GatheringType
	updatedInput.Name = input.Name

	db.Model(&data).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// DELETE /members/:id
// Delete a member
func DeleteGatheringType(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var data models.GatheringType
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&data)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

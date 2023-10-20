package controllers

import (
	"go-crud/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateMemberInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UpdateMemberInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// GET /members
// Get all members
func FindMembers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var members []models.Member
	db.Find(&members)

	c.JSON(http.StatusOK, gin.H{"data": members})
}

// POST /members
// Create new member
func CreateMember(c *gin.Context) {
	// Validate input
	var input CreateMemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	member := models.Member{FirstName: input.FirstName, LastName: input.LastName, Email: input.Email}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&member)

	c.JSON(http.StatusOK, gin.H{"data": member})
}

// GET /members/:id
// Find a member
func FindMember(c *gin.Context) { // Get model if exist
	var member models.Member

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}

// PATCH /members/:id
// Update a member
func UpdateMember(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var member models.Member
	if err := db.Where("id = ?", c.Param("id")).First(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateMemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Member
	updatedInput.FirstName = input.FirstName
	updatedInput.LastName = input.LastName
	updatedInput.Email = input.Email

	db.Model(&member).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": member})
}

// DELETE /members/:id
// Delete a member
func DeleteMember(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var member models.Member
	if err := db.Where("id = ?", c.Param("id")).First(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&member)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

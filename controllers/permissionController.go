package controllers

import (
	"basicAuth/initializers"
	"basicAuth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionInput struct {
	Name string `json:"name"`
}

func CreatePermission(ctx *gin.Context) {
	var permissionInput PermissionInput
	if err := ctx.ShouldBindJSON(&permissionInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var permissionMoel models.Permission
	initializers.DB.Where("name=?", permissionInput.Name).First(&permissionMoel)
	if permissionMoel.ID != 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Permission already exists"})
		return
	}

	permission := models.Permission{Name: permissionInput.Name}
	initializers.DB.Create(&permission)
	ctx.JSON(http.StatusOK, gin.H{"message": "Permission Created Successfully"})
}

func AssignPermeissions(ctx *gin.Context) {
	var assignPermissionInput struct {
		UserID        int    `json:"user_id"`
		PermissionIDs []uint `json:"permission_ids"`
	}
	if err := ctx.ShouldBindJSON(&assignPermissionInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var user models.User
	initializers.DB.Where("id=?", assignPermissionInput.UserID).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch the permissions from the database
	var permissions []models.Permission
	if err := initializers.DB.Where("id IN ?", assignPermissionInput.PermissionIDs).Find(&permissions).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Permissions not found"})
		return
	}

	// Prepare a slice of permission objects
	var userPermissions []models.Permission

	for _, permission := range permissions {
		userPermissions = append(userPermissions, permission)
	}
	// Sync the user's permissions
	if err := initializers.DB.Model(&user).Association("Permission").Replace(userPermissions); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign permissions"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Permission Assigned Successfully"})
}

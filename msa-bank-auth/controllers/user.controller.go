package controllers

import (
	"msa-bank-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Email:     currentUser.Email,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
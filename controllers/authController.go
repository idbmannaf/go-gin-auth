package controllers

import (
	"basicAuth/initializers"
	"basicAuth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *gin.Context) {
	var authInput models.AuthInput
	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("user_name =?", authInput.Username).Find(&userFound)
	if userFound.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
	}
	paswordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		UserName: authInput.Username,
		Password: string(paswordHash),
	}
	if authInput.IsAdmin {
		user.IsAdmin = true
	}

	initializers.DB.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(ctx *gin.Context) {
	var authInput models.AuthInput

	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	// initializers.DB.Where("user_name =?", authInput.Username).First(&userFound)
	if err := initializers.DB.Preload("Permission").Where("user_name = ?", authInput.Username).First(&userFound).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	if userFound.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      userFound.ID,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "ManageError": "Something Wrong to generate token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString, "user": userFound})
}

func GetUserProfile(ctx *gin.Context) {
	user, _ := ctx.Get("currentUser")
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

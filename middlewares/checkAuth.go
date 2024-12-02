package middlewares

import (
	"basicAuth/initializers"
	"basicAuth/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth(ctx *gin.Context) {
	authHeder := ctx.GetHeader("Authorization")
	if authHeder == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeder, " ")

	if len(authToken) != 2 || authToken[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Format"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}

	// Check for the expiration claim
	exp, ok := claims["expires"].(float64)
	if !ok || exp == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing expiration claim"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > exp {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	// initializers.DB.Where("ID=?", claims["id"]).Find(&user)
	if err := initializers.DB.Preload("Permission").First(&user, claims["id"]).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	if user.ID == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("currentUser", user)

	ctx.Next()
}

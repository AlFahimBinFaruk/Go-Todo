package user

import (
	"go-todo/initializers"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	type REQ struct {
		Username string
		Email    string
		Password string
	}
	var body REQ
	if err := c.Bind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request."})
		return
	}
	var user models.USER
	exixts := initializers.DB.First(&user, "email = ?", body.Email)
	if exixts.Error == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User Exixts."})
		return
	}
	// hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate password."})
		return
	}
	newUser := models.USER{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hashPassword),
	}
	if err := initializers.DB.Create(&newUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "User Registered Successfully."})
}

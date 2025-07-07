package user

import (
	"go-todo/initializers"
	"go-todo/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	type REQ struct {
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
	if exixts.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User Don't Exixts."})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong password."})
		return
	}
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		Subject:   user.ID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedtoken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": signedtoken,
	})

}

package todo

import (
	"go-todo/initializers"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTodos(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Interner Server Error in => GetTodos.",
			})
		}
	}()

	userIdStr, _ := c.Get("userId")
	userId, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request."})
		return
	}

	var user models.USER
	exixts := initializers.DB.Preload("Todos").First(&user, "id = ?", userId)
	if exixts.Error != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"error": "No user found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"todos": user.Todos,
	})

}

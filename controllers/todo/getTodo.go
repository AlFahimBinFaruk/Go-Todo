package todo

import (
	"go-todo/initializers"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTodo(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Interner Server Error in => GetTodo.",
			})
		}
	}()

	id := c.Param("id")
	userIdStr, _ := c.Get("userId")
	userId, err := uuid.Parse(userIdStr.(string))
	if err != nil || id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request."})
		return
	}

	// see if user exixts
	var user models.USER
	exixts := initializers.DB.First(&user, "id = ?", userId)
	if exixts.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "User not found.",
		})
	}
	var todo models.TODO
	result := initializers.DB.First(&todo, "id = ?", id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "failed to get todo!!",
		})
		return
	}
	// if user not authorized
	if todo.UserId != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized to perform this.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}

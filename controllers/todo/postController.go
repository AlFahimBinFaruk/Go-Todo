package todo

import (
	"go-todo/initializers"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTodo(c *gin.Context) {
	type REQ struct {
		Title string
		Desc  string
	}
	var body REQ
	if err := c.Bind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request."})
		return
	}
	userIdStr, _ := c.Get("userId")
	userId, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request."})
		return
	}

	userExixts := initializers.DB.First(&models.USER{}, "id = ?", userId)
	if userExixts.Error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User not found."})
		return
	}

	todo := models.TODO{Title: body.Title, Desc: body.Desc, UserId: userId}
	result := initializers.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "Failed to add Todo.",
		})
		return
	}
	c.JSON(202, gin.H{
		"todo": todo,
	})
}

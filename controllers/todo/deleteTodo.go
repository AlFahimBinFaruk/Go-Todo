package todo

import (
	"go-todo/initializers"
	"go-todo/models"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	result := initializers.DB.Delete(&models.TODO{}, id)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "failed to get todo!!",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "Deleted successfully!!",
	})
}

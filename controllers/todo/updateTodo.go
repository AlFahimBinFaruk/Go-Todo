package todo

import (
	"go-todo/initializers"
	"go-todo/models"

	"github.com/gin-gonic/gin"
)

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.TODO
	result := initializers.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "failed to get todo!!",
		})
		return
	}
	type REQ struct {
		Title string
		Desc  string
	}
	var body REQ
	c.Bind(&body)
	result = initializers.DB.Model(&todo).Updates(models.TODO{
		Title: body.Title,
		Desc:  body.Desc,
	})
	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "failed to update todo!!",
		})
		return
	}
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

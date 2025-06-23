package migrate

import (
	"go-todo/initializers"
	"go-todo/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.TODO{})
}

package main

import (
	"sekolah/config"
	userController "sekolah/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Connect()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/data", userController.Index)
	r.POST("/daftar", userController.Create)
	r.PUT("/update/:id", userController.Update)
	r.DELETE("/hapus/:id", userController.Delete)
	r.Run(":3000")
}

// func loadEnv() {
// 	err := godotenv.Load(".env.local")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

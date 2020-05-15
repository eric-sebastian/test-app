package main

import (
	"test-app/src/internal/handler/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/users", http.GetUsers())
	r.GET("/api/user/:id", http.GetUser())
	r.POST("/api/users", http.CreateUser())
	r.PUT("/api/user/:id", http.UpdateUser())
	r.DELETE("/api/user/:id", http.DeleteUser())
	r.Run(":8000")
}

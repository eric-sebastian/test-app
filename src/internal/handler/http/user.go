package http

import (
	"test-app/src/internal/entity/user_model"
	"test-app/src/internal/repo/user"

	"github.com/gin-gonic/gin"
)

//GetUsers :
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "",
			"data":    user.GetUsers(),
		})
	}
}

//GetUser :
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status":  200,
			"message": "",
			"data":    user.GetUser(id),
		})
	}
}

//CreateUser :
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser user_model.UserForm
		if c.BindJSON(&newUser) == nil {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "User created!",
				"data":    user.CreateUser(newUser),
			})
		} else {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Bad Request",
				"data":    "",
			})
		}
	}
}

//UpdateUser :
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updateUser user_model.UserForm
		if c.BindJSON(&updateUser) == nil {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "User updated!",
				"data":    user.UpdateUser(id, updateUser),
			})
		} else {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Bad Request",
				"data":    "",
			})
		}
	}
}

//DeleteUser :
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user.DeleteUser(id)
		c.JSON(200, gin.H{
			"status":  200,
			"message": "User deleted!",
			"data":    "",
		})
	}
}

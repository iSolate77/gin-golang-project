package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
  Username string `json:"username"`
  Email string `json:"email"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
}

var users = []User{
  {Username: "user1", Email: "test@test.c", FirstName: "John", LastName: "Doe"},
  {Username: "user2", Email: "test2@test.c", FirstName: "Jane", LastName: "Doe"},
}

func main() {
	r := gin.Default()
  r.GET("/users", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"users": users})
  })
    r.Run(":8080")
}

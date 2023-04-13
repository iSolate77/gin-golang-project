package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
  "gin/internals"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		data := gin.H{"title": "Home page"}
		c.HTML(http.StatusOK, "homepage.tmpl", data)
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"users": internals.Users})
	})

	r.Run()
}

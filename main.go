package main

import (
	"github.com/gin-gonic/gin"
)

func getStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Vini",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", getStudents)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

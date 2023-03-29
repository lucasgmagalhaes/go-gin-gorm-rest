package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/alunos", alunosHandler)

	r.Run(":5000")
}

func alunosHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Lucas",
	})
}

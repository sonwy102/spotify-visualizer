package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// func main() {
//   // Set the router as the default one shipped with Gin
//   router := gin.Default()

//   // Serve frontend static files
//   // router.Use(static.Serve("/", static.LocalFile("./views", true)))
  
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
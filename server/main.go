package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func main() {
	var err error

	router := gin.Default()

	// Serve frontend static files and exclude /api routes
	router.Use(func(context *gin.Context) {
		if strings.HasPrefix(context.Request.URL.Path, "/api") {
			context.Next()
		} else {
			http.ServeFile(context.Writer, context.Request, "./website/build"+context.Request.URL.Path)
			context.Abort()
		}
	})

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err = router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}

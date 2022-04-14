package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {

	port := os.Getenv("PORT")

	r = gin.Default()

	r.LoadHTMLGlob("static/*")

	intializeRoutes()

	r.Run(":" + port)
}
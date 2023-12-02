package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
	"net/http"
)

func main() {
	ginS.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello Jenkins")
	})
	_ = ginS.Run(":8081")
}

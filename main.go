package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
	"net/http"
)

func main() {
	ginS.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello docker")
	})
	_ = ginS.Run(":8081")
}

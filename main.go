package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	okMessage struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}
)

func ResponseOKMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, okMessage{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

func main() {
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ResponseOKMessage(ctx)
	})

	_ = router.Run(":8080")
}

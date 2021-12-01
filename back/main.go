package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.New()

	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "Hello",
		})
	})

	err := engine.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}

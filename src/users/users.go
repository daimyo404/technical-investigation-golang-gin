package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type fugaType struct {
	id      string
	message string
}

var fugavar = fugaType{
	id:      "fuga",
	message: "fugafuga",
}

func fuga(c *gin.Context) {
	c.JSON(http.StatusOK, fugavar.id)
}

func Users() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("hoge")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/fuga", fuga)

	router.Run()
}

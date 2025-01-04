package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()


	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello test completed successfully"})
	})

	return r
}

func main() {
	if inLambda() {
		fmt.Println("running aws lambda in aws")
		log.Fatal(gateway.ListenAndServe(":8080", setupRouter()))
	} else {
		fmt.Println("running aws lambda in local")
		log.Fatal(http.ListenAndServe(":8080", setupRouter()))
	}
}

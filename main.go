package main

import (
	  "fmt"
    "github.com/samarsingh/golang-gin-jwt-auth-crud/api/router"
    "github.com/samarsingh/golang-gin-jwt-auth-crud/config"
    "github.com/samarsingh/golang-gin-jwt-auth-crud/db/initializers"
    "github.com/gin-gonic/gin"
)
// go mod init github.com/samarsingh@bitcot.com/golang-gin-jwt-auth-crud

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello auth")
	r := gin.Default()
	router.GetRoute(r)

	r.Run()
}

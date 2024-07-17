package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SignInInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type SignUpInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func main() {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"https://login-web-test.vercel.app/"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    r.POST("/signin", func(c *gin.Context) {
        var input SignInInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        fmt.Printf("Email: %s, Password: %s\n", input.Email, input.Password)

        c.JSON(http.StatusOK, gin.H{"message": "Sign in successful"})
    })

	r.POST("/signup", func(c *gin.Context) {
        var input SignInInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        fmt.Printf("Sign up Email: %s, Password: %s\n", input.Email, input.Password)

        c.JSON(http.StatusOK, gin.H{"message": "Sign in successful"})
    })

    r.Run(":8080")
}

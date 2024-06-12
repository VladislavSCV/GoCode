package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions"
    // "github.com/gin-contrib/sessions/cookie"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("Before the handler")
        c.Next()
        fmt.Println("After the handler")
    }
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob(`C:\Users\VladislavSCV\OneDrive\Desktop\VSC\GO\GoCode\web\templates\*`)
	r.Static("/static", "../static")

	// store := cookie.NewStore([]byte("secret"))
    // r.Use(sessions.Sessions("mysession", store))
	r.Use(LoggerMiddleware())

	r.GET("/", GetIndex)
	r.GET("/SignUp", GetSignUP)
	r.GET("/Login", GetLogin)
	r.GET("/GoCode", GetMainPage)
	r.GET("/Task", GetTask)
	r.POST("/Task", PostTask)
	r.GET("/Profile", GetProfile)
	r.GET("/Resourses", GetResourses)
	r.GET("/Catalog", GetCatalog)
	r.POST("/Catalog", PostCatalog)

	r.Run(":8000")
}


package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions"
    // "github.com/gin-contrib/sessions/cookie"
)





type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob(`..\templates/*`)
	r.Static("/static", "../static")

	r.Use(LoggerMiddleware())

	r.GET("/", GetIndex)
	r.GET("/SignUp", GetSignUP)
	r.POST("/PostSign", PostSign)
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


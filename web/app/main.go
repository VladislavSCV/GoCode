package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Создаем хранилище сессий на основе файлов cookie
	store := cookie.NewStore([]byte("secret"))
	r.LoadHTMLGlob(`..\templates/*`)
	r.Static("/static", "../static")

	r.Use(LoggerMiddleware())
	// Используем middleware для сессий
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", GetIndex)
	r.GET("/signup", GetSignUP)
	r.GET("/signupnextstep", GetSignUpNextStep)
	r.POST("/signup/step1", SignUpSaveENP)
	r.POST("/signup/step2", PostSign)
	r.GET("/login", GetLogin)
	r.POST("/PostLogin", PostLogin)
	r.GET("/main", GetMainPage)
	r.GET("/task", GetTask)
	r.POST("/task", PostTask)
	r.GET("/profile", GetProfile)
	r.GET("/resourses", GetResourses)
	r.GET("/catalog", GetCatalog)
	r.POST("/catalog", PostCatalog)
	r.POST("/ReSaveUserData", ReSaveUserData)
	r.GET("/QuitFromProfile", QuitFromProfile)

	r.Run(":8000")
}

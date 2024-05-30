package main

import (
	"fmt"
	"net/http"

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
	r.LoadHTMLGlob("../templates/*")
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


func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "TITLE",
	})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show SignUp Page
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func GetSignUP(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{
		 "status": http.StatusOK, 
	})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show Login Page
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func GetLogin(c *gin.Context) {
	
	c.HTML(http.StatusOK, "login.html", gin.H{
		 "status": http.StatusOK, 
	})
}



// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show main page
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func GetMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"status": http.StatusOK,
	})

}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show task page
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func GetTask(c *gin.Context) {
	// id := c.Param("id")
	// fmt.Println("id: " + id)
	c.HTML(http.StatusOK, "task.html", gin.H{
		"status": http.StatusOK,
	})

}


// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Post answer for task
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Param code formData string true "Code for check"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [post]

type CheckCode struct {
	Code string `json:"code" form:"code"`
}

func PostTask(c *gin.Context) {
	var checkCode CheckCode
	if err := c.ShouldBind(&checkCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: Проверка кода
	if CheckSolution(checkCode.Code) {
		c.JSON(http.StatusOK, gin.H{"message": "Решение принято"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Решение не принято"})
	}	
}


// Функция, которая проверяет решение на правильность
func CheckSolution(code string) bool {
	fmt.Println(code)
	// Здесь вы можете написать код, который будет проверять решение пользователя
	// Например, вы можете запустить тесты или выполнить код и проверить результат
	// В зависимости от результата возвращайте true или false
	// Например, если решение правильное, то возвращаем true, иначе false
	return true
}


// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show profile
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func GetProfile(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"status": http.StatusOK,
	})
}


// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show resurs page
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func GetResourses(c *gin.Context) {
	c.HTML(http.StatusOK, "resourses.html", gin.H{
		"status": http.StatusOK,
	})
}


// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show catalog
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func GetCatalog(c *gin.Context) {
	c.HTML(http.StatusOK, "catalog.html", gin.H{
		"status": http.StatusOK,
	})
}

func PostCatalog(c *gin.Context) {
    difficulty := c.PostForm("difficulty")
    category := c.PostForm("category")

    // TODO: get data from db
    // ...

    // TODO: render data instead of redirection
    c.HTML(http.StatusOK, "catalog.html", gin.H{
        "status": http.StatusOK,
        "difficulty": difficulty,
        "category": category,
    })
    c.Abort()
}

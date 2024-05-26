package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("Before the handler")
        c.Next()
        fmt.Println("After the handler")
    }
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../template/*")
	r.Static("/static", "../static")

	r.Use(LoggerMiddleware())

	r.GET("/", Index)
	r.GET("/SignUp", SignUP)
	r.GET("/Login", Login)
	r.GET("/GoCode", MainPage)
	r.GET("/Task", GetTask)
	r.POST("/Task", PostTask)

	r.Run(":8000")
}


func Index(c *gin.Context) {
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
func SignUP(c *gin.Context) {
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
func Login(c *gin.Context) {
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
func MainPage(c *gin.Context) {
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
	Verify_code string `json:"verify_code"`
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

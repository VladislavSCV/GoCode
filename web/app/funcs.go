package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	// "github.com/VladislavSCV/GoCode/pkg"
	pb "github.com/VladislavSCV/GoCode/api/grpc/gen/pb-go/com.user_data"
	"github.com/VladislavSCV/GoCode/pkg"
)

var (
	conn, errConGRPC = grpc.Dial(":50051", grpc.WithInsecure())
	grpcFunc = pb.NewUserDataMessageServiceClient(conn)
)

func init() {
	if errConGRPC != nil {
		log.Fatalf("did not connect: %v", errConGRPC)
	}
	defer conn.Close()
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
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func PostSign(c *gin.Context) {
	var HashedPassw string
	// получить данные из формы
	var formData struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
		Phone    string `form:"phone"`
	}
	if err := c.Bind(&formData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	passIsValid, err := pkg.CheckPassword(formData.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if passIsValid {
		HashedPassw, err = pkg.HashFuncPassword(formData.Password)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// use HashedPassw here or in other parts of the code
	fmt.Println("Hashed password:", HashedPassw)

	// работать с данными
	fmt.Println("post sign up")
	fmt.Println("username:", formData.Username)
	fmt.Println("email:", formData.Email)
	fmt.Println("password:", formData.Password)
	fmt.Println("phone:", formData.Phone)

	// Username     string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Email        string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Phone        string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	// AvatarUrl    string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// Status       string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// Role         string `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	// PasswordHash string `protobuf:"bytes,7,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	// DateOfBirth  *Date  `protobuf:"bytes,8,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`

	// r, err := grpcFunc.SignUp(context.Background(), &pb.SignUserDataRequest{
		
	// })
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }

	// Вывести страницу после успешной обработки данных
	c.HTML(http.StatusOK, "index.html", gin.H{
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

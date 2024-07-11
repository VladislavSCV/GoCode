package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	pb "github.com/VladislavSCV/GoCode/api/grpc/gen/pb-go/com.user_data"
	"github.com/VladislavSCV/GoCode/pkg"
)

var (
	conn, errConGRPC = grpc.Dial(":50051", grpc.WithInsecure())
	grpcFunc         = pb.NewUserDataMessageServiceClient(conn)
)

type FormDataSignUpFirstStep struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Phone    string `form:"phone"`
}

type FormDataSignUpSecondStep struct {
	Username    string `form:"username"`
	AvatarUrl   string `form:"avatar_url"`
	Role        string `form:"role"`
	DateOfBirth string `form:"date_of_birth"`
}

type FormDataLogin struct {
	UserEmail    string `form:"email"`
	UserPassword string `form:"password"`
}

func init() {
	if errConGRPC != nil {
		log.Fatalf("did not connect: %v", errConGRPC)
	}
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "TITLE",
	})
}

func GetSignUP(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{
		"status": http.StatusOK,
	})
}

func GetSignUpNextStep(c *gin.Context) {
	c.HTML(
		http.StatusOK, "signupNextStep.html", gin.H{
			"status": http.StatusOK,
		})
}

func SignUpSaveENP(c *gin.Context) {
	var dataENP FormDataSignUpFirstStep
	if err := c.Bind(&dataENP); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	UserEmail := dataENP.Email
	UserPassword := dataENP.Password
	UserPhone := dataENP.Phone

	passIsValid, err := pkg.CheckPassword(UserPassword)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var HashedPassw string
	if passIsValid {
		HashedPassw, err = pkg.HashFuncPassword(UserPassword)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	session := sessions.Default(c)
	session.Set("email", UserEmail)
	session.Set("password", HashedPassw)
	session.Set("phone", UserPhone)
	session.Save()

	// Redirect with appropriate status code
	c.Redirect(http.StatusFound, "/signupnextstep")
}

func PostSign(c *gin.Context) {
	var dataNURD FormDataSignUpSecondStep
	if err := c.Bind(&dataNURD); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	session := sessions.Default(c)
	emailInterface := session.Get("email")
	passwordInterface := session.Get("password")
	phoneInterface := session.Get("phone")

	if emailInterface == nil || passwordInterface == nil || phoneInterface == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	UserEmail, emailOK := emailInterface.(string)
	UserPassword, passwordOK := passwordInterface.(string)
	UserPhone, phoneOK := phoneInterface.(string)

	if !emailOK || !passwordOK || !phoneOK {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Printf("Sending gRPC request: %+v", &pb.SignUserDataRequest{
		Username:     dataNURD.Username,
		Email:        UserEmail,
		Phone:        UserPhone,
		AvatarUrl:    dataNURD.AvatarUrl,
		Role:         dataNURD.Role,
		PasswordHash: UserPassword,
		DateOfBirth:  dataNURD.DateOfBirth,
	})

	r, err := grpcFunc.SignUp(context.Background(), &pb.SignUserDataRequest{
		Username:     dataNURD.Username,
		Email:        UserEmail,
		Phone:        UserPhone,
		AvatarUrl:    dataNURD.AvatarUrl,
		Role:         dataNURD.Role,
		PasswordHash: UserPassword,
		DateOfBirth:  dataNURD.DateOfBirth,
	})
	if err != nil {
		log.Fatalf("could not sign up: %v", err)
	}

	log.Println(r.String())

	session.Set("name", dataNURD.Username)
	session.Save()

	c.HTML(http.StatusOK, "main.html", gin.H{
		"status": http.StatusOK,
	})
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"status": http.StatusOK,
	})
}

func PostLogin(c *gin.Context) {
	var dataLogin FormDataLogin
	if err := c.Bind(&dataLogin); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	r, err := grpcFunc.Login(context.Background(), &pb.LoginUserDataRequest{
		Email:        dataLogin.UserEmail,
		PasswordHash: dataLogin.UserPassword,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Println(r.String())
	session := sessions.Default(c)
	session.Set("description", r.Description)
	session.Set("name", r.UserName)
	session.Set("email", r.Email)
	session.Set("phone", r.Phone)
	session.Set("avatar_url", r.AvatarUrl)
	session.Set("status", r.Status)
	session.Set("role", r.Role)
	session.Set("date_of_birth", r.DateOfBirth)
	session.Set("privacy_settings", r.PrivacySettings)
	session.Set("is_active", r.IsActive)
	session.Set("last_login", r.LastLogin)
	session.Set("confirmation_token", r.ConfirmationToken)
	session.Set("social_profiles", r.SocialProfiles)
	session.Save()

	if r != nil {
		c.Redirect(http.StatusFound, "http://127.0.0.1:8000/main")
	} else {
		c.Redirect(http.StatusFound, "http://127.0.0.1:8000/signup")
	}
}

func GetMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"status": http.StatusOK,
	})
}

func GetTask(c *gin.Context) {
	c.HTML(http.StatusOK, "task.html", gin.H{
		"status": http.StatusOK,
	})
}

type CheckCode struct {
	Code string `json:"code" form:"code"`
}

func PostTask(c *gin.Context) {
	var checkCode CheckCode
	if err := c.ShouldBind(&checkCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if CheckSolution(checkCode.Code) {
		c.JSON(http.StatusOK, gin.H{"message": "Решение принято"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Решение не принято"})
	}
}

func CheckSolution(code string) bool {
	fmt.Println(code)
	return true
}

func GetProfile(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"username":        session.Get("name"),
		"email":           session.Get("email"),
		"phone":           session.Get("phone"),
		"avatar_url":      session.Get("avatar_url"),
		"status":          session.Get("status"),
		"role":            session.Get("role"),
		"date_of_birth":   session.Get("date_of_birth"),
		"is_active":       session.Get("is_active"),
		"last_login":      session.Get("last_login"),
		"social_profiles": session.Get("social_profiles"),
	})
}

func GetResourses(c *gin.Context) {
	c.HTML(http.StatusOK, "resourses.html", gin.H{
		"status": http.StatusOK,
	})
}

type Task struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Difficulty  string `json:"difficulty" form:"difficulty"`
	Category    string `json:"category" form:"category"`
	Solution    string `json:"solution" form:"solution"`
}

var Tasks []Task

// router.GET("/array", func(c *gin.Context) {
// 	var values []int
// 	for i := 0; i < 5; i++ {
// 		values = append(values, i)
// 	}

// 	c.HTML(http.StatusOK, "array.tmpl", gin.H{"values": values})
// })

// <ul>
//   {{ range .values }}
//   <li>{{ . }}</li>
//   {{ end }}
// </ul>

func GetCatalog(c *gin.Context) {
	c.HTML(http.StatusOK, "catalog.html", gin.H{
		"tasks": Tasks,
	})
}

func PostCatalog(c *gin.Context) {
	difficulty := c.PostForm("difficulty")
	category := c.PostForm("category")

	c.HTML(http.StatusOK, "catalog.html", gin.H{
		"status":     http.StatusOK,
		"difficulty": difficulty,
		"category":   category,
	})
}

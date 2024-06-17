package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// User is a struct
type User struct {
    ID   int
    Name string
}

var users = []User{
    {ID: 1, Name: "Duotify"},
    {ID: 2, Name: "Duotify"},
}

func main() {

    router := gin.Default()

    router.GET("/user", func(c *gin.Context) {
        c.JSON(http.StatusOK, users)
    })

    router.GET("/user/:id", func(c *gin.Context) {

        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
            return
        }

        for i := 0; i < len(users); i++ {
            if users[i].ID == id {
                c.JSON(http.StatusOK, users[i])
                break
            }
        }
    })

    router.POST("/user", func(c *gin.Context) {
        var u User
        if err := c.ShouldBind(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
            return
        }
        users = append(users, u)
        c.Status(http.StatusOK)
    })

    router.PUT("/user/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
            return
        }

        var u User
        if err := c.ShouldBind(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
            return
        }

        for i := 0; i < len(users); i++ {
            if users[i].ID == id {
                users[i].Name = u.Name
                break
            }
        }
        c.Status(http.StatusNoContent)
    })

    router.DELETE("/user/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
            return
        }

        for i, user := range users {
            if user.ID == id {
                users = append(users[0:i], users[i+1:]...)
                break
            }
        }

        c.Status(http.StatusNoContent)
    })

    router.Run(":8080")
}

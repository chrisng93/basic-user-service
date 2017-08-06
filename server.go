package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "github.com/chrisng93/basic-user-service/controllers"
)

func main() {
    r := httprouter.New()

    uc := controllers.NewUserController()

    r.GET("/user/:id", uc.GetUser)
    r.POST("/user", uc.CreateUser)
    r.DELETE("/user/:id", uc.RemoveUser)

    http.ListenAndServe("localhost:3000", r)
}

package main

import (
    "net/http"

    "github.com/julienschmidt/httprouter"
    "github.com/chrisng93/basic-user-service/controllers"

    "gopkg.in/mgo.v2"
)

func main() {
    r := httprouter.New()

    uc := controllers.NewUserController(getSession())

    r.GET("/user/:id", uc.GetUser)
    r.POST("/user", uc.CreateUser)
    r.DELETE("/user/:id", uc.RemoveUser)

    http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session {
    s, err := mgo.Dial("mongodb://localhost")

    if err != nil {
        panic(err)
    }
    return s
}

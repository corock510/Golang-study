package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        v := 42 // change me!
        fmt.Printf("v is of type %T\n", v)

        c.JSON(200, gin.H{
            "message": "pong",
            "v": v,
        })
    })
    r.Run() // デフォルトで :8080 でリッスン
}
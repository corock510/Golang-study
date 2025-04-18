package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "emo-tracking/models" // models パッケージをインポート
)

var db *gorm.DB

func initDB() {
    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
        "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
        os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Database connected!")
}

func main() {
    initDB()

    r := gin.Default()

    r.GET("/todos", func(c *gin.Context) {
        var todos []models.Todo
        if err := db.Find(&todos).Error; err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, todos)
    })

    r.POST("/todos", func(c *gin.Context) {
        var todo models.Todo
        if err := c.ShouldBindJSON(&todo); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        if err := db.Create(&todo).Error; err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(201, todo)
    })

    r.PUT("/todos/:id", func(c *gin.Context) {
        id := c.Param("id")
        var todo models.Todo
        if err := db.First(&todo, id).Error; err != nil {
            c.JSON(404, gin.H{"error": "Todo not found"})
            return
        }
        if err := c.ShouldBindJSON(&todo); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        db.Save(&todo)
        c.JSON(200, todo)
    })

    r.DELETE("/todos/:id", func(c *gin.Context) {
        id := c.Param("id")
        if err := db.Delete(&models.Todo{}, id).Error; err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "Todo deleted"})
    })

    r.Run() // デフォルトで :8080 でリッスン
}
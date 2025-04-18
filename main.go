package main

import (
    "emo-tracking/api/routes"
    "emo-tracking/api/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func main() {
    // ...existing code...
    db := setupDatabase() // 既存のDB初期化コード
    todoService := services.NewTodoService(db)

    r := gin.Default()
    routes.RegisterTodoRoutes(r, todoService)
    // ...existing code...
}
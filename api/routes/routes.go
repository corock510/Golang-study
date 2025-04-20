package routes

import (
    "emo-tracking/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    // TODO関連のルート
    todoController := controllers.NewTodoController(db)
    r.GET("/todos", todoController.GetTodos)
    r.POST("/todos", todoController.CreateTodo)
    r.PUT("/todos/:id", todoController.UpdateTodo)
    r.DELETE("/todos/:id", todoController.DeleteTodo)
}
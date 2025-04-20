package controllers

import (
    "emo-tracking/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

type TodoController struct {
    service *services.TodoService
}

func NewTodoController(db *gorm.DB) *TodoController {
    return &TodoController{
        service: services.NewTodoService(db),
    }
}

func (ctrl *TodoController) GetTodos(c *gin.Context) {
    todos, err := ctrl.service.GetTodos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, todos)
}

func (ctrl *TodoController) CreateTodo(c *gin.Context) {
    var input services.TodoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    todo, err := ctrl.service.CreateTodo(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, todo)
}

func (ctrl *TodoController) UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var input services.TodoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    todo, err := ctrl.service.UpdateTodo(id, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    if err := ctrl.service.DeleteTodo(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
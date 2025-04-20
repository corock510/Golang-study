package services

import (
    "emo-tracking/models"
    "gorm.io/gorm"
)

type TodoService struct {
    db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
    return &TodoService{db: db}
}

type TodoInput struct {
    Title       string `json:"title"`
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
    var todos []models.Todo
    if err := s.db.Find(&todos).Error; err != nil {
        return nil, err
    }
    return todos, nil
}

func (s *TodoService) CreateTodo(input TodoInput) (models.Todo, error) {
    todo := models.Todo{
        Title:       input.Title,
    }
    if err := s.db.Create(&todo).Error; err != nil {
        return models.Todo{}, err
    }
    return todo, nil
}

func (s *TodoService) UpdateTodo(id string, input TodoInput) (models.Todo, error) {
    var todo models.Todo
    if err := s.db.First(&todo, id).Error; err != nil {
        return models.Todo{}, err
    }
    todo.Title = input.Title
    if err := s.db.Save(&todo).Error; err != nil {
        return models.Todo{}, err
    }
    return todo, nil
}

func (s *TodoService) DeleteTodo(id string) error {
    if err := s.db.Delete(&models.Todo{}, id).Error; err != nil {
        return err
    }
    return nil
}
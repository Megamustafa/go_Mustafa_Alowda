package repositories

import (
	"testing"

	"go-todos-api/models"
	"go-todos-api/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	mockRepo := mocks.NewTodoRepository(t)

	input := models.TodoInput{
		Title: "New Todo",
	}
	expectedTodo := models.Todo{
		ID:    1,
		Title: "New Todo",
	}

	mockRepo.On("Create", input).Return(expectedTodo, nil)

	todo, err := mockRepo.Create(input)

	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, todo)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTodos(t *testing.T) {
	mockRepo := mocks.NewTodoRepository(t)

	todos := []models.Todo{
		{ID: 1, Title: "Todo 1"},
		{ID: 2, Title: "Todo 2"},
	}

	mockRepo.On("GetAll").Return(todos, nil)

	result, err := mockRepo.GetAll()

	assert.NoError(t, err)
	assert.ElementsMatch(t, todos, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTodoByID(t *testing.T) {
	mockRepo := mocks.NewTodoRepository(t)

	expectedTodo := models.Todo{
		ID:    1,
		Title: "Todo 1",
	}

	mockRepo.On("GetByID", "1").Return(expectedTodo, nil)

	todo, err := mockRepo.GetByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, todo)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTodo(t *testing.T) {
	mockRepo := mocks.NewTodoRepository(t)

	input := models.TodoInput{
		Title: "Updated Todo",
	}
	updatedTodo := models.Todo{
		ID:    1,
		Title: "Updated Todo",
	}

	mockRepo.On("Update", input, "1").Return(updatedTodo, nil)

	todo, err := mockRepo.Update(input, "1")

	assert.NoError(t, err)
	assert.Equal(t, updatedTodo, todo)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTodo(t *testing.T) {
	mockRepo := mocks.NewTodoRepository(t)

	mockRepo.On("Delete", "1").Return(nil)

	err := mockRepo.Delete("1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

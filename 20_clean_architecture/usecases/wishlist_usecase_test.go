package usecases_test

import (
	"go-wishlist-api/entities"
	"go-wishlist-api/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWishlistRepository struct {
	mock.Mock
}

func (m *MockWishlistRepository) Save(wishlist *entities.Wishlist) error {
	args := m.Called(wishlist)
	return args.Error(0)
}

func (m *MockWishlistRepository) FindByID(id string) (*entities.Wishlist, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Wishlist), args.Error(1)
}

func (m *MockWishlistRepository) Update(wishlist *entities.Wishlist) error {
	args := m.Called(wishlist)
	return args.Error(0)
}

func (m *MockWishlistRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateWishlist(t *testing.T) {
	mockRepo := new(MockWishlistRepository)
	mockRepo.On("Save", mock.Anything).Return(nil)

	usecase := usecases.NewWishlistUsecase(mockRepo)
	err := usecase.CreateWishlist(&entities.Wishlist{
		Name: "Test Wishlist",
	})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindWishlist(t *testing.T) {
	mockRepo := new(MockWishlistRepository)
	mockRepo.On("FindByID", "1").Return(&entities.Wishlist{Name: "Test Wishlist"}, nil)

	usecase := usecases.NewWishlistUsecase(mockRepo)
	wishlist, err := usecase.GetWishlistByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, wishlist)
	assert.Equal(t, "Test Wishlist", wishlist.Name)
	mockRepo.AssertExpectations(t)
}

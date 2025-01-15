package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetUserByID(id string) (Person, error) {
	args := m.Called(id)
	return args.Get(0).(Person), args.Error(1)
}

func TestGetPersonInfo(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockService := new(MockStorage)
		user := Person{
			ID:    "123",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}
		expectedUserInfo := fmt.Sprintf("id = %s, name = %s, email = %s", user.ID, user.Name, user.Email)

		mockService.On("GetUserByID", "123").Return(user, nil)

		userInfo, err := GetUserInfo(mockService, "123")

		assert.NoError(t, err)
		assert.Equal(t, expectedUserInfo, userInfo)

		mockService.AssertExpectations(t)
	})
}

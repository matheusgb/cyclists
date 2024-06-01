package tests

import (
	"testing"

	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/user"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetAllUsersService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockIUser(ctrl)
	service := services.Init(repository)

	t.Run("Success", func(t *testing.T) {
		pagination := &domainsP.Pagination{
			Limit: 1,
			Page:  1,
			Sort:  "",
		}

		response := &domainsP.Pagination{
			Limit:      1,
			Page:       1,
			Sort:       "Id desc",
			TotalRows:  1,
			TotalPages: 1,
			Rows: []entities.User{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Name:  "Test",
					Email: "test@mail.com",
				},
			},
		}

		repository.EXPECT().GetAllUsers(pagination, "").Return(response, nil)
		response, err := service.GetAllUsers(pagination, "")

		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.Rows.([]entities.User)[0].ID)
		assert.Equal(t, "Test", response.Rows.([]entities.User)[0].Name)
		assert.Equal(t, "test@mail.com", response.Rows.([]entities.User)[0].Email)
	})
}

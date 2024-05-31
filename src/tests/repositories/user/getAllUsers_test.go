package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func initGetAllUsersPaginationMockedDomain() *domains.Pagination {
	domain := &domains.Pagination{
		Page:  0,
		Limit: 0,
		Sort:  "",
	}

	return domain
}

func TestGetAllUsersRepository(t *testing.T) {
	db, mock := mocks.MockDatabase()
	domain := initGetAllUsersPaginationMockedDomain()
	t.Run("SuccessWithoutParam", func(t *testing.T) {

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1).AddRow(2))

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"email"}).
				AddRow("test@mail.com").AddRow("test2@mail.com"))

		repository := repositories.Init(db)
		user, err := repository.GetAllUsers(domain, "")

		assert.NoError(t, err)
		assert.Equal(t, 2, len(user.Rows.([]entities.User)))
		assert.Equal(t, "test@mail.com", user.Rows.([]entities.User)[0].Email)
		assert.Equal(t, "test2@mail.com", user.Rows.([]entities.User)[1].Email)
		assert.Equal(t, int64(2), user.TotalRows)
		assert.Equal(t, 1, user.TotalPages)
		assert.Equal(t, 1, user.Page)
		assert.Equal(t, 10, user.Limit)
		assert.Equal(t, "Id desc", user.Sort)
	})

	t.Run("SuccessWithParam", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs("%test@mail.com%").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1))

		mock.ExpectQuery("SELECT").
			WithArgs("%test@mail.com%").
			WillReturnRows(sqlmock.NewRows([]string{"email"}).
				AddRow("test@mail.com"))

		repository := repositories.Init(db)
		user, err := repository.GetAllUsers(domain, "test@mail.com")

		assert.NoError(t, err)
		assert.Equal(t, 1, len(user.Rows.([]entities.User)))
		assert.Equal(t, "test@mail.com", user.Rows.([]entities.User)[0].Email)
		assert.Equal(t, int64(1), user.TotalRows)
		assert.Equal(t, 1, user.TotalPages)
		assert.Equal(t, 1, user.Page)
		assert.Equal(t, 10, user.Limit)
		assert.Equal(t, "Id desc", user.Sort)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs("%test%").
			WillReturnError(db.Error)

		mock.ExpectQuery("SELECT").
			WithArgs("%test%").
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		_, err := repository.GetAllUsers(domain, "test")

		assert.Error(t, err)
	})
}

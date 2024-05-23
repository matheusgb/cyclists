package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/bikeEvent"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func initCreateMockedDomain() *domains.BikeEvent {
	request := requests.CreateBikeEvent{
		Name:                  "Test",
		StartPlace:            "Test",
		StartDateRegistration: time.Now(),
		Organizer:             1,
		StartDate:             time.Now(),
		EndDateRegistration:   time.Now(),
	}

	domain := domains.InitCreate(request.Name,
		request.StartPlace, request.AditionalInformation,
		request.StartDate, request.StartDateRegistration,
		request.EndDateRegistration, request.Organizer, request.ParticipantsLimit)

	return domain
}

func TestCreateBikeEventRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initCreateMockedDomain()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			domain.Name,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			domain.StartPlace,
			domain.Organizer,
			domain.AditionalInformation,
			domain.ParticipantsLimit,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "name", "start_place", "organizer"},
			).AddRow(
				1,
				domain.Name,
				domain.StartPlace,
				domain.Organizer,
			),
		)

	repository := repositories.Init(db)
	bikeEvent, err := repository.CreateBikeEvent(*domain)

	assert.NoError(t, err)
	assert.Equal(t, domain.Name, bikeEvent.Name)
	assert.Equal(t, domain.StartPlace, bikeEvent.StartPlace)
	assert.Equal(t, domain.Organizer, bikeEvent.Organizer)
}

func TestCreateBikeEventRepositoryError(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initCreateMockedDomain()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			domain.Name,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			domain.StartPlace,
			domain.Organizer,
			domain.AditionalInformation,
			domain.ParticipantsLimit,
		).
		WillReturnError(db.Error)
	mock.ExpectCommit()

	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnError(db.Error)

	repository := repositories.Init(db)
	_, err := repository.CreateBikeEvent(*domain)

	assert.Error(t, err)
}

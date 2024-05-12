package tests

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/bikeEvent"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func initUpdateMockedDomain() *domains.BikeEvent {
	request := requests.UpdateBikeEvent{
		Name:                  "Test",
		StartPlace:            "Test",
		StartDateRegistration: time.Now(),
		Organizer:             1,
		StartDate:             time.Now(),
		EndDateRegistration:   time.Now(),
	}

	domain := domains.InitUpdate("1", request.Name,
		request.StartPlace, request.AditionalInformation,
		request.StartDate, request.StartDateRegistration,
		request.EndDateRegistration, request.ParticipantsLimit, request.Organizer)
	return domain
}

func TestUpdateUserRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initUpdateMockedDomain()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
		WithArgs(sqlmock.AnyArg(), domain.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), domain.StartPlace, domain.Organizer, domain.ID, domain.Organizer).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
		WithArgs(domain.ID, domain.Organizer).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, domain.Name))

	repository := repositories.Init(db)
	bikeEvent, err := repository.UpdateBikeEvent(*domain, 1)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), bikeEvent.ID)
	assert.Equal(t, domain.Name, bikeEvent.Name)
}

func TestUpdateUserRepositoryError(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initUpdateMockedDomain()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
		WithArgs(sqlmock.AnyArg(), domain.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), domain.StartPlace, domain.Organizer, domain.ID, domain.Organizer).
		WillReturnError(db.Error)
	mock.ExpectCommit()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
		WithArgs(domain.ID, domain.Organizer).
		WillReturnError(db.Error)

	repository := repositories.Init(db)
	_, err := repository.UpdateBikeEvent(*domain, 1)

	assert.Error(t, err)
}

package domains

import "time"

type BikeEvent struct {
	ID                    string
	Name                  string
	StartDate             time.Time
	StartDateRegistration time.Time
	EndDateRegistration   time.Time
	StartPlace            string
	Organizer             uint
	AditionalInformation  *string
	ParticipantsLimit     *uint
}

func InitCreate(name, startPlace string, aditionalInformation *string, startDate, startDateRegistration, endDateRegistration time.Time, organizer uint, participantsLimit *uint) *BikeEvent {
	return &BikeEvent{
		Name:                  name,
		StartDate:             startDate,
		StartDateRegistration: startDateRegistration,
		EndDateRegistration:   endDateRegistration,
		StartPlace:            startPlace,
		Organizer:             organizer,
		AditionalInformation:  aditionalInformation,
		ParticipantsLimit:     participantsLimit,
	}
}

func InitUpdate(id, name, startPlace string, aditionalInformation *string, startDate, startDateRegistration, endDateRegistration time.Time, participantsLimit *uint) *BikeEvent {
	return &BikeEvent{
		ID:                    id,
		Name:                  name,
		StartDate:             startDate,
		StartDateRegistration: startDateRegistration,
		EndDateRegistration:   endDateRegistration,
		StartPlace:            startPlace,
		AditionalInformation:  aditionalInformation,
		ParticipantsLimit:     participantsLimit,
	}
}

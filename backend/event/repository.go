package event

import (
	"rsvp/person"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(e Event) (*Event, error)
	Get(id string) (*Event, error)
	GetAll() ([]Event, error)
	Delete(id string) (*Event, error)
	Update(id string, details Event) (*Event, error)
	GetAttendance(eventID string) (*EventAttendance, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Create(e Event) (*Event, error) {
	id := uuid.New()
	e.ID = id.String()

	err := repo.db.Create(&e).Error

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (repo *repository) Get(id string) (*Event, error) {
	event := Event{}

	err := repo.db.First(&event, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (repo *repository) GetAll() ([]Event, error) {
	event := []Event{}

	err := repo.db.Find(&event).Error
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (repo *repository) Delete(id string) (*Event, error) {
	event := Event{}

	err := repo.db.Delete(&event, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &event, nil
}

type PersonAttendance struct {
	person.Person
	Going          string
	BringingFriend string
	InvitationID   string
}

type EventAttendance struct {
	Event
	Attendance []PersonAttendance
}

func (repo *repository) GetAttendance(eventID string) (*EventAttendance, error) {
	event, err := repo.Get(eventID)
	if err != nil {
		return nil, err
	}

	attendance := EventAttendance{Event: *event, Attendance: []PersonAttendance{}}

	err = repo.db.Raw(`SELECT p.*,
	i.id as invitation_id,
	r.going,
	r.bringing_friend
	FROM people p
	INNER JOIN (
		SELECT * FROM invitations
		WHERE event_id = ?
		AND deleted_at IS NULL
	) i
	ON p.id = i.person_id

	LEFT JOIN (
		SELECT r.* FROM (
			SELECT max(created_at) created_at, invitation_id
			FROM rsvps
			WHERE event_id = ?
			GROUP BY invitation_id
		) AS newest
		INNER JOIN rsvps r
		ON newest.created_at = r.created_at
		AND newest.invitation_id = r.invitation_id
	) r

	ON r.invitation_id = i.id
	`, eventID, eventID).Scan(&attendance.Attendance).Error

	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

func (repo *repository) Update(id string, details Event) (*Event, error) {
	details.ID = ""

	match := Event{}

	err := repo.db.First(&match, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Where("id = ?", id).Updates(&details).Error

	if err != nil {
		return nil, err
	}

	return &details, nil
}

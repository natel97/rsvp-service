package event

import (
	"fmt"
	"rsvp/person"
	timeoption "rsvp/time_option"
	timeselection "rsvp/time_selection"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=event
type Repository interface {
	Create(e Event) (*Event, error)
	Get(id string) (*Event, error)
	GetAll() ([]Event, error)
	Delete(id string) (*Event, error)
	Update(id string, details Event) (*Event, error)
	GetAttendance(eventID string) (*EventAttendance, error)
	GetTimeOptionData(eventId string, invitationId string) ([]*TimeOption, error)
	GetEventsBetween(hoursStart uint, hoursEnd uint) ([]Event, error)
	GetUnmarkedStaleEvents() ([]Event, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetEventsBetween(hoursStart uint, hoursEnd uint) ([]Event, error) {
	events := []Event{}
	durationStart, _ := time.ParseDuration(fmt.Sprintf("%dh", hoursStart))
	durationEnd, _ := time.ParseDuration(fmt.Sprintf("%dh", hoursEnd))
	now := time.Now().UTC()
	startTime := now.Add(durationStart)
	endTime := now.Add(durationEnd)

	fmt.Println(time.Now(), "checking dates between ", startTime, " ", endTime)
	err := repo.db.Find(&events, "date BETWEEN ? AND ?", startTime, endTime).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return events, nil
}

func (repo *repository) GetUnmarkedStaleEvents() ([]Event, error) {
	events := []Event{}
	now := time.Now().UTC()
	stale := GetStateID(STALE)
	err := repo.db.Find(&events, "date < ? AND state != ?", now, stale).Error
	if err != nil {
		return nil, err
	}

	return events, nil
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
	Attendance  []PersonAttendance
	TimeOptions []*TimeOption
}

func getTimeOptionDetail(db *gorm.DB, timeOptionID string, invitationID string) (*TimeOption, error) {
	selections := []timeselection.TimeSelection{}

	err := db.Table("time_selections").Find(&selections, "time_option_id = ?", timeOptionID).Error
	if err != nil {
		return nil, err
	}

	to := TimeOption{}

	for _, selection := range selections {
		if selection.Acceptable == nil {
			continue
		}

		if selection.InvitationID == invitationID {
			to.IsDownvote = !*selection.Acceptable
			to.IsUpvote = *selection.Acceptable
		}

		if *selection.Acceptable {
			to.Upvote += 1
		} else {
			to.Downvote += 1
		}
	}
	return &to, nil
}

func (repo *repository) GetTimeOptionData(eventId string, invitationId string) ([]*TimeOption, error) {
	timeOptions := []timeoption.TimeOption{}

	err := repo.db.Table("time_options").Find(&timeOptions, "event_id = ? AND deleted_at IS NULL", eventId).Error
	if err != nil {
		return nil, err
	}

	options := []*TimeOption{}

	for _, option := range timeOptions {
		opt, err := getTimeOptionDetail(repo.db, option.ID, invitationId)
		if err != nil {
			return nil, err
		}
		opt.Time = option.Time
		opt.ID = option.ID
		options = append(options, opt)
	}

	return options, nil
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

	timeOptions, err := repo.GetTimeOptionData(eventID, "")
	if err != nil {
		return nil, err
	}

	attendance.TimeOptions = timeOptions

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

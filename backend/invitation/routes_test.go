package invitation_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"rsvp/event"
	"rsvp/invitation"
	"rsvp/invitation/types"
	"rsvp/notifications"
	"rsvp/person"
	"rsvp/rsvp"
	timeoption "rsvp/time_option"
	timeselection "rsvp/time_selection"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

var invitationRepository *types.MockRepository
var eventRepository *event.MockRepository
var rsvpRepository *rsvp.MockRepository
var notificationService *notifications.MockService
var personRepository *person.MockRepository
var timeOptionRepository *timeoption.MockRepository
var timeSelectionRepository *timeselection.MockRepository

func buildRoute(t *testing.T) invitation.Controller {
	ctrl := gomock.NewController(t)

	invitationRepository = types.NewMockRepository(ctrl)
	eventRepository = event.NewMockRepository(ctrl)
	rsvpRepository = rsvp.NewMockRepository(ctrl)
	notificationService = notifications.NewMockService(ctrl)
	personRepository = person.NewMockRepository(ctrl)
	timeOptionRepository = timeoption.NewMockRepository(ctrl)
	timeSelectionRepository = timeselection.NewMockRepository(ctrl)

	return *invitation.NewController(invitationRepository,
		eventRepository,
		rsvpRepository,
		notificationService,
		personRepository,
		timeOptionRepository,
		timeSelectionRepository)
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func getTimeOptions(durations ...string) []*event.TimeOption {
	timeOptions := []*event.TimeOption{}
	now := time.Now()
	for _, dur := range durations {
		duration, _ := time.ParseDuration((dur))
		newTime := now.Add(duration)
		option := event.TimeOption{
			ID:         dur,
			Time:       &newTime,
			Upvote:     3,
			Downvote:   3,
			IsUpvote:   true,
			IsDownvote: false,
		}
		timeOptions = append(timeOptions, &option)
	}

	return timeOptions
}

func TestGetReservation(t *testing.T) {
	router := setUpRouter()
	controller := buildRoute(t)
	controller.HandleRoutes(router.Group(""))
	req, _ := http.NewRequest("GET", "/id", nil)
	w := httptest.NewRecorder()

	invite := types.Invitation{
		ID:       "id",
		PersonID: "person",
		EventID:  "event",
	}

	invitationRepository.EXPECT().Get("id").Return(&invite, nil)

	now := time.Now()
	event := event.Event{
		ID:           "event",
		Title:        "test event",
		Street:       "111 test street",
		City:         "Melbourne, VIC 3000",
		Description:  "test description",
		InternalNote: "hidden note",
		Date:         &now,
		State:        event.GetStateID(event.PLANNING),
	}

	eventRepository.EXPECT().Get(invite.EventID).Return(&event, nil)

	rsvps := []rsvp.RSVP{
		{Going: "Yes", BringingFriend: "Yes"},
		{Going: "Maybe", BringingFriend: "Maybe"},
		{Going: "No", BringingFriend: "No"},
		{Going: "Maybe", BringingFriend: "No"},
	}
	rsvpRepository.EXPECT().GetEventRSVPs(event.ID).Return(rsvps, nil)

	rsvpYes := uint(2)
	rsvpMaybe := uint(3)
	rsvpNo := uint(1)

	me := rsvp.RSVP{
		Going:          "Yes",
		BringingFriend: "No",
	}
	rsvpRepository.EXPECT().GetLatestRSVPByInvitation(invite.ID).Return(&me, nil)

	notificationService.EXPECT().GetIsSubscribed(invite.ID).Return(false)

	timeOptions := getTimeOptions("1d", "2d", "3d", "4d")
	eventRepository.EXPECT().GetTimeOptionData(event.ID, invite.ID).Return(timeOptions, nil)

	router.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("Status Code Mismatch - Expected %d Got %d", http.StatusOK, w.Result().StatusCode)
	}

	parsed := invitation.GetInvitationResponse{}
	json.Unmarshal(responseData, &parsed)

	if parsed.Title != event.Title {
		t.Fatalf("expected %s got %s", event.Title, parsed.Title)
	}

	if parsed.Street != event.Street {
		t.Fatalf("expected %s got %s", event.Street, parsed.Street)
	}

	if parsed.City != event.City {
		t.Fatalf("expected %s got %s", event.City, parsed.City)
	}

	if parsed.Description != event.Description {
		t.Fatalf("expected %s got %s", event.Description, parsed.Description)
	}

	if parsed.MyAttendance != me.Going {
		t.Fatalf("expected %s got %s", parsed.MyAttendance, me.Going)
	}

	if parsed.Attendance.Yes != rsvpYes {
		t.Fatalf("expected %d got %d", parsed.Attendance.Yes, rsvpYes)
	}

	if parsed.Attendance.Maybe != rsvpMaybe {
		t.Fatalf("expected %d got %d", parsed.Attendance.Maybe, rsvpMaybe)
	}

	if parsed.Attendance.No != rsvpNo {
		t.Fatalf("expected %d got %d", parsed.Attendance.No, rsvpNo)
	}

	t.Log(string(responseData))
}

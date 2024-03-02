package invitation

import (
	"errors"
	"fmt"
	"net/http"
	"rsvp/event"
	"rsvp/invitation/types"
	"rsvp/notifications"
	"rsvp/person"
	"rsvp/rsvp"
	timeoption "rsvp/time_option"
	timeselection "rsvp/time_selection"
	"time"

	ical "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron"
	"gorm.io/gorm"
)

type Controller struct {
	repository              *repository
	eventRepository         event.Repository
	rsvpRepository          rsvp.Repository
	notifications           *notifications.Service
	personRepository        person.Repository
	timeOptionRepository    timeoption.Repository
	timeSelectionRepository timeselection.Repository
}

func getRSVPMessage(rsvp *rsvp.RSVP) string {
	if rsvp == nil {
		return "You have not RSVP'd yet 😥️"
	}

	messages := map[string]string{
		"Yes":   "See you soon!",
		"No":    "It's not too late to change your mind if you'd like to RSVP.",
		"Maybe": "We hope we'll see you, but feel free to update your RSVP!",
	}

	return messages[rsvp.Going]
}

func (ctrl *Controller) notifyUpcomingEvent(hoursBefore uint, hourRange uint) func() {
	return func() {
		fmt.Println("Running cron: Notify Upcoming Event")
		events, err := ctrl.eventRepository.GetEventsBetween(hoursBefore, hoursBefore+hourRange)
		if err != nil {
			fmt.Println("Error in notification", err)
		}

		for _, event := range events {
			fmt.Println("Sending out notifications for event ", event.ID)
			invitations, err := ctrl.repository.GetByEvent(event.ID)
			if err != nil {
				fmt.Println("Error getting invites for upcoming event", event.ID)
				continue
			}

			for _, invite := range invitations {
				rsvp, err := ctrl.rsvpRepository.GetLatestRSVPByInvitation(invite.ID)
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					fmt.Println("Error getting rsvp to notify before event: ", event.Title, invite.ID)
					continue
				}

				rsvpMessage := getRSVPMessage(rsvp)
				message := fmt.Sprintf("Upcoming event %s! %s", event.Title, rsvpMessage)
				ctrl.notifications.NotifyInvite("push-notify", message, invite.ID)
			}
		}
	}
}

func (ctrl *Controller) closeStaleEvents() {
	events, err := ctrl.eventRepository.GetUnmarkedStaleEvents()
	fmt.Println("Checking stale events")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, e := range events {
		e.SetEventState(event.STALE)
		_, err := ctrl.eventRepository.Update(e.ID, e)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Updated stale event", e.ID)
	}
}

func (ctrl *Controller) AddCrons(cron *cron.Cron) {
	job := ctrl.notifyUpcomingEvent(5, 1)
	cron.AddFunc("@hourly", job)
	cron.AddFunc("@daily", ctrl.closeStaleEvents)
}

func (ctrl *Controller) get(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	invitation, err := ctrl.repository.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	ev, err := ctrl.eventRepository.Get(invitation.EventID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	attendance := Attendance{}
	responses, err := ctrl.rsvpRepository.GetEventRSVPs(ev.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	for _, response := range responses {
		if response.Going == "Yes" {
			attendance.Yes += 1
		}
		if response.Going == "Maybe" {
			attendance.Maybe += 1
		}
		if response.Going == "No" {
			attendance.No += 1
		}
		if response.Going != "No" && response.BringingFriend == "Yes" {
			attendance.Yes += 1
		}
		if response.Going != "No" && response.BringingFriend == "Maybe" {
			attendance.Maybe += 1
		}
	}

	me, err := ctrl.rsvpRepository.GetLatestRSVPByInvitation(invitation.ID)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	if me == nil {
		me = &rsvp.RSVP{}
	}

	subscribed := ctrl.notifications.GetIsSubscribed(id)
	timeOptions, err := ctrl.eventRepository.GetTimeOptionData(ev.ID, invitation.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	inviteState := ev.GetEventState()
	var date *time.Time
	if inviteState != event.PLANNING {
		date = ev.Date
	}

	response := GetInvitationResponse{
		Title:           ev.Title,
		Date:            date,
		Street:          ev.Street,
		City:            ev.City,
		Attendance:      attendance,
		InvitationState: inviteState,
		MyAttendance:    me.Going,
		MyFriend:        me.BringingFriend,
		Description:     ev.Description,
		Subscribed:      subscribed,
		TimeOptions:     timeOptions,
	}

	ctx.JSON(http.StatusOK, response)
}

type createBody struct {
	EventID   string `json:"eventID"`
	PersonIdD string `json:"personID"`
}

func (ctrl *Controller) create(ctx *gin.Context) {
	body := createBody{}

	err := ctx.BindJSON(&body)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	create := types.Invitation{
		PersonID: body.PersonIdD,
		EventID:  body.EventID,
	}
	val, err := ctrl.repository.Create(create)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, val)
}

func (ctrl *Controller) inviteGroup(ctx *gin.Context) {
	body := InviteGroup{}

	err := ctx.BindJSON(&body)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	err = ctrl.repository.InviteGroup(body)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "")
}

func (ctrl *Controller) post(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	body := rsvp.UpdateRSVP{}
	ctx.BindJSON(&body)

	invitation, err := ctrl.repository.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	response := rsvp.RSVP{
		InvitationID:   id,
		EventID:        invitation.EventID,
		Going:          body.Going,
		BringingFriend: body.BringingFriend,
	}

	_, err = ctrl.rsvpRepository.Create(response)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	go func() {
		person, _ := ctrl.personRepository.Get(invitation.PersonID)
		e, _ := ctrl.eventRepository.Get(invitation.EventID)
		message := fmt.Sprintf("%s %s just RSVP'd to %s. Going %s, Friend %s", person.First, person.Last, e.Title, body.Going, body.BringingFriend)
		ctrl.notifications.NotifyGroup("admin", "push-notify", message, fmt.Sprintf("/admin/event/%s/invite", invitation.EventID))
	}()

	ctx.JSON(http.StatusCreated, "Success")
}

func (ctrl *Controller) getCalendarFile(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	invitation, err := ctrl.repository.Get(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	event, err := ctrl.eventRepository.Get(invitation.EventID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	deref := *event.Date
	endAt := deref.Add(2 * time.Hour)

	cal := ical.NewCalendar()
	cal.SetMethod(ical.MethodRequest)
	calEvent := cal.AddEvent(invitation.ID)
	calEvent.SetDtStampTime(event.CreatedAt)
	calEvent.SetLocation(fmt.Sprintf("%s, %s", event.Street, event.City))
	calEvent.SetStartAt(*event.Date)
	calEvent.SetEndAt(endAt)
	calEvent.SetSummary(event.Title)
	calEvent.SetURL(fmt.Sprintf("%s/invitation/%s", ctx.Request.Header.Get("Origin"), id))
	calEvent.SetDescription(fmt.Sprintf("%s\n\n%s/invitation/%s", event.Description, ctx.Request.Header.Get("Origin"), id))

	calString := cal.Serialize()
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.ics", event.Title))
	ctx.Data(http.StatusOK, "application/octet-stream", []byte(calString))
}

func (ctrl *Controller) delete(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	err := ctrl.repository.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.Status(http.StatusAccepted)
}

type ReservationSubscriptionInput struct {
	Subscription string `json:"subscription"`
}

func (ctrl *Controller) unsubscribe(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := ctrl.notifications.RemoveByInvitation(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusAccepted, "deleted")
}

func (ctrl *Controller) selectTime(ctx *gin.Context) {
	invitationID := ctx.Param("id")
	timeOptionID := ctx.Param("time-id")
	body := timeselection.TimeSelection{}
	ctx.BindJSON(&body)
	body.TimeOptionID = timeOptionID
	body.InvitationID = invitationID

	err := ctrl.timeSelectionRepository.UpdateSelection(&body)

	if err != nil {
		fmt.Println(err, "error making time selection")
		ctx.JSON(http.StatusBadRequest, "db error")
		return
	}

	ctx.JSON(http.StatusAccepted, body)
}

func (ctrl *Controller) subscribe(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	rsi := ReservationSubscriptionInput{}
	err := ctx.BindJSON(&rsi)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	ctrl.notifications.RegisterForInvitation(id, rsi.Subscription, "invited")
	go func() {
		ctrl.notifications.Notify(rsi.Subscription, "push-notify", "Subscription Successful!", "/invitation/"+id)
		ctrl.notifications.NotifyGroup("admin", "push-notify", "New Subscriber to push!", "/admin")
	}()
}

func (ctrl *Controller) HandleRoutes(group *gin.RouterGroup) {
	group.GET("/:id", ctrl.get)
	group.POST("/:id/rsvp", ctrl.post)
	group.GET("/:id/download", ctrl.getCalendarFile)
	group.POST("/:id/subscribe", ctrl.subscribe)
	group.DELETE("/:id/subscribe", ctrl.unsubscribe)
	group.PUT("/:id/time-selection/:time-id", ctrl.selectTime)
}

func (ctrl *Controller) HandleAdminRoutes(group *gin.RouterGroup) {
	group.POST("", ctrl.create)
	group.DELETE(":id", ctrl.delete)
	group.POST("group", ctrl.inviteGroup)
}

func NewController(repository *repository, eventRepository event.Repository, rsvpRepository rsvp.Repository, notifications *notifications.Service, personRepository person.Repository, timeOptionRepository timeoption.Repository, timeSelectionRepository timeselection.Repository) *Controller {
	return &Controller{
		repository:              repository,
		eventRepository:         eventRepository,
		rsvpRepository:          rsvpRepository,
		notifications:           notifications,
		personRepository:        personRepository,
		timeOptionRepository:    timeOptionRepository,
		timeSelectionRepository: timeSelectionRepository,
	}
}

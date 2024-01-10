package invitation

import (
	"errors"
	"fmt"
	"net/http"
	"rsvp/event"
	"rsvp/rsvp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	repository      *repository
	eventRepository event.Repository
	rsvpRepository  rsvp.Repository
}

func (ctrl *Controller) get(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	invitation, err := ctrl.repository.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	if invitation == nil {
		ctx.JSON(http.StatusNotFound, "Not Found")
		fmt.Println(err)
		return
	}

	event, err := ctrl.eventRepository.Get(invitation.EventID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	if event == nil {
		ctx.JSON(http.StatusNotFound, "Event Not Found")
		fmt.Println(err)
		return
	}

	attendance := Attendance{}
	responses, err := ctrl.rsvpRepository.GetEventRSVPs(event.ID)
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

	response := GetInvitationResponse{
		Title:        event.Title,
		Date:         event.Date,
		Street:       event.Street,
		City:         event.City,
		Attendance:   attendance,
		MyAttendance: me.Going,
		MyFriend:     me.BringingFriend,
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *Controller) post(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	body := rsvp.UpdateRSVP{}
	ctx.BindJSON(&body)

	invitation, err := ctrl.repository.Get(id)

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

	ctx.JSON(http.StatusCreated, "Success")
}

func (ctrl *Controller) HandleRoutes(group *gin.RouterGroup) {
	group.GET("/:id", ctrl.get)
	group.POST("/:id/rsvp", ctrl.post)
}

func NewController(repository *repository, eventRepository event.Repository, rsvpRepository rsvp.Repository) *Controller {
	return &Controller{
		repository:      repository,
		eventRepository: eventRepository,
		rsvpRepository:  rsvpRepository,
	}
}

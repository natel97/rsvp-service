package event

import (
	"fmt"
	"net/http"
	types "rsvp/invitation/types"
	"rsvp/notifications"
	"rsvp/person"
	timeoption "rsvp/time_option"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Controller struct {
	repository           *repository
	invitationRepository types.Repository
	personRepository     person.Repository
	notifify             *notifications.Service
	timeOptionRepository timeoption.Repository
}

func (ctrl *Controller) getAll(ctx *gin.Context) {
	vals, err := ctrl.repository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, vals)
}

func (ctrl *Controller) get(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	val, err := ctrl.repository.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, val)
}

func (ctrl *Controller) post(ctx *gin.Context) {
	event := Event{}

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	vals, err := ctrl.repository.Create(event)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, vals)
}

func (ctrl *Controller) update(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	existingEvent, err := ctrl.repository.Get(id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		fmt.Println(err)
		return
	}

	event := Event{}

	err = ctx.ShouldBindBodyWith(&event, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	rawMap := map[string]interface{}{}

	err = ctx.ShouldBindBodyWith(&rawMap, binding.JSON)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	vals, err := ctrl.repository.Update(id, event)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}

	updatedItems := ""

	for k, i := range rawMap {
		if k != "InternalNote" {
			updatedItems = fmt.Sprintf("%s, %s: %s", updatedItems, k, i)
		}
	}

	message := fmt.Sprintf("Event %s updated%s", existingEvent.Title, updatedItems)

	if len(updatedItems) > 0 {
		go ctrl.notifify.NotifyEvent(id, "push-notify", message, "/")
	}

	ctx.JSON(http.StatusOK, vals)
}

func (ctrl *Controller) delete(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	existing, err := ctrl.repository.Get("id")
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	vals, err := ctrl.repository.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctrl.notifify.NotifyEvent(id, "push-notify", fmt.Sprintf("Event %s has been deleted", existing.Title), "/")

	ctx.JSON(http.StatusOK, vals)
}

func (ctrl *Controller) getPeopleInvitedToEvent(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	people, err := ctrl.personRepository.GetHavingEvent(id)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, people)
}

func (ctrl *Controller) getAttendance(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	attendance, err := ctrl.repository.GetAttendance(id)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, attendance)
}

func (ctrl *Controller) addTimeOption(ctx *gin.Context) {
	body := timeoption.TimeOption{}
	err := ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println("error bind json: ", err)
		return
	}

	id := ctx.Param("id")
	body.EventID = id

	if body.Time == nil {
		ctx.JSON(http.StatusBadRequest, "Missing Time field")
		fmt.Println("Missing Time field")
		return
	}

	err = ctrl.timeOptionRepository.CreateTimeOption(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println("error bind json: ", err)
		return
	}

	ctx.JSON(http.StatusCreated, body)
}

func (ctrl *Controller) deleteTimeOption(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	err := ctrl.timeOptionRepository.DeleteTimeOption(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println("Error: ", err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

func (ctrl *Controller) establishTime(ctx *gin.Context) {
	id := ctx.Param("id")
	timeID := ctx.Param("time-id")
	timeOption, err := ctrl.timeOptionRepository.GetTimeOption(timeID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println("Error setting time", err)
		return
	}

	event := Event{Date: timeOption.Time}
	event.SetEventState(UPCOMING)

	_, err = ctrl.repository.Update(id, event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println("Error setting time", err)
		return
	}

	fullEvent, err := ctrl.repository.Get(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println("Error setting time", err)
		return
	}

	ctrl.notifify.NotifyEvent(id, "push-notify", fmt.Sprintf("%s update - Time set!", fullEvent.Title), "/")
}

func (ctrl *Controller) HandleRoutes(group *gin.RouterGroup) {
	// Nothing public yet
}

func (ctrl *Controller) HandleAdminRoutes(group *gin.RouterGroup) {
	group.GET("", ctrl.getAll)
	group.POST("", ctrl.post)
	group.PUT("/:id", ctrl.update)
	group.GET("/:id", ctrl.get)
	group.DELETE("/:id", ctrl.delete)
	group.GET("/:id/people", ctrl.getPeopleInvitedToEvent)
	group.GET("/:id/attendance", ctrl.getAttendance)
	group.POST(":id/time-option", ctrl.addTimeOption)
	group.PUT(":id/select-time/:time-id", ctrl.establishTime)
	group.DELETE("/time-option/:id", ctrl.deleteTimeOption)
}

func NewController(repository *repository, invitationRepository types.Repository, personRepository person.Repository, notify *notifications.Service, timeOptionRepository timeoption.Repository) *Controller {
	return &Controller{
		repository:           repository,
		invitationRepository: invitationRepository,
		personRepository:     personRepository,
		notifify:             notify,
		timeOptionRepository: timeOptionRepository,
	}
}

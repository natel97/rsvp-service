package group

import (
	"net/http"
	persongroup "rsvp/person_group"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository   *repository
	pgRepository persongroup.Repository
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
	group := Group{}

	err := ctx.ShouldBindJSON(&group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	vals, err := ctrl.repository.Create(group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, vals)
}

func (ctrl *Controller) update(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	group := Group{}

	err := ctx.ShouldBindJSON(&group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	vals, err := ctrl.repository.Update(id, group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, vals)
}

func (ctrl *Controller) delete(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	vals, err := ctrl.repository.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, vals)
}

func (ctrl *Controller) addToGroup(ctx *gin.Context) {
	groupID, _ := ctx.Params.Get("id")
	personID, _ := ctx.Params.Get("personID")

	personGroup := &persongroup.PersonGroup{
		PersonID: personID,
		GroupID:  groupID,
	}

	err := ctrl.pgRepository.CreateUserGroup(personGroup)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, personGroup)
}

func (ctrl *Controller) deleteFromGroup(ctx *gin.Context) {
	groupID, _ := ctx.Params.Get("id")
	personID, _ := ctx.Params.Get("personID")
	err := ctrl.pgRepository.DeleteUserGroup(groupID, personID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}

func (ctrl *Controller) getUsersWithGroup(ctx *gin.Context) {
	groupID, _ := ctx.Params.Get("id")
	users, err := ctrl.pgRepository.GetUsersInGroup(groupID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (ctrl *Controller) HandleRoutes(group *gin.RouterGroup) {
	group.GET("", ctrl.getAll)
	group.POST("", ctrl.post)
	group.PUT("/:id", ctrl.update)
	group.GET("/:id", ctrl.get)
	group.DELETE("/:id/person/:personID", ctrl.deleteFromGroup)
	group.DELETE("/:id", ctrl.delete)
	group.POST("/:id/person/:personID", ctrl.addToGroup)
	group.GET("/:id/person", ctrl.getUsersWithGroup)
}

func NewController(repository *repository, pgRepository persongroup.Repository) *Controller {
	return &Controller{
		repository:   repository,
		pgRepository: pgRepository,
	}
}

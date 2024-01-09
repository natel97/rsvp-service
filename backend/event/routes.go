package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository *repository
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
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	vals, err := ctrl.repository.Create(event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, vals)
}

func (ctrl *Controller) update(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	event := Event{}

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	vals, err := ctrl.repository.Update(id, event)
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

func (ctrl *Controller) HandleRoutes(group *gin.RouterGroup) {
	group.GET("/", ctrl.getAll)
	group.POST("/", ctrl.post)
	group.PUT("/:id", ctrl.update)
	group.GET("/:id", ctrl.get)
	group.DELETE("/:id", ctrl.delete)
}

func NewController(repository *repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

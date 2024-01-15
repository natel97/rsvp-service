package notifications

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewController(service *Service, repo *Repository) *Controller {
	return &Controller{
		service: service,
		repo:    repo,
	}
}

type Controller struct {
	service *Service
	repo    *Repository
}

type SubscriptionInput struct {
	Subscription string `json:"subscription"`
	Kind         string `json:"kind"`
}

func (ctrl *Controller) subscribe(ctx *gin.Context) {
	subscription := SubscriptionInput{}
	err := ctx.BindJSON(&subscription)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	subscription.Kind = "admin"

	err = ctrl.repo.Create(subscription.Subscription, subscription.Kind)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctrl.service.Notify(subscription.Subscription, "push-notify", "Push Notifications Set Up", "/admin")

	ctx.JSON(http.StatusOK, "OK")
}

func (ctrl Controller) HandleRoutes(routes *gin.RouterGroup) {
	routes.POST("/subscribe", ctrl.subscribe)
}

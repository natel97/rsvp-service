package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"rsvp/event"
	"rsvp/group"
	"rsvp/invitation"
	"rsvp/invitation/types"
	"rsvp/notifications"
	"rsvp/person"
	persongroup "rsvp/person_group"
	"rsvp/rsvp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTestData(personRepository person.Repository, eventRepository event.Repository, invitationRepository types.Repository, run bool) {
	if !run {
		return
	}

	p, _ := personRepository.Create(person.Person{
		First: "John",
		Last:  "Doe",
	})

	now := time.Now()

	e, _ := eventRepository.Create(event.Event{
		Title:  "Test Event",
		Date:   &now,
		Street: "111 Flinders Street",
		City:   "Melbourne, VIC 3000",
	})

	invitation, _ := invitationRepository.Create(types.Invitation{
		PersonID: p.ID,
		EventID:  e.ID,
	})

	fmt.Println(invitation)
}

func getConfig() (map[string]string, error) {
	config := viper.New()
	config.SetConfigName(".env.local")
	config.SetConfigType("env")
	config.AddConfigPath(".")

	c := map[string]string{}

	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = config.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&event.Event{}, &types.Invitation{}, &person.Person{}, &rsvp.RSVP{}, &group.Group{}, &persongroup.PersonGroup{}, &notifications.Subscription{}, &notifications.InvitationSubscription{})

	eventRepository := event.NewRepository(db)
	rsvpRepository := rsvp.NewRepository(db)
	invitationRepository := invitation.NewRepository(db)
	personRepository := person.NewRepository(db)
	groupRepository := group.NewRepository(db)
	pgRepository := persongroup.NewRepository(db)

	notificationRepo := notifications.NewRepository(db)

	c, err := getConfig()
	if err != nil {
		fmt.Println(err)
		panic("config error" + err.Error())
	}

	notifyService := notifications.NewService(c["vapid_private_key"], c["vapid_public_key"], notificationRepo)
	invitationController := invitation.NewController(invitationRepository, eventRepository, rsvpRepository, notifyService, personRepository)
	eventController := event.NewController(eventRepository, invitationRepository, personRepository, notifyService)
	personController := person.NewController(personRepository)
	groupController := group.NewController(groupRepository, pgRepository)

	createTestData(personRepository, eventRepository, invitationRepository, false)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	adminRoutes := server.Group("admin")
	adminRoutes.Use(cors.Default())

	adminRoutes.Use(func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth != c["api_key"] {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
	})

	notifyCtrl := notifications.NewController(notifyService, notificationRepo)

	invitationController.HandleRoutes(server.Group("invitation"))
	invitationController.HandleAdminRoutes(adminRoutes.Group("invitation"))
	eventController.HandleRoutes(adminRoutes.Group("event"))
	personController.HandleRoutes(adminRoutes.Group("people"))
	groupController.HandleRoutes(adminRoutes.Group("group"))
	notifyCtrl.HandleRoutes(adminRoutes.Group("notify"))

	fmt.Println("Starting, http://localhost:9083")
	server.Run(fmt.Sprintf(":%d", 9083))
}

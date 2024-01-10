package main

import (
	"fmt"
	"log/slog"
	"os"
	"rsvp/event"
	"rsvp/invitation"
	"rsvp/person"
	"rsvp/rsvp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTestData(personRepository person.Repository, eventRepository event.Repository, invitationRepository invitation.Repository, run bool) {
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

	invitation, _ := invitationRepository.Create(invitation.Invitation{
		UserID:  p.ID,
		EventID: e.ID,
	})

	fmt.Println(invitation)
}

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&event.Event{}, &invitation.Invitation{}, &person.Person{}, &rsvp.RSVP{})

	eventRepository := event.NewRepository(db)
	rsvpRepository := rsvp.NewRepository(db)
	invitationRepository := invitation.NewRepository(db)
	personRepository := person.NewRepository(db)
	invitationController := invitation.NewController(invitationRepository, eventRepository, rsvpRepository)

	createTestData(personRepository, eventRepository, invitationRepository, false)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(gin.Recovery())

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	invitationController.HandleRoutes(server.Group("invitation"))
	// eventRoutes.HandleRoutes(server.Group("admin/event"))

	fmt.Println("Starting, http://localhost:9083")

	server.Run(fmt.Sprintf(":%d", 9083))
}

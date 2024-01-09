package main

import (
	"fmt"
	"log/slog"
	"os"
	"rsvp/event"
	"rsvp/invitation"
	"rsvp/person"
	"rsvp/rsvp"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&event.Event{}, &invitation.Invitation{}, &person.Person{}, &rsvp.RSVP{})

	_ = event.NewRepository(db)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(gin.Recovery())

	fmt.Println("Starting")
	server.Run(fmt.Sprintf(":%d", 9083))
}

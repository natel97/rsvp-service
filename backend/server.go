package main

import (
	"fmt"
	"log/slog"
	"os"
	"rsvp/event"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	_ = event.NewRepository(db)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(gin.Recovery())

	fmt.Println("Starting")
	server.Run(fmt.Sprintf(":%d", 9083))
}

package main

import (
	"myEchoTest/chapter2/handlers"
	"myEchoTest/chapter2/models"

	"github.com/labstack/echo"
	//_ "github.com/mattn/go-sqlite3"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// Signing Key for our auth middleware
	var signingKey = []byte("superdupersecret!")
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.SigningContextKey, signingKey)
			return next(c)
		}
	})

	// add database to context
	/* db, err := sql.Open("sqlite3", "./service.db")
	if err != nil {
		log.Fatalf("error opening database: %v\n", err)
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.DBContextKey, db)
			return next(c)
		}
	})

	reminderGroup := e.Group("/reminder")
	reminderGroup.Use(middleware.JWT(signingKey))
	reminderGroup.POST("", handlers.CreateReminder)
	*/

	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)
	// Authentication routes
	e.POST("/login", handlers.Login)
	e.POST("/logout", handlers.Logout)
	g := e.Group("/v1")
	g.POST("/login", handlers.Login)
	g.POST("/logout", handlers.Logout)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

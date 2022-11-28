package main

import (
	"log"
	"time"

	"github.com/aurelius15/lf/internal/model"
	"github.com/aurelius15/lf/internal/transaction"
	"github.com/aurelius15/lf/internal/verification"
	"github.com/aurelius15/lf/internal/web"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				go func() {
					log.Println("verify job is started")
					verification.VerifyJob()
					log.Println("verify job is finished")
				}()

				go func() {
					log.Println("transaction job is started")
					transaction.TransJob()
					log.Println("transaction job is finished")
				}()
			}
		}
	}()

	// Echo instance
	e := echo.New()

	e.Validator = &model.CustomValidator{Validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST(web.ApiPrefix+"/user", web.CreateUser)
	e.POST(web.ApiPrefix+"/transaction", web.CreateTransaction)
	e.GET(web.ApiPrefix+"/users", web.AllUsers)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

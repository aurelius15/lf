package web

import (
	"net/http"

	"github.com/aurelius15/lf/internal/model"
	"github.com/aurelius15/lf/internal/storage"
	"github.com/aurelius15/lf/internal/transaction"
	"github.com/aurelius15/lf/internal/verification"
	"github.com/labstack/echo/v4"
)

const (
	ApiPrefix = "/api/v1"
)

// CreateUser is a handler for creating a new user
func CreateUser(c echo.Context) error {
	uReq := new(model.CreateUserRequest)
	if err := c.Bind(uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.CreateUser(uReq.Name)
	_, _ = storage.SaveUser(user)
	verification.VerifyUser(user)

	return c.JSON(http.StatusOK, user)
}

// AllUsers is a handler for getting all users
func AllUsers(c echo.Context) error {
	users, _ := storage.GetAllUsers()

	return c.JSON(http.StatusOK, users)
}

// CreateTransaction is a handler for creating a new transaction
func CreateTransaction(c echo.Context) error {
	tReq := new(model.CreateTransactionRequest)
	if err := c.Bind(tReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(tReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !transaction.BaseCheck(tReq.ReceiverID, tReq.SenderID, tReq.Amount) {
		return echo.NewHTTPError(http.StatusBadRequest, "transaction can't be processed")
	}

	transaction.CreateTransaction(tReq.ReceiverID, tReq.SenderID, tReq.Amount)

	return c.JSON(http.StatusOK, "transaction is created")
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/transactions-service/api/model"
	"github.com/italorfeitosa/transactions-service/pkg/errs"
)

type Account struct{}

func (a *Account) Register(r gin.IRouter) {
	r.POST("/accounts", a.OpenAccount)
	r.GET("/accounts/{accountID}", a.OpenAccount)
}

// swagger:route POST /api/v1/accounts accounts openAccount
//
// Open Account
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Parameters:
//	  + name: Open Account Body
//	    in: body
//	    required: true
//	    type: OpenAccountRequest
//
//	Responses:
//	  201: AccountResponse
//	  400: OpenAccountValidationError
//	  422: OpenAccountBusinessRuleError
//	  500: InternalServerError
func (a *Account) OpenAccount(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, model.ErrorResponse{Error: errs.Error{Message: "TODO: implement open account"}})
}

// swagger:route GET /api/v1/accounts/{accountID} accounts getAccount
//
// Get Account
//
//	Produces:
//	- application/json
//
//	Parameters:
//	  + name: accountID
//	    in: path
//	    required: true
//	    type: string
//
//	Responses:
//	  200: AccountResponse
//	  404: AccountNotFoundResponse
//	  500: InternalServerError
func (a *Account) GetAccount(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, model.ErrorResponse{Error: errs.Error{Message: "TODO: implement get account"}})
}

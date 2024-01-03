package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/transactions-service/api/model"
	"github.com/italorfeitosa/transactions-service/internal/app"
)

type Account struct {
	accountUseCase app.AccountUseCase
}

func NewAccount(accountUseCase app.AccountUseCase) *Account {
	return &Account{accountUseCase}
}

func (a *Account) Register(r gin.IRouter) {
	r.POST("/accounts", a.OpenAccount)
	r.GET("/accounts/:accountID", a.GetAccount)
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
	var req model.OpenAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(ErrMalformedOpenAccountRequest.WithError(err))
		return
	}

	account, err := a.accountUseCase.OpenAccount(c.Request.Context(), app.OpenAccountInput{
		DocumentType:         req.DocumentType,
		DocumentNumber:       req.DocumentNumber,
		AvailableCreditLimit: req.AvailableCreditLimit,
	})

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, model.AccountResponse{
		Data: model.Account{
			AccountID:            account.AccountID,
			DocumentType:         account.DocumentType,
			DocumentNumber:       account.DocumentNumber,
			AvailableCreditLimit: account.AvailableCreditLimit,
			CreatedAt:            account.CreatedAt,
			UpdatedAt:            account.UpdatedAt,
		},
	})
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
	accountIDStr := c.Param("accountID")

	accountID, err := strconv.Atoi(accountIDStr)

	if err != nil {
		c.Error(ErrMalformedAccountID)
		return
	}

	account, err := a.accountUseCase.GetAccount(c.Request.Context(), accountID)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.AccountResponse{
		Data: model.Account{
			AccountID:            account.AccountID,
			DocumentType:         account.DocumentType,
			DocumentNumber:       account.DocumentNumber,
			AvailableCreditLimit: account.AvailableCreditLimit,
			CreatedAt:            account.CreatedAt,
			UpdatedAt:            account.UpdatedAt,
		},
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/transactions-service/api/model"
	"github.com/italorfeitosa/transactions-service/internal/app"
)

type Transaction struct {
	transactionUseCase app.TransactionUseCase
}

func NewTransaction(transactionUseCase app.TransactionUseCase) *Transaction {
	return &Transaction{transactionUseCase}
}

func (t *Transaction) Register(r gin.IRouter) {
	r.POST("/transactions", t.SaveTransaction)
}

// swagger:route POST /api/v1/transactions transactions saveTransaction
//
// Save Transaction
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Parameters:
//	  + name: Save Transaction Body
//	    in: body
//	    required: true
//	    type: SaveTransactionRequest
//
//	Responses:
//	  201: TransactionResponse
//	  400: SaveTransactionValidationError
//	  422: SaveTransactionBusinessRuleError
//	  500: InternalServerError
func (t *Transaction) SaveTransaction(c *gin.Context) {
	var req model.SaveTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(ErrMalformedSaveTransactionRequest.WithError(err))
		return
	}

	transaction, err := t.transactionUseCase.SaveTransaction(c.Request.Context(), app.SaveTransactionInput{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
	})

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, model.TransactionResponse{
		Data: model.Transaction{
			TransactionID:   transaction.TransactionID,
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount,
			CreatedAt:       transaction.CreatedAt,
		},
	})
}

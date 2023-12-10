package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/transactions-service/api/model"
	"github.com/italorfeitosa/transactions-service/pkg/errs"
)

type Transaction struct{}

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
	c.JSON(http.StatusNotImplemented, model.ErrorResponse{Error: errs.Error{Message: "TODO: implement save transaction"}})
}

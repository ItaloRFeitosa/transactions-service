package app_test

import (
	"fmt"
	"testing"

	"github.com/italorfeitosa/transactions-service/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestValidateOperationType(t *testing.T) {
	type testCase struct {
		operationTypeID, amount int
		wantError               bool
	}

	cases := []testCase{
		{
			operationTypeID: app.PaymentOperation,
			amount:          1000,
		},
		{
			operationTypeID: app.PaymentOperation,
			amount:          -1000,
			wantError:       true,
		},
		{
			operationTypeID: app.PurchaseInInstallmentsOperation,
			amount:          -1000,
		},
		{
			operationTypeID: app.PurchaseInInstallmentsOperation,
			amount:          1000,
			wantError:       true,
		},
		{
			operationTypeID: app.PurchaseOperation,
			amount:          -1000,
		},
		{
			operationTypeID: app.WithdrawOperation,
			amount:          -1000,
		},
		{
			operationTypeID: app.PurchaseOperation,
			amount:          1000,
			wantError:       true,
		},
		{
			operationTypeID: app.WithdrawOperation,
			amount:          1000,
			wantError:       true,
		},
		{
			operationTypeID: 5,
			amount:          1000,
			wantError:       true,
		},
		{
			operationTypeID: -1,
			amount:          1000,
			wantError:       true,
		},
	}

	for _, c := range cases {
		t.Run("validate operation type tests", func(t *testing.T) {
			err := app.ValidateOperationType(c.operationTypeID, c.amount)
			if c.wantError {
				assert.ErrorContains(t, err, fmt.Sprint(c.operationTypeID))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

package app_test

import (
	"testing"

	"github.com/italorfeitosa/transactions-service/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestValidateDocument(t *testing.T) {

	type testCase struct {
		documentType   string
		documentNumber string
		wantError      bool
	}

	cases := []testCase{
		{
			documentType:   app.CPF,
			documentNumber: "55621619072",
		},
		{
			documentType:   app.CPF,
			documentNumber: "556.216.190-72",
			wantError:      true,
		},
		{
			documentType:   app.CPF,
			documentNumber: "43452839060",
		},
		{
			documentType:   app.CPF,
			documentNumber: "20909605076",
		},
		{
			documentType:   app.CPF,
			documentNumber: "20909605076",
		},
		{
			documentType:   app.CPF,
			documentNumber: "11111111111",
			wantError:      true,
		},
		{
			documentType:   app.CPF,
			documentNumber: "99999999999",
			wantError:      true,
		},
		{
			documentType:   app.CNPJ,
			documentNumber: "91490873000185",
		},
		{
			documentType:   app.CNPJ,
			documentNumber: "19991837000178",
		},
		{
			documentType:   app.CNPJ,
			documentNumber: "48430952000171",
		},
		{
			documentType:   app.CNPJ,
			documentNumber: "48430952000171",
		},
		{
			documentType:   app.CNPJ,
			documentNumber: "48.430.952/0001-71",
			wantError:      true,
		},
		{
			documentType:   app.CNPJ,
			documentNumber: "11111111111111",
			wantError:      true,
		},
		{
			documentType:   app.NINO,
			documentNumber: "CS020744C",
		},
		{
			documentType:   app.NINO,
			documentNumber: "WP524362C",
		},
		{
			documentType:   app.NINO,
			documentNumber: "wp524362c",
			wantError:      true,
		},
		{
			documentType:   app.NINO,
			documentNumber: "WP 52 43 62 C",
			wantError:      true,
		},
		{
			documentType:   app.NINO,
			documentNumber: "1S020744C",
			wantError:      true,
		},
		{
			documentType:   app.NINO,
			documentNumber: "CS02A744C",
			wantError:      true,
		},
	}

	for _, c := range cases {
		t.Run("validate document tests", func(t *testing.T) {
			err := app.ValidateDocument(c.documentType, c.documentNumber)
			if c.wantError {
				assert.ErrorContains(t, err, c.documentType)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

package app

const (
	PurchaseOperation = iota + 1
	PurchaseInInstallmentsOperation
	WithdrawOperation
	PaymentOperation
)

func ValidateOperationType(operationTypeID, amount int) error {
	if operationTypeID < PurchaseOperation || operationTypeID > PaymentOperation {
		return ErrInvalidOperationType.WithArgs(operationTypeID)
	}

	if operationTypeID == PaymentOperation && amount < 0 {
		return ErrNegativeAmount.WithArgs(operationTypeID)
	}

	if operationTypeID < PaymentOperation && amount > 0 {
		return ErrPositiveAmount.WithArgs(operationTypeID)
	}

	return nil
}

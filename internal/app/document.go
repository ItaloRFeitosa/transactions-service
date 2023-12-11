package app

import (
	"regexp"
	"strconv"
)

const (
	CPF  = "CPF"
	CNPJ = "CNPJ"
	NINO = "NINO"
)

var justNumbersRegex = regexp.MustCompile(`\d`)

func ValidateDocument(documentType string, documentNumber string) error {
	switch documentType {
	case CPF:
		return ValidateCPF(documentNumber)
	case CNPJ:
		return ValidateCNPJ(documentNumber)
	case NINO:
		return ValidateNINO(documentNumber)
	default:
		return ErrDocumentNotAllowed.WithArgs(documentType)
	}
}

func ValidateCPF(documentNumber string) error {
	if !justNumbersRegex.MatchString(documentNumber) {
		return ErrInvalidDocumentNumber.WithArgs(CPF)
	}

	if len(documentNumber) != 11 {
		return ErrInvalidDocumentNumber.WithArgs(CPF)
	}

	var cpfDigits [11]int
	for i, char := range documentNumber {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return ErrInvalidDocumentNumber.WithArgs(CPF)
		}
		cpfDigits[i] = digit
	}

	// Check if all CPF digits are the same (a common invalid CPF condition)
	allSame := true
	for i := 1; i < 11; i++ {
		if cpfDigits[i] != cpfDigits[i-1] {
			allSame = false
			break
		}
	}
	if allSame {
		return ErrInvalidDocumentNumber.WithArgs(CPF)
	}

	// Validate the CPF using the algorithm
	sum := 0
	for i := 0; i < 9; i++ {
		sum += cpfDigits[i] * (10 - i)
	}
	remainder := sum % 11

	// Calculate the first verification digit
	expectedDigit1 := 11 - remainder
	if expectedDigit1 == 10 || expectedDigit1 == 11 {
		expectedDigit1 = 0
	}
	if cpfDigits[9] != expectedDigit1 {
		return ErrInvalidDocumentNumber.WithArgs(CPF)
	}

	// Calculate the second verification digit
	sum = 0
	for i := 0; i < 10; i++ {
		sum += cpfDigits[i] * (11 - i)
	}
	remainder = sum % 11

	expectedDigit2 := 11 - remainder
	if expectedDigit2 == 10 || expectedDigit2 == 11 {
		expectedDigit2 = 0
	}
	if cpfDigits[10] != expectedDigit2 {
		return ErrInvalidDocumentNumber.WithArgs(CPF)
	}

	return nil
}

func ValidateCNPJ(documentNumber string) error {
	if !justNumbersRegex.MatchString(documentNumber) {
		return ErrInvalidDocumentNumber.WithArgs(CNPJ)
	}

	if len(documentNumber) != 14 {
		return ErrInvalidDocumentNumber.WithArgs(CNPJ)
	}

	var cnpjDigits [14]int
	for i, char := range documentNumber {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return ErrInvalidDocumentNumber.WithArgs(CNPJ)
		}
		cnpjDigits[i] = digit
	}

	sum := 0
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 12; i++ {
		sum += cnpjDigits[i] * weights[i]
	}
	remainder := sum % 11

	expectedDigit1 := 11 - remainder
	if expectedDigit1 >= 10 {
		expectedDigit1 = 0
	}
	if cnpjDigits[12] != expectedDigit1 {
		return ErrInvalidDocumentNumber.WithArgs(CNPJ)
	}

	sum = 0
	weights = append([]int{6}, weights...) // Include the first position for the second digit calculation
	for i := 0; i < 13; i++ {
		sum += cnpjDigits[i] * weights[i]
	}
	remainder = sum % 11

	expectedDigit2 := 11 - remainder
	if expectedDigit2 >= 10 {
		expectedDigit2 = 0
	}
	if cnpjDigits[13] != expectedDigit2 {
		return ErrInvalidDocumentNumber.WithArgs(CNPJ)
	}

	return nil
}

var ninoRegex = regexp.MustCompile(`^[A-Z]{2}\d{6}[A-Z]{1}$`)

func ValidateNINO(documentNumber string) error {
	if !ninoRegex.MatchString(documentNumber) {
		return ErrInvalidDocumentNumber.WithArgs(NINO)
	}
	return nil
}

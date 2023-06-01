package service

import (
	models "golang-credit-simulator/models/credit_simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRateCredit(t *testing.T) {
	t.Log("check result rate, case rate car not equals respon rate")

	cs := models.CreditSimulator{
		VehicleType:      "CAR",
		VehicleCondition: "NEW",
		Year:             2022,
		DownPayment:      80_000_000,
		Total:            200_000_000,
		Tenor:            4,
	}

	handler := NewCreditSimulator()
	rate, err := handler.CheckRateAndValidationCredit(cs)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.NotEqual(t, rate, 9)
}

func TestValidationRequiredCredit(t *testing.T) {
	t.Log("validation request required, case request year, and tenor is empty")

	cs := models.CreditSimulator{
		VehicleType:      "CAR",
		VehicleCondition: "NEW",
		DownPayment:      80_000_000,
		Total:            200_000_000,
	}

	handler := NewCreditSimulator()
	_, err := handler.CheckRateAndValidationCredit(cs)
	if err != nil {
		assert.Equal(t, err.Error(), "request year, tenor is required")
	}
	assert.Zero(t, cs.Year)
	assert.Zero(t, cs.Tenor)
}

func TestValidationDownPaymentCredit(t *testing.T) {
	t.Log("validation down payment, case downpayment is less then the minimum (minimum is 70_000_000)")

	cs := models.CreditSimulator{
		VehicleType:      "CAR",
		VehicleCondition: "NEW",
		Year:             2022,
		DownPayment:      20_000_000,
		Total:            200_000_000,
		Tenor:            4,
	}

	handler := NewCreditSimulator()
	_, err := handler.CheckRateAndValidationCredit(cs)
	if err != nil {
		assert.Equal(t, err.Error(), "down payment is less than the minimum")
	}
}

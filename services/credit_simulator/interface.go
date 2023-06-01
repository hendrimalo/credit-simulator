package service

import (
	models "golang-credit-simulator/models/credit_simulator"
)

type CreditSimulatorUsecase interface {
	CheckRateAndValidationCredit(req models.CreditSimulator) (rate float32, err error)
	GenerateCreditSimulation(rate float32, req models.CreditSimulator) (res []models.CreditSimulatorRespon)
}

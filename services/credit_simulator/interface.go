package service

import (
	models "golang-credit-simulator/models/credit_simulator"

	"github.com/gin-gonic/gin"
)

type CreditSimulatorUsecase interface {
	ReadFile(c *gin.Context) (res models.CreditSimulator, err error)
	CheckRateAndValidationCredit(req models.CreditSimulator) (rate float32, err error)
	GenerateCreditSimulation(rate float32, req models.CreditSimulator) (res []models.CreditSimulatorRespon)
}

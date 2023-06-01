package credit_simulator

import (
	models "golang-credit-simulator/models/credit_simulator"
	service "golang-credit-simulator/services/credit_simulator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreditSimulatorHandler struct {
	uCS service.CreditSimulatorUsecase
}

func NewCreditSimulatorHandler(c *gin.Engine, ucs service.CreditSimulatorUsecase) {
	handler := &CreditSimulatorHandler{
		uCS: ucs,
	}
	csr := c.Group("v1/credit-simulator")
	csr.POST("/", handler.CreateCreditSimulator)

}

func (ch *CreditSimulatorHandler) CreateCreditSimulator(c *gin.Context) {
	var cs models.CreditSimulator
	if err := c.ShouldBindJSON(&cs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	rate, err := ch.uCS.CheckRateAndValidationCredit(cs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	res := ch.uCS.GenerateCreditSimulation(rate, cs)
	c.JSON(http.StatusOK, res)
}
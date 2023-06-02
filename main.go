package main

import (
	"golang-credit-simulator/controllers/credit_simulator"
	service "golang-credit-simulator/services/credit_simulator"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//credit-simulator routes
	ucs := service.NewCreditSimulator()
	credit_simulator.NewCreditSimulatorHandler(r, ucs)

	r.Run("0.0.0.0:3000")
}

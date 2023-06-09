package service

import (
	"errors"
	"fmt"
	models "golang-credit-simulator/models/credit_simulator"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CreditSimulator struct{}

func NewCreditSimulator() CreditSimulatorUsecase {
	return &CreditSimulator{}
}

func (cs CreditSimulator) ReadFile(c *gin.Context) (res models.CreditSimulator, err error) {
	//read file txt
	path := fmt.Sprintf("./files/input_credit/%s.txt", c.Param("filename"))
	content, err := os.ReadFile(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//validate format txt
	data := strings.Split(string(content), ",")
	if len(data) != 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid format txt",
		})
		return
	}

	//desctruction txt to variable
	vtype, condition, year, downPayment, total, tenor := data[0], data[1], data[2], data[3], data[4], data[5]
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return
	}

	//coverting format string to specific request format
	downPaymentF32, _ := strconv.ParseFloat(downPayment, 32)
	if err != nil {
		return
	}

	totalF32, _ := strconv.ParseFloat(total, 32)
	if err != nil {
		return
	}

	tenorF32, err := strconv.ParseFloat(tenor, 32)
	if err != nil {
		return
	}

	res = models.CreditSimulator{
		VehicleType:      vtype,
		VehicleCondition: condition,
		Year:             yearInt,
		DownPayment:      float32(downPaymentF32),
		Total:            float32(totalF32),
		Tenor:            int(tenorF32),
	}
	return
}

func (cs CreditSimulator) CheckRateAndValidationCredit(req models.CreditSimulator) (rate float32, err error) {
	var (
		temp           []string
		percentPayment int
	)

	//validate required request
	if len(req.VehicleType) == 0 {
		temp = append(temp, "vehicle_type")
	}
	if len(req.VehicleCondition) == 0 {
		temp = append(temp, "vehicle_condition")
	}
	if req.Year == 0 {
		temp = append(temp, "year")
	}
	if req.DownPayment == 0 {
		temp = append(temp, "down_payment")
	}
	if req.Total == 0 {
		temp = append(temp, "total")
	}
	if req.Tenor == 0 {
		temp = append(temp, "tenor")
	}
	if len(temp) != 0 {
		required := strings.Join(temp, ", ")
		err = fmt.Errorf("request %s is required", required)
		return
	}

	//check vehicle type
	switch strings.ToUpper(req.VehicleType) {
	case "CAR":
		rate = 8.00
	case "MOTORCYCLE":
		rate = 9.00
	default:
		err = errors.New("request vehicle_type invalid")
		return
	}

	//validate tenor
	if req.Tenor > 6 {
		err = errors.New("tenor cannot be more than 6 years")
		return
	}

	//validate year vehicle conditional
	switch strings.ToUpper(req.VehicleCondition) {
	case "NEW":
		if req.Year < time.Now().Year()-1 {
			err = errors.New("new vehicle must not be more than 2 years old")
			return
		}
		percentPayment = 35
	case "SECOND":
		percentPayment = 35
	default:
		err = errors.New("request vehicle_condition invalid")
		return
	}

	//validate downpayment
	if req.DownPayment < req.Total*(float32(percentPayment)/100) {
		err = errors.New("down payment is less than the minimum")
		return
	}
	return
}

func (cs CreditSimulator) GenerateCreditSimulation(rate float32, req models.CreditSimulator) (res []models.CreditSimulatorRespon) {
	loanPrincipal := (req.Total - req.DownPayment) / float32(req.Tenor)
	for i := 1; i <= req.Tenor; i++ {
		start := i != 1
		if i%2 == 0 && start {
			rate += 0.10
		} else if i%2 == 1 && start {
			rate += 0.50
		}

		installmentPayment := loanPrincipal + (loanPrincipal * (rate / 100))
		loanPrincipal = installmentPayment
		res = append(res, models.CreditSimulatorRespon{
			Tahun:              i,
			Rate:               rate,
			InstallmentPayment: installmentPayment,
		})
	}
	return
}

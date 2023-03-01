package handlers

import (
	"msa-bank-credit-cs/models"
	"msa-bank-credit-cs/pkg/services"
	"net/http"
	"math"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	CreditService    *services.Credit
	RepaymentService *services.EarlyRepayment
}

func NewHandler(creditService *services.Credit, repaymentService *services.EarlyRepayment) *Handler {
	return &Handler{
		CreditService:    creditService,
		RepaymentService: repaymentService,
	}
}

func (h Handler) Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "The msa-bank-credit-cs microservice!")
}

// Create new credit
// (POST /credit)
func (h Handler) PostCredit(ctx echo.Context) error {

	req := &models.PostCreditJSONRequestBody{}
	if err := ctx.Bind(req); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	reqCredit := &models.Credit{
		Amount:   req.Amount,
		ClientId: req.ClientId,
		Months:   req.Months,
		Rate:     req.Rate,
	}

	resCredit, err := h.CreditService.PostCredit(reqCredit)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusCreated, resCredit)
}

// Get credit paymentPlan from the store
// (POST /credit/paymentPlan)
func (h Handler) PostCreditPaymentPlan(ctx echo.Context) error {
	req := &models.PostCreditPaymentPlanJSONRequestBody{}
	if err := ctx.Bind(req); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	reqRate := &models.Rate{
		Amount:   req.Amount,
		Months:   req.Months,
		Rate:     req.Rate,
	}
	
	fullAmount := *reqRate.Amount
	var totalAmount float32
	var allpayment models.AllPayment
	var monthlyRate = (*reqRate.Rate/12)/100
	var annuityRatio = float32(float64(monthlyRate) * (math.Pow(float64(1+monthlyRate), float64(*reqRate.Months)))/((math.Pow(float64(1+monthlyRate), float64(*reqRate.Months)))-1))

	for i := 1; i < *reqRate.Months+1; i++ {
		totalAmount += fullAmount * annuityRatio
	}

	for i := 1; i < *reqRate.Months+1; i++ {
		var monthPayment = fullAmount * annuityRatio
		totalAmount -= monthPayment
		allpayment = append(allpayment, models.PaymentPlan{
		MonthPayment: monthPayment,
		Month : i,
		Amount: totalAmount,
		},)
	}

	return ctx.JSON(http.StatusOK, allpayment)

}

// Create new earlyRepayment
// (POST /credit/earlyRepayment)
func (h Handler) PostCreditEarlyRepayment(ctx echo.Context) error {
	req := &models.PostCreditEarlyRepaymentJSONRequestBody{}
	if err := ctx.Bind(req); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res, err := h.RepaymentService.PostEarlyRepayment(req)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusCreated, res)
}


// Creating a request for full repayment
// (POST /credit/fullRepayment)
func (h Handler) PostCreditFullRepayment(ctx echo.Context) error {
	req := &models.PostCreditFullRepaymentJSONRequestBody{}
	if err := ctx.Bind(req); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	reqRepayment := &models.EarlyRepayment{
		Amount:   0.0,
		Id: req.Id,
	}

	res, err := h.RepaymentService.PostEarlyRepayment(reqRepayment)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusCreated, res)
}

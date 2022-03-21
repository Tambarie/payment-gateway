package handler

import (
	"github.com/Tambarie/payment-gateway/domain/helpers"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"github.com/Tambarie/payment-gateway/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Authorize handler authorizes the merchant to fund his wallet and gives him a unique ID to capture and refund payments
func (h *Handler) Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)
		userReference := context.Param("user-reference")

		var merchant = &domain.Card{}

		merchant.ID = uuid.New().String()
		merchant.UserReference = userReference

		// Binding the merchant json
		if err := helpers.Decode(context, &merchant); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// check if user Reference exists in the database
		_, err := h.PaymentGatewayService.CheckIfUserExists(userReference)
		if err != nil {
			response.JSON(context, http.StatusBadRequest, nil, nil, "user does not exists")
			return
		}

		// checking the validation of the credit card using Luhn algorithm
		if false == helpers.Valid(merchant.CardNumber) {
			response.JSON(context, http.StatusBadRequest, nil, nil, "your credit card is not valid")
			return
		}

		checker := 119
		// Checking for the authorisation of the credit card
		res := helpers.AuthorisationFailure(merchant.CardNumber, checker)
		if res == true {
			response.JSON(context, http.StatusBadRequest, nil, nil, "Authorisation failure")
			return
		}

		// method that handles the authorization of the credit card
		card, err := h.PaymentGatewayService.Authorise(merchant)
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		//JSON response to the client
		response.JSON(context, http.StatusCreated, gin.H{
			"Unique ID": card.ID,
			"Amount":    card.Amount,
			"Currency":  card.Currency,
		},
			nil,
			"Success")
	}
}

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

func (h *Handler) Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		helpers.LogRequest(context)
		userReference := context.Param("user-reference")
		var merchant = &domain.Card{}
		merchant.ID = uuid.New().String()
		merchant.UserReference = userReference

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

		// check if credit card is valid
		if false == helpers.Valid(merchant.CardNumber) {
			response.JSON(context, http.StatusBadRequest, nil, nil, "your credit card is not valid")
			return
		}

		checker := 119
		res := helpers.AuthorisationFailure(merchant.CardNumber, checker)
		if res == true {
			response.JSON(context, http.StatusBadRequest, nil, nil, "Authorisation failure")
			return
		}

		card, err := h.PaymentGatewayService.Authorise(merchant)
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		response.JSON(context, http.StatusCreated, gin.H{
			"Unique ID": card.ID,
			"Amount":    card.Amount,
			"Currency":  card.Currency,
		},
			nil,
			"Success")
	}
}

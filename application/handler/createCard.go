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
		var merchant = &domain.Card{}
		merchant.ID = uuid.New().String()

		if err := helpers.Decode(context, &merchant); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		checker := 119
		log.Println(merchant.CardNumber)
		res := helpers.AuthorisationFailure(merchant.CardNumber, checker)
		log.Println(res)
		if res == true {
			response.JSON(context, http.StatusForbidden, nil, nil, "You can't use this card")
			return
		}

		card, err := h.PaymentGatewayService.CreateMerchant(merchant)
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

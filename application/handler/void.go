package handler

import (
	"github.com/Tambarie/payment-gateway/domain/helpers"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"github.com/Tambarie/payment-gateway/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) Void() gin.HandlerFunc {
	return func(context *gin.Context) {
		card := &domain.Card{}
		void := &domain.Void{}

		if err := helpers.Decode(context, &void); err != nil {
			log.Fatal(err)
			return
		}

		cardDB, err := h.PaymentGatewayService.GetCardByID(void.AuthorizationID)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, nil, "No documents in results, please enter a valid authorisation token")
			return
		}

		card.VoidCard()

		_, err = h.PaymentGatewayService.VoidCard(void.AuthorizationID, card.Void)
		if err != nil {
			log.Fatal(err)
			return
		}
		response.JSON(context, 201, gin.H{
			"message":         "your card has been blocked",
			"account balance": cardDB["amount"],
		}, nil, "")
	}
}

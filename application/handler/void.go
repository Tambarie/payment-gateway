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
		// Logging the context
		helpers.LogRequest(context)

		card := &domain.Card{}
		void := &domain.Void{}

		// Binding the void json
		if err := helpers.Decode(context, &void); err != nil {
			log.Fatal(err)
			return
		}

		// method that handles the getting of card using the ID
		cardDB, err := h.PaymentGatewayService.GetCardByID(void.AuthorizationID)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, nil, "No documents in results, please enter a valid authorisation token")
			return
		}

		card.VoidCard()

		// method that voids the merchant's card
		_, err = h.PaymentGatewayService.VoidCard(void.AuthorizationID, card.Void)
		if err != nil {
			log.Fatal(err)
			return
		}
		//JSON response to the client
		response.JSON(context, 201, gin.H{
			"message":         "your card has been blocked",
			"account balance": cardDB["amount"],
		}, nil, "")
	}
}

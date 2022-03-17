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

		if err := helpers.Decode(context, &card); err != nil {
			log.Fatal(err)
			return
		}

		cardDB, err := h.PaymentGatewayService.GetCardByID(card.ID)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, nil, "No documents in results, please enter a valid authorisation token")
			return
		}

		if card.ID != cardDB["id"] {
			if err != nil {
				response.JSON(context, http.StatusForbidden, nil, nil, "Please enter a valid unique ID")
				return
			}

			card.VoidCard()

			_, err := h.PaymentGatewayService.UpdateAccount(card.ID)

		}
	}
}

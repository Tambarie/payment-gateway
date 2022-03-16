package handler

import (
	"github.com/Tambarie/payment-gateway/domain/helpers"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"github.com/Tambarie/payment-gateway/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) Capture() gin.HandlerFunc {
	return func(context *gin.Context) {
		var captured = &domain.Capture{}

		if err := helpers.Decode(context, &captured); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		card, err := h.PaymentGatewayService.GetID(captured.AuthorizationID)
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		checker := 259
		var cardNumber = card["card_number"].(int64)
		res := helpers.AuthorisationFailure(cardNumber, checker)
		if res == true {
			response.JSON(context, http.StatusForbidden, nil, nil, "You can't use this card")
			return
		}

		if captured.AuthorizationID != card["id"] {
			response.JSON(context, http.StatusBadRequest, nil, nil, "Please enter a valid Unique ID")
			return
		}

		var balance = card["amount"].(float64)
		var currentBalance float64

		if balance < captured.Amount {
			response.JSON(context, http.StatusBadRequest, nil, nil, "you have insufficient balance")
			return
		} else {
			currentBalance = balance - captured.Amount
			_, err := h.PaymentGatewayService.UpdateAccount(currentBalance, captured.AuthorizationID)
			if err != nil {
				log.Fatalf("Error %v", err)
				return
			}

			_, err = h.PaymentGatewayService.SaveCapturedTransaction(captured)
			if err != nil {
				log.Fatalf("Error %v", err)
				return
			}

			currentDB, err := h.PaymentGatewayService.GetID(captured.AuthorizationID)
			if err != nil {
				log.Fatalf("Error %v", err)
				return
			}

			response.JSON(context, 201, gin.H{
				"Amount Debited":  captured.Amount,
				"Account Balance": currentDB["amount"],
			}, nil, "successfully captured")
		}
	}
}

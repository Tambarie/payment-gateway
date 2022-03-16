package handler

import (
	"github.com/Tambarie/payment-gateway/domain/helpers"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"github.com/Tambarie/payment-gateway/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) Refund() gin.HandlerFunc {
	return func(context *gin.Context) {
		refund := domain.Refund{}
		merchant := domain.Card{}

		if err := helpers.Decode(context, &refund); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		capturedTransaction, err := h.PaymentGatewayService.GetCapturedTransaction(refund.AuthorizationID)
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		merchantCard, err := h.PaymentGatewayService.GetID(refund.AuthorizationID)
		log.Println(merchantCard["id"])
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		if capturedTransaction["transaction_id"] != refund.TransactionID {
			response.JSON(context, http.StatusBadRequest, nil, nil, "refund ID not valid")
			return
		}

		//if merchantCard == nil {
		//	response.JSON(context, http.StatusBadRequest, nil, nil, "refund ID not valid")
		//	return
		//}

		checker := 3238
		var cardNumber = merchantCard["card_number"].(int64)
		res := helpers.AuthorisationFailure(cardNumber, checker)
		if res == true {
			response.JSON(context, http.StatusForbidden, nil, nil, "You can't use this card")
			return
		}

		capturedTransaction, err = h.PaymentGatewayService.GetCapturedTransaction(refund.AuthorizationID)
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		log.Println(capturedTransaction["amount"])

		capturedAmount := capturedTransaction["amount"].(float64)
		authorisedAmount := merchantCard["amount"].(float64)

		if capturedAmount < refund.Amount && authorisedAmount < refund.Amount {
			response.JSON(context, http.StatusBadRequest, nil, nil, "you are not entitled to be refunded")
			return
		}

		merchant.RefundMerchant(refund.Amount)

		_, err = h.PaymentGatewayService.UpdateAccount(merchant.Amount, refund.AuthorizationID)
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}

	}
}

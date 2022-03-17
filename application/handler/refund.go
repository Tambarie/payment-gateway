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
		refund := &domain.Refund{}
		refundTracker := &domain.RefundTracker{}

		if err := helpers.Decode(context, &refund); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		capturedTransaction, err := h.PaymentGatewayService.GetCapturedTransactionByTransactionID(refund.TransactionID)
		if err != nil {
			response.JSON(context, http.StatusBadRequest, nil, nil, "transaction ID is not valid")
			return
		}

		merchantCard, err := h.PaymentGatewayService.GetCardByID(refund.AuthorizationID)
		if err != nil {
			response.JSON(context, http.StatusBadRequest, nil, nil, "please enter a valid authorizationID")
			return
		}

		if capturedTransaction.TransactionID != refund.TransactionID {
			response.JSON(context, http.StatusBadRequest, nil, nil, "refund ID not valid")
			return
		}
		if refund.Amount > capturedTransaction.Amount {
			response.JSON(context, http.StatusBadRequest, nil, nil, "sorry, you can't be refunded")
			return
		}

		checker := 3238
		var cardNumber = merchantCard["card_number"].(int64)
		res := helpers.AuthorisationFailure(cardNumber, checker)
		if res == true {
			response.JSON(context, http.StatusForbidden, nil, nil, "You can't use this card")
			return
		}

		capturedAmount := capturedTransaction.Amount
		authorisedAmount := merchantCard["amount"].(float64)

		if capturedAmount < refund.Amount && authorisedAmount < refund.Amount {
			response.JSON(context, http.StatusBadRequest, nil, nil, "you are not entitled to be refunded")
			return
		}

		var refundBalance float64
		var balance = merchantCard["amount"].(float64)
		//var count = merchantCard["count"].(int32)

		refTracker, err := h.PaymentGatewayService.GetRefundTrackerByTransactionID(refund.TransactionID)
		if err == nil {
			var count = refTracker["count"].(int32)
			log.Println("am here")
			if count > 0 {
				response.JSON(context, http.StatusBadRequest, nil, nil, "Sorry you have already been refunded your money")
				return
			}
		}

		refundBalance = refund.Amount + balance
		_, err = h.PaymentGatewayService.RefundUpdateAccount(refundBalance, refund.AuthorizationID, 1)
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}
		refundTracker.Count = 1
		refundTracker.TransactionID = capturedTransaction.TransactionID

		_, err = h.PaymentGatewayService.SaveRefundTracker(refundTracker)
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}
		final, err := h.PaymentGatewayService.GetCardByID(refund.AuthorizationID)
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}
		response.JSON(context, 201, gin.H{
			"Amount Refunded": refund.Amount,
			"Account Balance": final["amount"],
		}, nil, "successfully refunded")
	}
}

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

		helpers.LogRequest(context)

		refund := &domain.Refund{}
		refundTracker := &domain.RefundTracker{}

		// Binding the refund json
		if err := helpers.Decode(context, &refund); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// getting captured transaction by transaction ID
		capturedTransaction, err := h.PaymentGatewayService.GetCapturedTransactionByTransactionID(refund.TransactionID)
		if err != nil {
			response.JSON(context, http.StatusBadRequest, nil, nil, "transaction ID is not valid")
			return
		}

		// method that handles the getting of card using the ID
		merchantCard, err := h.PaymentGatewayService.GetCardByID(refund.AuthorizationID)
		if err != nil {
			response.JSON(context, http.StatusBadRequest, nil, nil, "please enter a valid authorizationID")
			return
		}

		// Checking if card is void
		if merchantCard["void"] == true {
			response.JSON(context, http.StatusBadRequest, nil, nil, "Sorry, your card is not valid")
			return
		}

		// checking if transactionID is valid
		if capturedTransaction.TransactionID != refund.TransactionID {
			response.JSON(context, http.StatusBadRequest, nil, nil, "refund ID not valid")
			return
		}
		//checking if amount inputted to refund is  more than the actual amount
		if refund.Amount > capturedTransaction.Amount {
			response.JSON(context, http.StatusBadRequest, nil, nil, "sorry, you can't be refunded")
			return
		}

		// Checking if card  ends with 3238
		checker := 3238
		var cardNumber = merchantCard["card_number"].(int64)
		res := helpers.AuthorisationFailure(cardNumber, checker)
		if res == true {
			response.JSON(context, http.StatusBadRequest, nil, nil, "Refund failure")
			return
		}

		capturedAmount := capturedTransaction.Amount
		authorisedAmount := merchantCard["amount"].(float64)

		// checking if the refund amount is more than the  captured amount
		if capturedAmount < refund.Amount && authorisedAmount < refund.Amount {
			response.JSON(context, http.StatusBadRequest, nil, nil, "you are not entitled to be refunded")
			return
		}

		var refundBalance float64
		var balance = merchantCard["amount"].(float64)

		// method that gets the refund tracker by transaction ID
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
		// method that  refunds and update's the merchant account
		_, err = h.PaymentGatewayService.RefundUpdateAccount(refundBalance, refund.AuthorizationID, 1)
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}
		// this tracks if the merchant has been refunded or not
		refundTracker.Count = 1
		refundTracker.TransactionID = capturedTransaction.TransactionID

		// saves the refunded amount
		_, err = h.PaymentGatewayService.SaveRefundTracker(refundTracker)
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// method that handles the getting of card using the ID
		final, err := h.PaymentGatewayService.GetCardByID(refund.AuthorizationID)
		if err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		//JSON response to the client
		response.JSON(context, 201, gin.H{
			"Amount Refunded": refund.Amount,
			"Account Balance": final["amount"],
		}, nil, "successfully refunded")
	}
}

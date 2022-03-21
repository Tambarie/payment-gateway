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

// Capture captures the expenses of the merchant
func (h *Handler) Capture() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)

		var captured = &domain.Transaction{}
		var refund = &domain.RefundTracker{}

		//generating unique transaction ID
		captured.TransactionID = uuid.New().String()

		// Binding the transaction json
		if err := helpers.Decode(context, &captured); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// method that handles the getting of card using the ID
		card, err := h.PaymentGatewayService.GetCardByID(captured.AuthorizationID)
		if err != nil {
			response.JSON(context, http.StatusForbidden, nil, nil, "No documents in results, please enter a valid authorisation token")
			return
		}

		// Checking if card is void
		if card["void"] == true {
			response.JSON(context, http.StatusBadRequest, nil, nil, "Sorry, your card is not valid")
			return
		}

		checker := 259
		var cardNumber = card["card_number"].(int64)
		res := helpers.AuthorisationFailure(cardNumber, checker)
		if res == true {
			response.JSON(context, http.StatusBadRequest, nil, nil, "Capture failure")
			return
		}

		// checking if the authorisation id is not equal to the unique ID issued
		if captured.AuthorizationID != card["id"] {
			response.JSON(context, http.StatusBadRequest, nil, nil, "Please enter a valid Unique ID")
			return
		}

		refund.TransactionID = captured.TransactionID
		// method that gets the refund tracker by transaction ID
		refTracker, err := h.PaymentGatewayService.GetRefundTrackerByTransactionID(refund.TransactionID)
		if err == nil {
			var count = refTracker["count"].(int)
			log.Println("am here")
			if count > 0 {
				response.JSON(context, http.StatusBadRequest, nil, nil, "Sorry you have already been refunded your money")
				return
			}
		}

		var balance = card["amount"].(float64)
		var currentBalance float64

		// checks if the balance is less than the amount to be captured
		if balance < captured.Amount {
			response.JSON(context, http.StatusBadRequest, nil, nil, "you have insufficient balance")
			return
		} else {
			currentBalance = balance - captured.Amount
			// method that handles the updating of the account
			_, err := h.PaymentGatewayService.UpdateAccount(currentBalance, captured.AuthorizationID)
			if err != nil {
				log.Fatalf("Error %v", err)
				return
			}

			// method that saves the captured transaction
			_, err = h.PaymentGatewayService.SaveCapturedTransaction(captured)
			if err != nil {
				log.Fatalf("Error %v", err)
				return
			}

			// method that handles the getting of card using the ID
			currentDB, err := h.PaymentGatewayService.GetCardByID(captured.AuthorizationID)
			if err != nil {
				log.Fatalf("Error %v", err)
				return
			}
			//JSON response to the client
			response.JSON(context, 201, gin.H{
				"Amount Debited":  captured.Amount,
				"Account Balance": currentDB["amount"],
				"TransactionID":   captured.TransactionID,
			}, nil, "successfully captured")
		}
	}
}

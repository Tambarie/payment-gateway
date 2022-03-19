package test

import (
	"bytes"
	"encoding/json"
	"github.com/Tambarie/payment-gateway/application/handler"
	"github.com/Tambarie/payment-gateway/application/server"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"github.com/Tambarie/payment-gateway/domain/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRefund(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockPaymentService := service.NewMockPaymentGatewayService(controller)
	router := gin.Default()
	h := &handler.Handler{PaymentGatewayService: mockPaymentService}
	server.DefineRouter(router, h)

	refund := &domain.Refund{
		TransactionID:   "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		AuthorizationID: "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		Amount:          5,
	}

	marshalledRefund, err := json.Marshal(refund)
	if err != nil {
		log.Fatalf("An err occcured %v", err)
	}

	transaction := &domain.Transaction{
		TransactionID:   "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		AuthorizationID: "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		Amount:          5,
	}

	var card = bson.M{"UserReference": "user",
		"id":               "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		"card_number":      int64(374245455400126),
		"expiration_year":  0,
		"expiration_month": "2023",
		"cvv":              456,
		"amount":           0.00,
		"currency":         "NGN",
		"count":            0,
		"void":             false,
	}

	var refundDB = bson.M{
		"transaction_id": "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		"count":          int32(0),
	}

	var refundBalance float64 = 5

	mockPaymentService.EXPECT().GetCapturedTransactionByTransactionID(refund.TransactionID).Return(transaction, nil)
	mockPaymentService.EXPECT().GetCardByID(gomock.Any()).Return(card, nil).Times(2)
	mockPaymentService.EXPECT().GetRefundTrackerByTransactionID(gomock.Any()).Return(refundDB, nil)
	mockPaymentService.EXPECT().RefundUpdateAccount(refundBalance, refund.AuthorizationID, 1).Return(nil, nil)
	mockPaymentService.EXPECT().SaveRefundTracker(gomock.Any()).Return(nil, nil)

	t.Run("testing for if merchant has been refunded", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPut, "/api/v1/refund", bytes.NewBuffer(marshalledRefund))
		if err != nil {
			log.Fatalf("An err occcured %v", err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != 201 {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}
		var responseBodyTwo = `"Account Balance":0,"Amount Refunded":5},"errors":null,"message":"successfully refunded"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}

	})
}

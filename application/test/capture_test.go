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
	"go/types"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCapture(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockPaymentService := service.NewMockPaymentGatewayService(controller)
	router := gin.Default()
	h := &handler.Handler{PaymentGatewayService: mockPaymentService}
	server.DefineRouter(router, h)

	captured := &domain.Transaction{
		TransactionID:   "1",
		AuthorizationID: "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		Amount:          5,
	}
	refund := &domain.Refund{
		TransactionID:   "1",
		AuthorizationID: "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		Amount:          5,
	}

	marhalledCapture, err := json.Marshal(captured)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	var card = bson.M{"UserReference": "user",
		"id":               "1aa557e6-709b-4c5f-8910-9619c27c8adb",
		"card_number":      int64(374245455400126),
		"expiration_year":  0,
		"expiration_month": "2023",
		"cvv":              456,
		"amount":           10.00,
		"currency":         "NGN",
		"count":            0,
		"void":             false,
	}

	mockPaymentService.EXPECT().GetCardByID(gomock.Any()).Return(card, nil).Times(2)
	mockPaymentService.EXPECT().GetRefundTrackerByTransactionID(gomock.Any()).Return(card, nil)
	mockPaymentService.EXPECT().UpdateAccount(refund.Amount, captured.AuthorizationID).Return(types.Interface{}, nil)
	mockPaymentService.EXPECT().SaveCapturedTransaction(captured).Return(nil, nil)

	t.Run("testing for capture", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/api/v1/capture", bytes.NewBuffer(marhalledCapture))
		if err != nil {
			log.Fatalf("an error occured :%v", err)
		}

		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusCreated {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}
		var responseBodyTwo = `"Account Balance":10,"Amount Debited":5,"TransactionID":"1"},"errors":null,"message":"successfully captured"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})
}

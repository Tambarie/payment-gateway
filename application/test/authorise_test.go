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

func TestAuthorise(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := gomock.NewController(t)
	mockPaymentService := service.NewMockPaymentGatewayService(controller)
	router := gin.Default()
	h := &handler.Handler{PaymentGatewayService: mockPaymentService}
	server.DefineRouter(router, h)

	merchant := &domain.Card{
		UserReference:   "user",
		ID:              "1",
		CardNumber:      374245455400126,
		ExpirationYear:  0,
		ExpirationMonth: "2023",
		Cvv:             456,
		Amount:          10,
		Currency:        "NGN",
		Count:           0,
		Void:            false,
	}

	marshal, err := json.Marshal(merchant)
	if err != nil {
		log.Fatal(err)
	}

	var verify bson.M
	mockPaymentService.EXPECT().CheckIfUserExists(gomock.Any()).Return(verify, nil)
	mockPaymentService.EXPECT().Authorise(gomock.Any()).Return(merchant, nil)

	t.Run("Testing authorise card", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/api/v1/authorize/data", bytes.NewBuffer(marshal))
		if err != nil {
			log.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusCreated {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}

		var responseBodyTwo = `"Amount":10,"Currency":"NGN","Unique ID":"1"},"errors":null,"message":"Success","status":201,"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})
}

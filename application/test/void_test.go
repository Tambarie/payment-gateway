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

func TestVoid(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockPaymentService := service.NewMockPaymentGatewayService(controller)
	router := gin.Default()
	h := &handler.Handler{PaymentGatewayService: mockPaymentService}
	server.DefineRouter(router, h)

	void := &domain.Void{
		AuthorizationID: "1aa557e6-709b-4c5f-8910-9619c27c8adb",
	}
	//card:= &domain.Card{}

	marshalledVoid, err := json.Marshal(void)
	if err != nil {
		log.Fatalf("could not marshal json %v", err)
	}

	var cardDB = bson.M{"UserReference": "user",
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

	mockPaymentService.EXPECT().GetCardByID(void.AuthorizationID).Return(cardDB, nil)
	mockPaymentService.EXPECT().VoidCard(void.AuthorizationID, true).Return(nil, nil)

	t.Run("test for void", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPut, "/api/v1/void", bytes.NewBuffer(marshalledVoid))
		if err != nil {
			log.Fatalf("an error occured :%v", err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusCreated {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}
		var responseBodyTwo = `"account balance":10,"message":"your card has been blocked"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})
}

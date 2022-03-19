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
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockPaymentService := service.NewMockPaymentGatewayService(controller)
	router := gin.Default()
	h := &handler.Handler{PaymentGatewayService: mockPaymentService}
	server.DefineRouter(router, h)

	user := &domain.User{
		Reference:    "1",
		FirstName:    "Emmanuel",
		LastName:     "Gbaragbo",
		Email:        "gt.tammy@gmail.com",
		Password:     "kjkdd",
		HashPassword: []byte("dsdsdd"),
		TimeCreated:  "",
	}

	marshal, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	mockPaymentService.EXPECT().CheckIfEmailExists(user.Email).Return(nil, errors.New("error"))
	mockPaymentService.EXPECT().CreateUser(gomock.Any()).Return(user, nil)

	t.Run("test for creating user", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/api/v1/create-account", bytes.NewBuffer(marshal))
		if err != nil {
			log.Fatalf("an error occured:%v", err)
		}

		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusCreated {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}

		var responseBodyTwo = `"reference":"1","first_name":"Emmanuel","last_name":"Gbaragbo","email":"gt.tammy@gmail.com"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})
}

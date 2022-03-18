package handler

import (
	"github.com/Tambarie/payment-gateway/domain/helpers"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"github.com/Tambarie/payment-gateway/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) CreateAccount() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := &domain.User{}
		user.Reference = uuid.New().String()
		user.TimeCreated = time.Now().String()
		user.Email = strings.TrimSpace(user.Email)
		user.Password = strings.TrimSpace(user.Password)

		hashedPassword, err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			log.Fatalf("Error, %v", err)
			return
		}
		user.HashPassword = hashedPassword

		if err := helpers.Decode(context, &user); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// check if email exists in the database
		_, err = h.PaymentGatewayService.CheckIfEmailExists(user.Email)
		if err != nil {
			response.JSON(context, http.StatusBadRequest, nil, nil, "your email does not exists")
			return

		}

		//if user.Email != userDB["email"]{
		//	log.Fatalf("Error %v",err)
		//	return
		//}

		userD, err := h.PaymentGatewayService.CreateUser(user)
		if err != nil {
			log.Fatalf("Error, %v", err)
		}
		response.JSON(context, http.StatusCreated, userD, nil, "")
	}
}

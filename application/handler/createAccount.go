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
		// Logging the context
		helpers.LogRequest(context)

		user := &domain.User{}
		user.Reference = uuid.New().String()
		user.TimeCreated = time.Now().String()
		user.Email = strings.TrimSpace(user.Email)
		user.Password = strings.TrimSpace(user.Password)

		// function that handles the generation of hashed merchant's password
		hashedPassword, err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			log.Fatalf("Error, %v", err)
			return
		}
		user.HashPassword = hashedPassword

		// Binding the user json
		if err := helpers.Decode(context, &user); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// check if email exists in the database
		_, err = h.PaymentGatewayService.CheckIfEmailExists(user.Email)
		if err != nil {
			// creates user in the DB
			userD, err := h.PaymentGatewayService.CreateUser(user)
			if err != nil {
				log.Fatalf("Error, %v", err)
				return
			}
			//JSON response to the client
			response.JSON(context, http.StatusCreated, userD, nil, "account created successfully")
			return
		}
		//JSON response to the client
		response.JSON(context, http.StatusCreated, "", nil, "user already exists")
	}
}

package handler

import (
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Capture() gin.HandlerFunc {
	return func(context *gin.Context) {
		var capure = &domain.Capture{}

	}
}

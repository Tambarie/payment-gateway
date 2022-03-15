package handler

import (
	"github.com/Tambarie/payment-gateway/domain/service"
)

type Handler struct {
	PaymentGatewayService service.PaymentGatewayService
}

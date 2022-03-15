package service

import (
	"github.com/Tambarie/payment-gateway/application/paymentGatewayRepository"
)

// WalletService interface
type PaymentGatewayService interface{}

// DefaultWalletService struct
type DefaultWalletService struct {
	repo paymentGatewayRepository.Repository
}

func NewPaymentGatewayService(repo paymentGatewayRepository.Repository) *DefaultWalletService {
	return &DefaultWalletService{
		repo: repo,
	}
}

package service

import (
	"github.com/Tambarie/payment-gateway/application/payment-repository"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
)

// PaymentGatewayService interface
type PaymentGatewayService interface {
	CreateMerchant(card *domain.Card) (*domain.Card, error)
}

// DefaultWalletService struct
type DefaultWalletService struct {
	repo payment_repository.Repository
}

func NewPaymentGatewayService(repo payment_repository.Repository) *DefaultWalletService {
	return &DefaultWalletService{
		repo: repo,
	}
}

func (s *DefaultWalletService) CreateMerchant(card *domain.Card) (*domain.Card, error) {
	return s.repo.CreateMerchant(card)
}

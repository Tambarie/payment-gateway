package payment_repository

import domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"

// Repository interface
type Repository interface {
	CreateMerchant(card *domain.Card) (*domain.Card, error)
}

package service

import (
	"github.com/Tambarie/payment-gateway/application/payment-repository"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PaymentGatewayService interface
type PaymentGatewayService interface {
	CreateMerchant(card *domain.Card) (*domain.Card, error)
	GetID(id string) (bson.M, error)
	UpdateAccount(amount float64, id string) (*mongo.UpdateResult, error)
	SaveCapturedTransaction(capture *domain.Capture) (*domain.Capture, error)
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

func (s *DefaultWalletService) GetID(id string) (bson.M, error) {
	return s.repo.GetID(id)
}

func (s *DefaultWalletService) UpdateAccount(amount float64, id string) (*mongo.UpdateResult, error) {
	return s.repo.UpdateAccount(amount, id)
}

func (s *DefaultWalletService) SaveCapturedTransaction(capture *domain.Capture) (*domain.Capture, error) {
	return s.repo.SaveCapturedTransaction(capture)
}

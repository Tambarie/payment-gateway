package service

import (
	"github.com/Tambarie/payment-gateway/application/payment-repository"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PaymentGatewayService interface
type PaymentGatewayService interface {
	CheckIfEmailExists(email string) (bson.M, error)
	CheckIfUserExists(userReference string) (bson.M, error)
	CreateUser(user *domain.User) (*domain.User, error)
	Authorise(card *domain.Card) (*domain.Card, error)
	GetCardByID(id string) (bson.M, error)
	UpdateAccount(amount float64, id string) (interface{}, error)
	SaveCapturedTransaction(capture *domain.Transaction) (*mongo.InsertOneResult, error)
	GetCapturedTransactionByTransactionID(id string) (*domain.Transaction, error)
	GetCapturedTransactionByAuthorizationID(id string) (bson.M, error)
	RefundUpdateAccount(amount float64, id string, count int) (interface{}, error)
	GetRefundTrackerByTransactionID(id string) (bson.M, error)
	SaveRefundTracker(tracker *domain.RefundTracker) (*mongo.InsertOneResult, error)
	VoidCard(id string, void bool) (interface{}, error)
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

func (s *DefaultWalletService) CheckIfEmailExists(email string) (bson.M, error) {
	return s.repo.CheckIfEmailExists(email)
}

func (s *DefaultWalletService) CheckIfUserExists(userReference string) (bson.M, error) {
	return s.repo.CheckIfUserExists(userReference)
}

func (s *DefaultWalletService) CreateUser(user *domain.User) (*domain.User, error) {
	return s.repo.CreateUser(user)
}

func (s *DefaultWalletService) Authorise(card *domain.Card) (*domain.Card, error) {
	return s.repo.Authorise(card)
}

func (s *DefaultWalletService) GetCardByID(id string) (bson.M, error) {
	return s.repo.GetCardByID(id)
}

func (s *DefaultWalletService) UpdateAccount(amount float64, id string) (interface{}, error) {
	return s.repo.UpdateAccount(amount, id)
}

func (s *DefaultWalletService) SaveCapturedTransaction(capture *domain.Transaction) (*mongo.InsertOneResult, error) {
	return s.repo.SaveCapturedTransaction(capture)
}

func (s *DefaultWalletService) GetCapturedTransactionByTransactionID(id string) (*domain.Transaction, error) {
	return s.repo.GetCapturedTransactionByTransactionID(id)
}

func (s *DefaultWalletService) GetCapturedTransactionByAuthorizationID(id string) (bson.M, error) {
	return s.repo.GetCapturedTransactionByAuthorizationID(id)
}

func (s *DefaultWalletService) RefundUpdateAccount(amount float64, id string, count int) (interface{}, error) {
	return s.repo.RefundUpdateAccount(amount, id, count)
}

func (s *DefaultWalletService) GetRefundTrackerByTransactionID(id string) (bson.M, error) {
	return s.repo.GetRefundTrackerByTransactionID(id)
}

func (s *DefaultWalletService) SaveRefundTracker(tracker *domain.RefundTracker) (*mongo.InsertOneResult, error) {
	return s.repo.SaveRefundTracker(tracker)
}

func (s *DefaultWalletService) VoidCard(id string, void bool) (interface{}, error) {
	return s.repo.VoidCard(id, void)
}

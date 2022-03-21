run: |
	gofmt -w .
	go run main.go

mock-service:
	mockgen -source=domain/service/gateway.go -destination=domain/service/mock_payment_gateway_db.go -package=service





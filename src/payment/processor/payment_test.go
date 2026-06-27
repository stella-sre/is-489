package processor

import (
	"errors"
	"testing"
	"time"
)

func TestPaymentProcessor_ProcessesValidPayment(t *testing.T) {
	proc := NewProcessor(NewInMemoryRepository(), time.Now())

	payment, err := proc.Process(PaymentRequest{
		UserID: "user-1",
		Amount: 1000,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if payment.Status != StatusCompleted {
		t.Errorf("status = %s; want %s", payment.Status, StatusCompleted)
	}
	if payment.Amount != 1000 {
		t.Errorf("amount = %.2f; want 1000", payment.Amount)
	}
}

func TestPaymentProcessor_AppliesTax(t *testing.T) {
	proc := NewProcessor(NewInMemoryRepository(), time.Now())

	payment, err := proc.Process(PaymentRequest{
		UserID: "user-1",
		Amount: 2000,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if payment.Tax != 200 {
		t.Errorf("tax = %.2f; want 200", payment.Tax)
	}
	if payment.Total != 2200 {
		t.Errorf("total = %.2f; want 2200", payment.Total)
	}
}

func TestPaymentProcessor_RejectsBelowMinimum(t *testing.T) {
	proc := NewProcessor(NewInMemoryRepository(), time.Now())

	_, err := proc.Process(PaymentRequest{
		UserID: "user-1",
		Amount: 5,
	})

	if err == nil {
		t.Fatal("expected error for amount below minimum, got nil")
	}
	if !errors.Is(err, ErrBelowMinimum) {
		t.Errorf("expected ErrBelowMinimum, got %v", err)
	}
}

func TestPaymentProcessor_RejectsNegativeAmount(t *testing.T) {
	proc := NewProcessor(NewInMemoryRepository(), time.Now())

	_, err := proc.Process(PaymentRequest{
		UserID: "user-1",
		Amount: -50,
	})

	if err == nil {
		t.Fatal("expected error for negative amount, got nil")
	}
}

func TestPaymentProcessor_RejectsExceedingDailyLimit(t *testing.T) {
	repo := NewInMemoryRepository()
	now := time.Date(2026, 6, 27, 10, 0, 0, 0, time.UTC)
	proc := NewProcessor(repo, now)

	_, err := proc.Process(PaymentRequest{UserID: "user-1", Amount: 2000})
	if err != nil {
		t.Fatalf("primer pago falló: %v", err)
	}
	_, err = proc.Process(PaymentRequest{UserID: "user-1", Amount: 2000})
	if err != nil {
		t.Fatalf("segundo pago falló: %v", err)
	}
	_, err = proc.Process(PaymentRequest{UserID: "user-1", Amount: 2000})
	if err == nil {
		t.Fatal("tercer pago debió fallar por límite diario, no falló")
	}
	if !errors.Is(err, ErrDailyLimitExceeded) {
		t.Errorf("expected ErrDailyLimitExceeded, got %v", err)
	}
}

func TestPaymentProcessor_DailyLimitIsPerUser(t *testing.T) {
	repo := NewInMemoryRepository()
	now := time.Date(2026, 6, 27, 10, 0, 0, 0, time.UTC)
	proc := NewProcessor(repo, now)

	_, err := proc.Process(PaymentRequest{UserID: "user-A", Amount: 4500})
	if err != nil {
		t.Fatalf("pago user-A falló: %v", err)
	}
	_, err = proc.Process(PaymentRequest{UserID: "user-B", Amount: 4500})
	if err != nil {
		t.Fatalf("pago user-B falló, el límite debe ser por usuario: %v", err)
	}
}

func TestPaymentProcessor_RequiresUserID(t *testing.T) {
	proc := NewProcessor(NewInMemoryRepository(), time.Now())

	_, err := proc.Process(PaymentRequest{
		UserID: "",
		Amount: 100,
	})

	if err == nil {
		t.Fatal("expected error for empty user id, got nil")
	}
	if !errors.Is(err, ErrMissingUser) {
		t.Errorf("expected ErrMissingUser, got %v", err)
	}
}
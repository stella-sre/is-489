package refund

import (
	"errors"
	"testing"
	"time"

	"tdd/src/payment/processor"
)

func TestRefundProcessor_FullRefund_OfCompletedPayment(t *testing.T) {
	repo := newMockRepo()
	proc := NewProcessor(repo, time.Now())
	payment := &processor.Payment{
		ID:        "pay-1",
		UserID:    "user-1",
		Amount:    1000,
		Total:     1100,
		Tax:       100,
		Status:    processor.StatusCompleted,
		CreatedAt: time.Now().Add(-1 * time.Hour),
	}
	repo.payments[payment.ID] = payment

	refund, err := proc.Refund(RefundRequest{
		PaymentID: "pay-1",
		Amount:    1000,
		Reason:    "Solicitud del cliente",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if refund.Amount != 1000 {
		t.Errorf("refund amount = %.2f; want 1000", refund.Amount)
	}
	if refund.Status != StatusProcessed {
		t.Errorf("status = %s; want %s", refund.Status, StatusProcessed)
	}
	if payment.Status != processor.StatusCompleted {
		t.Errorf("payment status = %s; want %s", payment.Status, processor.StatusCompleted)
	}
}

func TestRefundProcessor_PartialRefund_Allowed(t *testing.T) {
	repo := newMockRepo()
	proc := NewProcessor(repo, time.Now())
	repo.payments["pay-1"] = &processor.Payment{
		ID:        "pay-1",
		UserID:    "user-1",
		Amount:    1000,
		Total:     1100,
		Status:    processor.StatusCompleted,
		CreatedAt: time.Now().Add(-1 * time.Hour),
	}

	refund, err := proc.Refund(RefundRequest{
		PaymentID: "pay-1",
		Amount:    400,
		Reason:    "Devolución parcial",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if refund.Amount != 400 {
		t.Errorf("refund amount = %.2f; want 400", refund.Amount)
	}
}

func TestRefundProcessor_RejectsRefundOfFailedPayment(t *testing.T) {
	repo := newMockRepo()
	proc := NewProcessor(repo, time.Now())
	repo.payments["pay-1"] = &processor.Payment{
		ID:        "pay-1",
		UserID:    "user-1",
		Amount:    1000,
		Total:     1100,
		Status:    processor.StatusFailed,
		CreatedAt: time.Now().Add(-1 * time.Hour),
	}

	_, err := proc.Refund(RefundRequest{
		PaymentID: "pay-1",
		Amount:    1000,
		Reason:    "Intento",
	})

	if err == nil {
		t.Fatal("expected error refunding a failed payment")
	}
	if !errors.Is(err, ErrPaymentNotRefundable) {
		t.Errorf("expected ErrPaymentNotRefundable, got %v", err)
	}
}

func TestRefundProcessor_RejectsRefundOlderThan30Days(t *testing.T) {
	repo := newMockRepo()
	now := time.Date(2026, 6, 27, 10, 0, 0, 0, time.UTC)
	proc := NewProcessor(repo, now)
	repo.payments["pay-1"] = &processor.Payment{
		ID:        "pay-1",
		UserID:    "user-1",
		Amount:    1000,
		Total:     1100,
		Status:    processor.StatusCompleted,
		CreatedAt: now.AddDate(0, 0, -31),
	}

	_, err := proc.Refund(RefundRequest{
		PaymentID: "pay-1",
		Amount:    1000,
		Reason:    "Fuera de plazo",
	})

	if err == nil {
		t.Fatal("expected error for payment older than 30 days")
	}
	if !errors.Is(err, ErrRefundWindowExpired) {
		t.Errorf("expected ErrRefundWindowExpired, got %v", err)
	}
}

func TestRefundProcessor_AllowsRefundAtDay30Boundary(t *testing.T) {
	repo := newMockRepo()
	now := time.Date(2026, 6, 27, 10, 0, 0, 0, time.UTC)
	proc := NewProcessor(repo, now)
	repo.payments["pay-1"] = &processor.Payment{
		ID:        "pay-1",
		UserID:    "user-1",
		Amount:    1000,
		Total:     1100,
		Status:    processor.StatusCompleted,
		CreatedAt: now.AddDate(0, 0, -30),
	}

	_, err := proc.Refund(RefundRequest{
		PaymentID: "pay-1",
		Amount:    1000,
		Reason:    "Dentro del límite",
	})

	if err != nil {
		t.Fatalf("se esperaba reembolso exitoso en el límite de 30 días: %v", err)
	}
}

func TestRefundProcessor_RejectsExceedingOriginalAmount(t *testing.T) {
	repo := newMockRepo()
	proc := NewProcessor(repo, time.Now())
	repo.payments["pay-1"] = &processor.Payment{
		ID:        "pay-1",
		UserID:    "user-1",
		Amount:    1000,
		Total:     1100,
		Status:    processor.StatusCompleted,
		CreatedAt: time.Now().Add(-1 * time.Hour),
	}

	_, err := proc.Refund(RefundRequest{
		PaymentID: "pay-1",
		Amount:    1500,
		Reason:    "Mayor al pago",
	})

	if err == nil {
		t.Fatal("expected error when refund amount exceeds original")
	}
	if !errors.Is(err, ErrRefundExceedsPayment) {
		t.Errorf("expected ErrRefundExceedsPayment, got %v", err)
	}
}

func TestRefundProcessor_RejectsRefundOfNonExistingPayment(t *testing.T) {
	proc := NewProcessor(newMockRepo(), time.Now())

	_, err := proc.Refund(RefundRequest{
		PaymentID: "no-existe",
		Amount:    100,
		Reason:    "X",
	})

	if err == nil {
		t.Fatal("expected error for non-existing payment")
	}
}

func TestRefundProcessor_RejectsDoubleRefund(t *testing.T) {
	repo := newMockRepo()
	proc := NewProcessor(repo, time.Now())
	repo.payments["pay-1"] = &processor.Payment{
		ID:        "pay-1",
		UserID:    "user-1",
		Amount:    1000,
		Total:     1100,
		Status:    processor.StatusCompleted,
		CreatedAt: time.Now().Add(-1 * time.Hour),
	}

	_, err := proc.Refund(RefundRequest{PaymentID: "pay-1", Amount: 600, Reason: "primera"})
	if err != nil {
		t.Fatalf("primer reembolso falló: %v", err)
	}

	_, err = proc.Refund(RefundRequest{PaymentID: "pay-1", Amount: 500, Reason: "segunda"})
	if err == nil {
		t.Fatal("segundo reembolso debió fallar porque excede lo restante (600+500 > 1000)")
	}
}

func TestRefundProcessor_RefundsProportionalTax(t *testing.T) {
	repo := newMockRepo()
	proc := NewProcessor(repo, time.Now())
	repo.payments["pay-1"] = &processor.Payment{
		ID:        "pay-1",
		UserID:    "user-1",
		Amount:    2000,
		Total:     2200,
		Tax:       200,
		Status:    processor.StatusCompleted,
		CreatedAt: time.Now().Add(-1 * time.Hour),
	}

	refund, err := proc.Refund(RefundRequest{PaymentID: "pay-1", Amount: 1000, Reason: "mitad"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if refund.RefundedTax != 100 {
		t.Errorf("refunded tax = %.2f; want 100", refund.RefundedTax)
	}
}
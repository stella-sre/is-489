package tax

import (
	"errors"
	"testing"
)

func TestTaxCalculator_NoTax_Below1000(t *testing.T) {
	calc := NewCalculator()
	got, err := calc.Calculate(500)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 0.0

	if got != want {
		t.Errorf("Calculate(500) = %.2f; want %.2f", got, want)
	}
}

func TestTaxCalculator_TenPercent_Between1000And10000(t *testing.T) {
	calc := NewCalculator()
	got, err := calc.Calculate(2000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 200.0

	if got != want {
		t.Errorf("Calculate(2000) = %.2f; want %.2f", got, want)
	}
}

func TestTaxCalculator_TenPercent_AtBoundary(t *testing.T) {
	calc := NewCalculator()
	got, err := calc.Calculate(10000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 1000.0

	if got != want {
		t.Errorf("Calculate(10000) = %.2f; want %.2f", got, want)
	}
}

func TestTaxCalculator_FifteenPercent_Above10000(t *testing.T) {
	calc := NewCalculator()
	got, err := calc.Calculate(20000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 3000.0

	if got != want {
		t.Errorf("Calculate(20000) = %.2f; want %.2f", got, want)
	}
}

func TestTaxCalculator_ReturnsError_OnNegativeAmount(t *testing.T) {
	calc := NewCalculator()
	_, err := calc.Calculate(-100)

	if err == nil {
		t.Fatal("expected error for negative amount, got nil")
	}
	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got %v", err)
	}
}

func TestTaxCalculator_ReturnsError_OnZeroAmount(t *testing.T) {
	calc := NewCalculator()
	_, err := calc.Calculate(0)

	if err == nil {
		t.Fatal("expected error for zero amount, got nil")
	}
}
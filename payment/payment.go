package payment

import (
	"errors"
	"fmt"
	"math"
	"time"
)

const (
	TaxRate       = 0.18    // IGV Perú: 18%
	MinAmount     = 1.00    // Monto mínimo permitido (S/.)
	MaxDaily      = 5000.00 // Límite diario por usuario (S/.)
	MaxRefundDays = 30      // Máximo de días para solicitar reembolso
)

type TransactionStatus string

const (
	StatusCompleted TransactionStatus = "completado"
	StatusRefunded  TransactionStatus = "reembolsado"
	StatusPartial   TransactionStatus = "reembolso_parcial"
)

type Transaction struct {
	ID        string
	UserID    string
	Amount    float64
	Tax       float64
	Total     float64
	Timestamp time.Time
	Status    TransactionStatus
}

type RefundResult struct {
	RefundID     string
	OriginalTxID string
	RefundAmount float64
	RefundTax    float64
	RefundTotal  float64
	ProcessedAt  time.Time
}

type TaxCalculation struct {
	Subtotal float64
	Tax      float64
	Total    float64
}

// CalculateTax calcula el IGV (18%) sobre un subtotal dado.
// Retorna error si el monto es negativo o cero.
func CalculateTax(subtotal float64) (TaxCalculation, error) {
	if subtotal <= 0 {
		return TaxCalculation{}, errors.New("el subtotal debe ser mayor a cero")
	}

	tax := round2(subtotal * TaxRate)
	total := round2(subtotal + tax)

	return TaxCalculation{
		Subtotal: round2(subtotal),
		Tax:      tax,
		Total:    total,
	}, nil
}

// ExtractSubtotalFromTotal extrae el subtotal dado un total con IGV incluido.
func ExtractSubtotalFromTotal(totalWithTax float64) (TaxCalculation, error) {
	if totalWithTax <= 0 {
		return TaxCalculation{}, errors.New("el total debe ser mayor a cero")
	}

	subtotal := round2(totalWithTax / (1 + TaxRate))
	tax := round2(totalWithTax - subtotal)

	return TaxCalculation{
		Subtotal: subtotal,
		Tax:      tax,
		Total:    round2(totalWithTax),
	}, nil
}

// PaymentProcessor gestiona pagos y lleva registro de transacciones diarias.
type PaymentProcessor struct {
	dailyTotals  map[string]map[string]float64
	transactions map[string]*Transaction
	clock        func() time.Time
}

// NewPaymentProcessor crea un nuevo procesador de pagos.
func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{
		dailyTotals:  make(map[string]map[string]float64),
		transactions: make(map[string]*Transaction),
		clock:        time.Now,
	}
}

// NewPaymentProcessorWithClock crea un procesador con reloj personalizado (para pruebas).
func NewPaymentProcessorWithClock(clock func() time.Time) *PaymentProcessor {
	return &PaymentProcessor{
		dailyTotals:  make(map[string]map[string]float64),
		transactions: make(map[string]*Transaction),
		clock:        clock,
	}
}

// ProcessPayment valida y procesa un pago para un usuario.
func (p *PaymentProcessor) ProcessPayment(userID string, amount float64) (*Transaction, error) {
	if amount < MinAmount {
		return nil, fmt.Errorf("el monto S/. %.2f es menor al mínimo permitido S/. %.2f", amount, MinAmount)
	}

	now := p.clock()
	today := now.Format("2006-01-02")

	if p.dailyTotals[userID] == nil {
		p.dailyTotals[userID] = make(map[string]float64)
	}

	accumulated := p.dailyTotals[userID][today]
	if accumulated+amount > MaxDaily {
		remaining := MaxDaily - accumulated
		return nil, fmt.Errorf(
			"límite diario excedido: acumulado S/. %.2f, disponible S/. %.2f, solicitado S/. %.2f",
			accumulated, remaining, amount,
		)
	}

	calc, err := CalculateTax(amount)
	if err != nil {
		return nil, err
	}

	txID := fmt.Sprintf("TXN-%s-%d", userID, now.UnixNano())
	tx := &Transaction{
		ID:        txID,
		UserID:    userID,
		Amount:    calc.Subtotal,
		Tax:       calc.Tax,
		Total:     calc.Total,
		Timestamp: now,
		Status:    StatusCompleted,
	}

	p.dailyTotals[userID][today] += amount
	p.transactions[txID] = tx
	return tx, nil
}

// GetDailyTotal retorna el total acumulado del día para un usuario.
func (p *PaymentProcessor) GetDailyTotal(userID string) float64 {
	today := p.clock().Format("2006-01-02")
	if p.dailyTotals[userID] == nil {
		return 0
	}
	return p.dailyTotals[userID][today]
}

// GetTransaction busca una transacción por ID.
func (p *PaymentProcessor) GetTransaction(txID string) (*Transaction, error) {
	tx, ok := p.transactions[txID]
	if !ok {
		return nil, fmt.Errorf("transacción %s no encontrada", txID)
	}
	return tx, nil
}

// --- Procesamiento de Reembolsos ---

// RefundPolicy define las reglas de reembolso
type RefundPolicy struct {
	MaxDaysForRefund int
	MaxRefundPct     float64
}

// DefaultRefundPolicy es la política de reembolso estándar
var DefaultRefundPolicy = RefundPolicy{
	MaxDaysForRefund: MaxRefundDays,
	MaxRefundPct:     1.0,
}

// ProcessRefund procesa un reembolso para una transacción existente.
func (p *PaymentProcessor) ProcessRefund(txID string, refundAmount float64, policy RefundPolicy) (*RefundResult, error) {
	tx, err := p.GetTransaction(txID)
	if err != nil {
		return nil, err
	}

	if tx.Status == StatusRefunded {
		return nil, errors.New("la transacción ya fue completamente reembolsada")
	}

	now := p.clock()
	daysSince := now.Sub(tx.Timestamp).Hours() / 24
	if daysSince > float64(policy.MaxDaysForRefund) {
		return nil, fmt.Errorf(
			"han pasado %.0f días desde la transacción; el plazo máximo es %d días",
			daysSince, policy.MaxDaysForRefund,
		)
	}

	maxRefund := round2(tx.Total * policy.MaxRefundPct)
	if refundAmount <= 0 {
		return nil, errors.New("el monto de reembolso debe ser positivo")
	}
	if refundAmount > maxRefund {
		return nil, fmt.Errorf(
			"monto de reembolso S/. %.2f excede el máximo permitido S/. %.2f",
			refundAmount, maxRefund,
		)
	}

	refundTax := round2(refundAmount * (TaxRate / (1 + TaxRate)))
	refundSubtotal := round2(refundAmount - refundTax)

	if refundAmount >= tx.Total {
		tx.Status = StatusRefunded
	} else {
		tx.Status = StatusPartial
	}

	return &RefundResult{
		RefundID:     fmt.Sprintf("REF-%s-%d", txID, now.UnixNano()),
		OriginalTxID: txID,
		RefundAmount: refundSubtotal,
		RefundTax:    refundTax,
		RefundTotal:  refundAmount,
		ProcessedAt:  now,
	}, nil
}

// round2 redondea a 2 decimales para evitar errores de punto flotante
func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

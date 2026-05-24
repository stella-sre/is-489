package payment

import (
	"strings"
	"testing"
	"time"
)

func TestCalculateTax_SubtotalPositivo(t *testing.T) {
	subtotal := 100.00

	result, err := CalculateTax(subtotal)

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if result.Tax != 18.00 {
		t.Errorf("IGV esperado: 18.00, obtenido: %.2f", result.Tax)
	}
	if result.Total != 118.00 {
		t.Errorf("Total esperado: 118.00, obtenido: %.2f", result.Total)
	}
	if result.Subtotal != 100.00 {
		t.Errorf("Subtotal esperado: 100.00, obtenido: %.2f", result.Subtotal)
	}
}

func TestCalculateTax_MontoDecimal(t *testing.T) {
	subtotal := 55.50

	result, err := CalculateTax(subtotal)

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	expectedTax := 9.99
	if result.Tax != expectedTax {
		t.Errorf("IGV esperado: %.2f, obtenido: %.2f", expectedTax, result.Tax)
	}
	expectedTotal := 65.49
	if result.Total != expectedTotal {
		t.Errorf("Total esperado: %.2f, obtenido: %.2f", expectedTotal, result.Total)
	}
}

func TestCalculateTax_MontoCero_DebeRetornarError(t *testing.T) {
	subtotal := 0.00

	_, err := CalculateTax(subtotal)

	if err == nil {
		t.Error("se esperaba un error con subtotal = 0, pero no se obtuvo ninguno")
	}
}

func TestCalculateTax_MontoNegativo_DebeRetornarError(t *testing.T) {
	subtotal := -50.00

	_, err := CalculateTax(subtotal)

	if err == nil {
		t.Error("se esperaba un error con subtotal negativo, pero no se obtuvo ninguno")
	}
}

func TestExtractSubtotalFromTotal_TotalConIGV(t *testing.T) {
	totalWithTax := 118.00

	result, err := ExtractSubtotalFromTotal(totalWithTax)

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if result.Subtotal != 100.00 {
		t.Errorf("Subtotal esperado: 100.00, obtenido: %.2f", result.Subtotal)
	}
	if result.Tax != 18.00 {
		t.Errorf("IGV esperado: 18.00, obtenido: %.2f", result.Tax)
	}
}

func TestProcessPayment_PagoValido(t *testing.T) {
	processor := NewPaymentProcessor()
	userID := "user-001"
	amount := 200.00

	tx, err := processor.ProcessPayment(userID, amount)

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if tx == nil {
		t.Fatal("se esperaba una transacción, se obtuvo nil")
	}
	if tx.Status != StatusCompleted {
		t.Errorf("estado esperado: %s, obtenido: %s", StatusCompleted, tx.Status)
	}
	if tx.Amount != 200.00 {
		t.Errorf("monto esperado: 200.00, obtenido: %.2f", tx.Amount)
	}
	if tx.Tax != 36.00 {
		t.Errorf("IGV esperado: 36.00, obtenido: %.2f", tx.Tax)
	}
	if tx.Total != 236.00 {
		t.Errorf("total esperado: 236.00, obtenido: %.2f", tx.Total)
	}
}

func TestProcessPayment_MontoMenorAlMinimo_DebeRetornarError(t *testing.T) {
	processor := NewPaymentProcessor()
	amountBelowMin := 0.50

	tx, err := processor.ProcessPayment("user-002", amountBelowMin)

	if err == nil {
		t.Error("se esperaba error por monto menor al mínimo")
	}
	if tx != nil {
		t.Error("no se esperaba transacción cuando hay error de validación")
	}
	if !strings.Contains(err.Error(), "mínimo") {
		t.Errorf("mensaje de error esperado que contenga 'mínimo', obtenido: %s", err.Error())
	}
}

func TestProcessPayment_ExactamenteMontoMinimo_DebeProcearse(t *testing.T) {
	processor := NewPaymentProcessor()
	tx, err := processor.ProcessPayment("user-003", MinAmount)

	if err != nil {
		t.Fatalf("el monto mínimo exacto debe ser aceptado, error: %v", err)
	}
	if tx == nil {
		t.Fatal("se esperaba transacción para el monto mínimo exacto")
	}
}

func TestProcessPayment_LimiteDiarioExcedido_DebeRetornarError(t *testing.T) {
	processor := NewPaymentProcessor()
	userID := "user-004"
	processor.ProcessPayment(userID, 4800.00)

	tx, err := processor.ProcessPayment(userID, 300.00)

	if err == nil {
		t.Error("se esperaba error por límite diario excedido")
	}
	if tx != nil {
		t.Error("no se esperaba transacción cuando el límite diario es excedido")
	}
	if !strings.Contains(err.Error(), "límite diario") {
		t.Errorf("mensaje debería mencionar 'límite diario', obtenido: %s", err.Error())
	}
}

func TestProcessPayment_ExactamenteLimiteDiario_DebeAceptarse(t *testing.T) {
	processor := NewPaymentProcessor()
	userID := "user-005"
	processor.ProcessPayment(userID, 4000.00)

	tx, err := processor.ProcessPayment(userID, 1000.00)

	if err != nil {
		t.Fatalf("el límite diario exacto debe ser aceptado, error: %v", err)
	}
	if tx == nil {
		t.Fatal("se esperaba transacción para el límite diario exacto")
	}
}

func TestProcessPayment_MultiplesUsuariosIndependientes(t *testing.T) {
	processor := NewPaymentProcessor()
	processor.ProcessPayment("user-A", 4900.00)

	tx, err := processor.ProcessPayment("user-B", 4000.00)

	if err != nil {
		t.Errorf("user-B no debería verse afectado por el límite de user-A, error: %v", err)
	}
	if tx == nil {
		t.Error("se esperaba transacción para user-B")
	}
}

func TestGetDailyTotal_AcumuladoCorrecto(t *testing.T) {
	processor := NewPaymentProcessor()
	userID := "user-006"

	processor.ProcessPayment(userID, 100.00)
	processor.ProcessPayment(userID, 250.00)
	total := processor.GetDailyTotal(userID)

	if total != 350.00 {
		t.Errorf("total diario esperado: 350.00, obtenido: %.2f", total)
	}
}

func TestProcessRefund_ReembolsoTotal_Exitoso(t *testing.T) {
	processor := NewPaymentProcessor()
	tx, _ := processor.ProcessPayment("user-007", 100.00)

	refund, err := processor.ProcessRefund(tx.ID, tx.Total, DefaultRefundPolicy)

	if err != nil {
		t.Fatalf("no se esperaba error en reembolso total: %v", err)
	}
	if refund.RefundTotal != tx.Total {
		t.Errorf("monto de reembolso esperado: %.2f, obtenido: %.2f", tx.Total, refund.RefundTotal)
	}
	updatedTx, _ := processor.GetTransaction(tx.ID)
	if updatedTx.Status != StatusRefunded {
		t.Errorf("estado esperado: %s, obtenido: %s", StatusRefunded, updatedTx.Status)
	}
}

func TestProcessRefund_ReembolsoParcial_Exitoso(t *testing.T) {
	processor := NewPaymentProcessor()
	tx, _ := processor.ProcessPayment("user-008", 200.00)

	refund, err := processor.ProcessRefund(tx.ID, 100.00, DefaultRefundPolicy)

	if err != nil {
		t.Fatalf("no se esperaba error en reembolso parcial: %v", err)
	}
	if refund.RefundTotal != 100.00 {
		t.Errorf("monto esperado: 100.00, obtenido: %.2f", refund.RefundTotal)
	}
	updatedTx, _ := processor.GetTransaction(tx.ID)
	if updatedTx.Status != StatusPartial {
		t.Errorf("estado esperado: %s, obtenido: %s", StatusPartial, updatedTx.Status)
	}
}

func TestProcessRefund_TransaccionNoExiste_DebeRetornarError(t *testing.T) {
	processor := NewPaymentProcessor()

	_, err := processor.ProcessRefund("TXN-INEXISTENTE-999", 50.00, DefaultRefundPolicy)

	if err == nil {
		t.Error("se esperaba error para transacción inexistente")
	}
}

func TestProcessRefund_TransaccionYaReembolsada_DebeRetornarError(t *testing.T) {
	processor := NewPaymentProcessor()
	tx, _ := processor.ProcessPayment("user-009", 100.00)
	processor.ProcessRefund(tx.ID, tx.Total, DefaultRefundPolicy)

	_, err := processor.ProcessRefund(tx.ID, tx.Total, DefaultRefundPolicy)

	if err == nil {
		t.Error("se esperaba error al reembolsar una transacción ya reembolsada")
	}
	if !strings.Contains(err.Error(), "reembolsada") {
		t.Errorf("mensaje debería mencionar 'reembolsada', obtenido: %s", err.Error())
	}
}

func TestProcessRefund_MontoExcedeTotal_DebeRetornarError(t *testing.T) {
	processor := NewPaymentProcessor()
	tx, _ := processor.ProcessPayment("user-010", 100.00)

	_, err := processor.ProcessRefund(tx.ID, 200.00, DefaultRefundPolicy)

	if err == nil {
		t.Error("se esperaba error cuando el monto de reembolso excede el total")
	}
}

func TestProcessRefund_FueraDelPlazo_DebeRetornarError(t *testing.T) {
	pastTime := time.Now().AddDate(0, 0, -35)
	processor := NewPaymentProcessorWithClock(func() time.Time { return pastTime })
	tx, _ := processor.ProcessPayment("user-011", 100.00)

	processor.clock = func() time.Time { return time.Now() }

	_, err := processor.ProcessRefund(tx.ID, tx.Total, DefaultRefundPolicy)

	if err == nil {
		t.Error("se esperaba error cuando el reembolso está fuera del plazo")
	}
	if !strings.Contains(err.Error(), "días") {
		t.Errorf("mensaje debería mencionar 'días', obtenido: %s", err.Error())
	}
}

func TestProcessRefund_MontoNegativo_DebeRetornarError(t *testing.T) {
	processor := NewPaymentProcessor()
	tx, _ := processor.ProcessPayment("user-012", 100.00)

	_, err := processor.ProcessRefund(tx.ID, -50.00, DefaultRefundPolicy)

	if err == nil {
		t.Error("se esperaba error con monto de reembolso negativo")
	}
}

func TestFlujoCompleto_Pago_y_ReembolsoParcial(t *testing.T) {
	processor := NewPaymentProcessor()
	userID := "user-integracion"

	tx, err := processor.ProcessPayment(userID, 500.00)
	if err != nil {
		t.Fatalf("error en pago: %v", err)
	}

	daily := processor.GetDailyTotal(userID)
	if daily != 500.00 {
		t.Errorf("total diario esperado: 500.00, obtenido: %.2f", daily)
	}

	refund, err := processor.ProcessRefund(tx.ID, 59.00, DefaultRefundPolicy)
	if err != nil {
		t.Fatalf("error en reembolso: %v", err)
	}

	if refund.RefundTotal != 59.00 {
		t.Errorf("monto de reembolso esperado: 59.00, obtenido: %.2f", refund.RefundTotal)
	}
	expectedRefundTax := round2(59.00 * (TaxRate / (1 + TaxRate)))
	if refund.RefundTax != expectedRefundTax {
		t.Errorf("IGV del reembolso esperado: %.2f, obtenido: %.2f", expectedRefundTax, refund.RefundTax)
	}
	updatedTx, _ := processor.GetTransaction(tx.ID)
	if updatedTx.Status != StatusPartial {
		t.Errorf("estado esperado: %s, obtenido: %s", StatusPartial, updatedTx.Status)
	}
}

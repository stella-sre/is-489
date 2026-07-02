package sales

import (
	"erp/types"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	svc := NewSalesService()

	product := svc.CreateProduct("Laptop", 1500.00, 10, "Electronica")
	if product == nil {
		t.Fatal("se esperaba un producto creado")
	}
	if product.Name != "Laptop" {
		t.Errorf("nombre esperado: Laptop, obtenido: %s", product.Name)
	}
	if product.Price != 1500.00 {
		t.Errorf("precio esperado: 1500.00, obtenido: %f", product.Price)
	}
	if product.Stock != 10 {
		t.Errorf("stock esperado: 10, obtenido: %d", product.Stock)
	}
}

func TestCreateQuote(t *testing.T) {
	svc := NewSalesService()

	client := &types.Client{ID: "cli-001", Name: "Test Client", Email: "test@test.com"}
	svc.clients[client.ID] = client
	prod := svc.CreateProduct("Mouse", 25.00, 100, "Accesorios")

	items := []types.InvoiceItem{
		{ProductID: prod.ID, Quantity: 2},
	}

	quote, err := svc.CreateQuote(client.ID, items)
	if err != nil {
		t.Fatalf("error creando cotizacion: %v", err)
	}
	if quote.Subtotal != 50.00 {
		t.Errorf("subtotal esperado: 50.00, obtenido: %f", quote.Subtotal)
	}
	if quote.TaxAmount != 9.00 {
		t.Errorf("impuesto esperado: 9.00, obtenido: %f", quote.TaxAmount)
	}
	if quote.Total != 59.00 {
		t.Errorf("total esperado: 59.00, obtenido: %f", quote.Total)
	}
}

func TestCreateQuoteProductOutOfStock(t *testing.T) {
	svc := NewSalesService()

	client := &types.Client{ID: "cli-001", Name: "Test Client", Email: "test@test.com"}
	svc.clients[client.ID] = client
	prod := svc.CreateProduct("Teclado", 45.00, 0, "Accesorios")

	items := []types.InvoiceItem{
		{ProductID: prod.ID, Quantity: 5},
	}

	_, err := svc.CreateQuote(client.ID, items)
	if err != types.ErrProductOutOfStock {
		t.Fatalf("se esperaba ErrProductOutOfStock, obtenido: %v", err)
	}
}

func TestApproveQuoteGeneratesInvoice(t *testing.T) {
	svc := NewSalesService()

	client := &types.Client{ID: "cli-001", Name: "Test Client", Email: "test@test.com"}
	svc.clients[client.ID] = client
	prod := svc.CreateProduct("Monitor", 300.00, 5, "Electronica")

	items := []types.InvoiceItem{
		{ProductID: prod.ID, Quantity: 1},
	}

	quote, _ := svc.CreateQuote(client.ID, items)
	invoice, err := svc.ApproveQuote(quote.ID)
	if err != nil {
		t.Fatalf("error aprobando cotizacion: %v", err)
	}
	if invoice == nil {
		t.Fatal("se esperaba una factura creada")
	}
	if invoice.Status != types.InvoiceStatusPending {
		t.Errorf("status esperado: Pendiente, obtenido: %s", invoice.Status)
	}
	if prod.Stock != 4 {
		t.Errorf("stock esperado después de venta: 4, obtenido: %d", prod.Stock)
	}
}

func TestPayInvoice(t *testing.T) {
	svc := NewSalesService()

	client := &types.Client{ID: "cli-001", Name: "Test Client", Email: "test@test.com"}
	svc.clients[client.ID] = client
	prod := svc.CreateProduct("Impresora", 200.00, 10, "Electronica")

	items := []types.InvoiceItem{
		{ProductID: prod.ID, Quantity: 1},
	}

	quote, _ := svc.CreateQuote(client.ID, items)
	invoice, _ := svc.ApproveQuote(quote.ID)

	err := svc.PayInvoice(invoice.ID)
	if err != nil {
		t.Fatalf("error pagando factura: %v", err)
	}

	paidInvoice, _ := svc.GetInvoice(invoice.ID)
	if paidInvoice.Status != types.InvoiceStatusPaid {
		t.Errorf("status esperado: Pagada, obtenido: %s", paidInvoice.Status)
	}
}

func TestCalculateTax(t *testing.T) {
	svc := NewSalesService()

	tax := svc.CalculateTax(100.00)
	if tax != 18.00 {
		t.Errorf("impuesto esperado: 18.00, obtenido: %f", tax)
	}
}

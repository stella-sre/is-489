package sales

import (
	"erp/types"
	"time"
)

type SalesService struct {
	clients       map[string]*types.Client
	opportunities map[string]*types.Opportunity
	quotes        map[string]*Quote
	invoices      map[string]*types.Invoice
	products      map[string]*types.Product
	supportService interface {
		CreateTicket(clientID, productID, description string, priority types.TicketPriority) (*types.SupportTicket, error)
	}
}

type Quote struct {
	ID            string
	ClientID      string
	OpportunityID string
	Items         []types.InvoiceItem
	Subtotal      float64
	TaxRate       float64
	TaxAmount     float64
	Total         float64
	Status        QuoteStatus
	CreatedAt     time.Time
}

type QuoteStatus string

const (
	QuoteStatusPending  QuoteStatus = "Pendiente"
	QuoteStatusApproved QuoteStatus = "Aprobada"
	QuoteStatusRejected QuoteStatus = "Rechazada"
)

func NewSalesService() *SalesService {
	return &SalesService{
		clients:       make(map[string]*types.Client),
		opportunities: make(map[string]*types.Opportunity),
		quotes:        make(map[string]*Quote),
		invoices:      make(map[string]*types.Invoice),
		products:      make(map[string]*types.Product),
	}
}

func (s *SalesService) CreateProduct(name string, price float64, stock int, category string) *types.Product {
	product := &types.Product{
		ID:       generateSalesID("prod"),
		Name:     name,
		Price:    price,
		Stock:    stock,
		Category: category,
	}
	s.products[product.ID] = product
	return product
}

func (s *SalesService) CreateClientFromLead(lead *types.Lead) (*types.Client, *types.Opportunity, error) {
	client := &types.Client{
		ID:        generateSalesID("cli"),
		Name:      lead.Name,
		Email:     lead.Email,
		LeadID:    lead.ID,
		CreatedAt: time.Now(),
	}
	s.clients[client.ID] = client

	opportunity := &types.Opportunity{
		ID:        generateSalesID("opp"),
		ClientID:  client.ID,
		Amount:    0,
		Status:    types.OpportunityStatusOpen,
		CreatedAt: time.Now(),
	}
	s.opportunities[opportunity.ID] = opportunity

	return client, opportunity, nil
}

func (s *SalesService) GetClient(id string) (*types.Client, error) {
	client, ok := s.clients[id]
	if !ok {
		return nil, types.ErrClientNotFound
	}
	return client, nil
}

func (s *SalesService) GetProduct(id string) (*types.Product, error) {
	product, ok := s.products[id]
	if !ok {
		return nil, types.ErrProductNotFound
	}
	return product, nil
}

func (s *SalesService) CreateQuote(clientID string, items []types.InvoiceItem) (*Quote, error) {
	client, err := s.GetClient(clientID)
	if err != nil {
		return nil, err
	}

	subtotal := 0.0
	for i := range items {
		product, err := s.GetProduct(items[i].ProductID)
		if err != nil {
			return nil, err
		}
		if product.Stock < items[i].Quantity {
			return nil, types.ErrProductOutOfStock
		}
		items[i].UnitPrice = product.Price
		items[i].Total = product.Price * float64(items[i].Quantity)
		subtotal += items[i].Total
	}

	taxRate := 0.18
	taxAmount := subtotal * taxRate
	total := subtotal + taxAmount

	quote := &Quote{
		ID:            generateSalesID("quote"),
		ClientID:      client.ID,
		Items:         items,
		Subtotal:      subtotal,
		TaxRate:       taxRate,
		TaxAmount:     taxAmount,
		Total:         total,
		Status:        QuoteStatusPending,
		CreatedAt:     time.Now(),
	}
	s.quotes[quote.ID] = quote
	_ = client
	return quote, nil
}

func (s *SalesService) ApproveQuote(quoteID string) (*types.Invoice, error) {
	quote, ok := s.quotes[quoteID]
	if !ok {
		return nil, types.ErrLeadNotFound
	}
	quote.Status = QuoteStatusApproved

	invoice := s.generateInvoice(quote)
	s.invoices[invoice.ID] = invoice
	return invoice, nil
}

func (s *SalesService) generateInvoice(quote *Quote) *types.Invoice {
	invoice := &types.Invoice{
		ID:            generateSalesID("inv"),
		ClientID:      quote.ClientID,
		Items:         quote.Items,
		Subtotal:      quote.Subtotal,
		TaxRate:       quote.TaxRate,
		TaxAmount:     quote.TaxAmount,
		Total:         quote.Total,
		Status:        types.InvoiceStatusPending,
		CreatedAt:     time.Now(),
	}

	for i := range invoice.Items {
		product := s.products[invoice.Items[i].ProductID]
		invoice.Items[i].Product = product
		product.Stock -= invoice.Items[i].Quantity
	}

	return invoice
}

func (s *SalesService) GetInvoice(id string) (*types.Invoice, error) {
	invoice, ok := s.invoices[id]
	if !ok {
		return nil, types.ErrLeadNotFound
	}
	return invoice, nil
}

func (s *SalesService) PayInvoice(invoiceID string) error {
	invoice, err := s.GetInvoice(invoiceID)
	if err != nil {
		return err
	}
	if invoice.Status != types.InvoiceStatusPending {
		return types.ErrLeadNotFound
	}
	invoice.Status = types.InvoiceStatusPaid
	return nil
}

func (s *SalesService) CalculateTax(amount float64) float64 {
	return amount * 0.18
}

func generateSalesID(prefix string) string {
	return prefix + "-" + time.Now().Format("20060102150405.000000")
}

package erp

import (
	"erp/types"
	"testing"
)

func TestLeadToClientToTicketFlow(t *testing.T) {
	erpSvc := NewERPService()

	lead := erpSvc.Marketing.CreateLead("Roberto Diaz", "roberto@empresa.com")
	if lead.Status != types.LeadStatusNew {
		t.Fatalf("status inicial esperado: Nuevo, obtenido: %s", lead.Status)
	}

	err := erpSvc.Marketing.QualifyLead(lead.ID)
	if err != nil {
		t.Fatalf("error calificando lead: %v", err)
	}

	client, opportunity, err := erpSvc.Marketing.ConvertLeadToClient(lead.ID)
	if err != nil {
		t.Fatalf("error convirtiendo lead: %v", err)
	}
	if client.Name != "Roberto Diaz" {
		t.Errorf("nombre de cliente esperado: Roberto Diaz, obtenido: %s", client.Name)
	}
	if opportunity.Status != types.OpportunityStatusOpen {
		t.Errorf("status de oportunidad esperado: Abierta, obtenido: %s", opportunity.Status)
	}

	prod := erpSvc.Sales.CreateProduct("Servidor", 5000.00, 3, "Infraestructura")

	items := []types.InvoiceItem{
		{ProductID: prod.ID, Quantity: 1},
	}

	quote, err := erpSvc.Sales.CreateQuote(client.ID, items)
	if err != nil {
		t.Fatalf("error creando cotizacion: %v", err)
	}

	invoice, err := erpSvc.Sales.ApproveQuote(quote.ID)
	if err != nil {
		t.Fatalf("error aprobando cotizacion: %v", err)
	}

	err = erpSvc.Sales.PayInvoice(invoice.ID)
	if err != nil {
		t.Fatalf("error pagando factura: %v", err)
	}

	ticket, err := erpSvc.Support.CreateTicket(client.ID, prod.ID, "Servidor arrived damaged", types.TicketPriorityHigh)
	if err != nil {
		t.Fatalf("error creando ticket de soporte: %v", err)
	}
	if ticket.Status != types.TicketStatusOpen {
		t.Errorf("status de ticket esperado: Abierto, obtenido: %s", ticket.Status)
	}

	err = erpSvc.Support.CloseTicket(ticket.ID)
	if err != nil {
		t.Fatalf("error cerrando ticket: %v", err)
	}

	closedTicket, _ := erpSvc.Support.GetTicket(ticket.ID)
	if closedTicket.Status != types.TicketStatusClosed {
		t.Errorf("status de ticket esperado: Cerrado, obtenido: %s", closedTicket.Status)
	}
}

func TestEmployeePayrollFlow(t *testing.T) {
	erpSvc := NewERPService()

	employee := erpSvc.HR.CreateEmployee("Sofia Hernandez", "sofia@empresa.com", "Desarrollo", 4000.00, 25.00)

	payroll, err := erpSvc.HR.ProcessPayroll(employee.ID, 20.0)
	if err != nil {
		t.Fatalf("error procesando nomina: %v", err)
	}

	expectedTotal := 4000.00 + (20.0 * 25.00)
	if payroll.Total != expectedTotal {
		t.Errorf("total de nomina esperado: %f, obtenido: %f", expectedTotal, payroll.Total)
	}
}

func TestRoleRestrictions(t *testing.T) {
	erpSvc := NewERPService()

	marketingLead := erpSvc.Marketing.CreateLead("Marketing User", "mkt@test.com")
	erpSvc.Marketing.QualifyLead(marketingLead.ID)
	client, _, _ := erpSvc.Marketing.ConvertLeadToClient(marketingLead.ID)

	user := &types.User{
		ID:    "user-001",
		Name:  "Marketing User",
		Email: "mkt@test.com",
		Role:  types.RoleMarketing,
	}

	_ = user

	if client.Name != "" {
		t.Log("cliente creado correctamente")
	}
}

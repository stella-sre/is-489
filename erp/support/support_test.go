package support

import (
	"erp/types"
	"testing"
)

func TestCreateTicket(t *testing.T) {
	svc := NewSupportService()

	ticket, err := svc.CreateTicket("cli-001", "prod-001", "Pantalla no enciende", types.TicketPriorityHigh)
	if err != nil {
		t.Fatalf("error creando ticket: %v", err)
	}
	if ticket == nil {
		t.Fatal("se esperaba un ticket creado")
	}
	if ticket.Description != "Pantalla no enciende" {
		t.Errorf("descripcion esperada: Pantalla no enciende, obtenido: %s", ticket.Description)
	}
	if ticket.Priority != types.TicketPriorityHigh {
		t.Errorf("prioridad esperada: Alta, obtenido: %s", ticket.Priority)
	}
	if ticket.Status != types.TicketStatusOpen {
		t.Errorf("status esperado: Abierto, obtenido: %s", ticket.Status)
	}
}

func TestGetTicket(t *testing.T) {
	svc := NewSupportService()

	created, _ := svc.CreateTicket("cli-001", "prod-001", "Test ticket", types.TicketPriorityMedium)
	retrieved, err := svc.GetTicket(created.ID)
	if err != nil {
		t.Fatalf("error obteniendo ticket: %v", err)
	}
	if retrieved.Description != "Test ticket" {
		t.Errorf("descripcion esperada: Test ticket, obtenido: %s", retrieved.Description)
	}
}

func TestGetTicketNotFound(t *testing.T) {
	svc := NewSupportService()

	_, err := svc.GetTicket("no-existe")
	if err != types.ErrTicketNotFound {
		t.Fatalf("se esperaba ErrTicketNotFound, obtenido: %v", err)
	}
}

func TestUpdateTicketStatus(t *testing.T) {
	svc := NewSupportService()

	ticket, _ := svc.CreateTicket("cli-001", "prod-001", "Problema menor", types.TicketPriorityLow)
	err := svc.UpdateTicketStatus(ticket.ID, types.TicketStatusProgress)
	if err != nil {
		t.Fatalf("error actualizando status: %v", err)
	}

	updated, _ := svc.GetTicket(ticket.ID)
	if updated.Status != types.TicketStatusProgress {
		t.Errorf("status esperado: En Progreso, obtenido: %s", updated.Status)
	}
}

func TestCloseTicket(t *testing.T) {
	svc := NewSupportService()

	ticket, _ := svc.CreateTicket("cli-001", "prod-001", "Problema resuelto", types.TicketPriorityMedium)
	err := svc.CloseTicket(ticket.ID)
	if err != nil {
		t.Fatalf("error cerrando ticket: %v", err)
	}

	closed, _ := svc.GetTicket(ticket.ID)
	if closed.Status != types.TicketStatusClosed {
		t.Errorf("status esperado: Cerrado, obtenido: %s", closed.Status)
	}
}

func TestGetClientTickets(t *testing.T) {
	svc := NewSupportService()

	svc.CreateTicket("cli-001", "prod-001", "Ticket 1", types.TicketPriorityHigh)
	svc.CreateTicket("cli-001", "prod-002", "Ticket 2", types.TicketPriorityMedium)
	svc.CreateTicket("cli-002", "prod-001", "Ticket 3", types.TicketPriorityLow)

	tickets := svc.GetClientTickets("cli-001")
	if len(tickets) != 2 {
		t.Errorf("se esperaban 2 tickets para cliente cli-001, obtenido: %d", len(tickets))
	}
}

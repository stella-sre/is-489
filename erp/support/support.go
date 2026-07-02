package support

import (
	"erp/types"
	"sync/atomic"
	"time"
)

type SupportService struct {
	tickets     map[string]*types.SupportTicket
	salesService interface {
		GetInvoice(id string) (*types.Invoice, error)
		GetClient(id string) (*types.Client, error)
	}
	ticketCounter uint64
}

func NewSupportService() *SupportService {
	return &SupportService{
		tickets: make(map[string]*types.SupportTicket),
	}
}

func (s *SupportService) SetSalesService(svc interface {
	GetInvoice(id string) (*types.Invoice, error)
	GetClient(id string) (*types.Client, error)
}) {
	s.salesService = svc
}

func (s *SupportService) CreateTicket(clientID, productID, description string, priority types.TicketPriority) (*types.SupportTicket, error) {
	if s.salesService != nil {
		_, err := s.salesService.GetClient(clientID)
		if err != nil {
			return nil, err
		}
	}

	ticket := &types.SupportTicket{
		ID:          s.generateSupportID("tkt"),
		ClientID:    clientID,
		ProductID:   productID,
		Description: description,
		Priority:    priority,
		Status:      types.TicketStatusOpen,
		CreatedAt:   time.Now(),
	}
	s.tickets[ticket.ID] = ticket
	return ticket, nil
}

func (s *SupportService) GetTicket(id string) (*types.SupportTicket, error) {
	ticket, ok := s.tickets[id]
	if !ok {
		return nil, types.ErrTicketNotFound
	}
	return ticket, nil
}

func (s *SupportService) UpdateTicketStatus(id string, status types.TicketStatus) error {
	ticket, err := s.GetTicket(id)
	if err != nil {
		return err
	}
	ticket.Status = status
	return nil
}

func (s *SupportService) CloseTicket(id string) error {
	return s.UpdateTicketStatus(id, types.TicketStatusClosed)
}

func (s *SupportService) GetClientTickets(clientID string) []*types.SupportTicket {
	result := make([]*types.SupportTicket, 0)
	for _, ticket := range s.tickets {
		if ticket.ClientID == clientID {
			result = append(result, ticket)
		}
	}
	return result
}

func (s *SupportService) generateSupportID(prefix string) string {
	count := atomic.AddUint64(&s.ticketCounter, 1)
	return prefix + "-" + time.Now().Format("20060102150405.000000") + "-" + string(rune('0'+count))
}

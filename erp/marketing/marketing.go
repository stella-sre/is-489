package marketing

import (
	"erp/types"
	"erp/sales"
	"time"
)

type MarketingService struct {
	leads        map[string]*types.Lead
	salesService *sales.SalesService
}

func NewMarketingService(salesSvc *sales.SalesService) *MarketingService {
	return &MarketingService{
		leads:        make(map[string]*types.Lead),
		salesService: salesSvc,
	}
}

func (s *MarketingService) CreateLead(name, email string) *types.Lead {
	lead := &types.Lead{
		ID:      generateID("lead"),
		Name:    name,
		Email:   email,
		Status:  types.LeadStatusNew,
		Created: time.Now(),
	}
	s.leads[lead.ID] = lead
	return lead
}

func (s *MarketingService) GetLead(id string) (*types.Lead, error) {
	lead, ok := s.leads[id]
	if !ok {
		return nil, types.ErrLeadNotFound
	}
	return lead, nil
}

func (s *MarketingService) QualifyLead(id string) error {
	lead, err := s.GetLead(id)
	if err != nil {
		return err
	}
	if lead.Status != types.LeadStatusNew {
		return types.ErrLeadNotQualified
	}
	lead.Status = types.LeadStatusQualified
	return nil
}

func (s *MarketingService) ConvertLeadToClient(id string) (*types.Client, *types.Opportunity, error) {
	lead, err := s.GetLead(id)
	if err != nil {
		return nil, nil, err
	}
	if lead.Status != types.LeadStatusQualified {
		return nil, nil, types.ErrLeadNotQualified
	}

	client, opportunity, err := s.salesService.CreateClientFromLead(lead)
	if err != nil {
		return nil, nil, err
	}

	lead.Status = types.LeadStatusConverted
	return client, opportunity, nil
}

func (s *MarketingService) ListLeads() []*types.Lead {
	result := make([]*types.Lead, 0, len(s.leads))
	for _, lead := range s.leads {
		result = append(result, lead)
	}
	return result
}

func generateID(prefix string) string {
	return prefix + "-" + time.Now().Format("20060102150405.000000")
}

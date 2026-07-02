package erp

import (
	"erp/hr"
	"erp/marketing"
	"erp/sales"
	"erp/support"
)

type ERPService struct {
	Marketing *marketing.MarketingService
	Sales     *sales.SalesService
	HR        *hr.HRService
	Support   *support.SupportService
}

func NewERPService() *ERPService {
	salesSvc := sales.NewSalesService()
	supportSvc := support.NewSupportService()
	supportSvc.SetSalesService(salesSvc)

	return &ERPService{
		Marketing: marketing.NewMarketingService(salesSvc),
		Sales:     salesSvc,
		HR:        hr.NewHRService(),
		Support:   supportSvc,
	}
}

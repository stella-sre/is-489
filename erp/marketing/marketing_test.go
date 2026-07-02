package marketing

import (
	"erp/sales"
	"erp/types"
	"testing"
)

func TestCreateLead(t *testing.T) {
	salesSvc := sales.NewSalesService()
	svc := NewMarketingService(salesSvc)

	lead := svc.CreateLead("Juan Perez", "juan@example.com")
	if lead == nil {
		t.Fatal("se esperaba un lead creado")
	}
	if lead.Name != "Juan Perez" {
		t.Errorf("nombre esperado: Juan Perez, obtenido: %s", lead.Name)
	}
	if lead.Email != "juan@example.com" {
		t.Errorf("email esperado: juan@example.com, obtenido: %s", lead.Email)
	}
	if lead.Status != types.LeadStatusNew {
		t.Errorf("status esperado: Nuevo, obtenido: %s", lead.Status)
	}
}

func TestQualifyLead(t *testing.T) {
	salesSvc := sales.NewSalesService()
	svc := NewMarketingService(salesSvc)

	lead := svc.CreateLead("Maria Lopez", "maria@example.com")
	err := svc.QualifyLead(lead.ID)
	if err != nil {
		t.Fatalf("error calificando lead: %v", err)
	}

	qualifiedLead, _ := svc.GetLead(lead.ID)
	if qualifiedLead.Status != types.LeadStatusQualified {
		t.Errorf("status esperado: Calificado, obtenido: %s", qualifiedLead.Status)
	}
}

func TestConvertLeadToClient(t *testing.T) {
	salesSvc := sales.NewSalesService()
	svc := NewMarketingService(salesSvc)

	lead := svc.CreateLead("Carlos Garcia", "carlos@example.com")
	svc.QualifyLead(lead.ID)

	client, opportunity, err := svc.ConvertLeadToClient(lead.ID)
	if err != nil {
		t.Fatalf("error convirtiendo lead a cliente: %v", err)
	}
	if client == nil {
		t.Fatal("se esperaba un cliente creado")
	}
	if opportunity == nil {
		t.Fatal("se esperaba una oportunidad creada")
	}
	if client.Name != "Carlos Garcia" {
		t.Errorf("nombre esperado: Carlos Garcia, obtenido: %s", client.Name)
	}

	leadAgain, _ := svc.GetLead(lead.ID)
	if leadAgain.Status != types.LeadStatusConverted {
		t.Errorf("status del lead esperado: Convertido, obtenido: %s", leadAgain.Status)
	}
}

func TestConvertLeadNotQualified(t *testing.T) {
	salesSvc := sales.NewSalesService()
	svc := NewMarketingService(salesSvc)

	lead := svc.CreateLead("Pedro", "pedro@example.com")
	_, _, err := svc.ConvertLeadToClient(lead.ID)
	if err != types.ErrLeadNotQualified {
		t.Fatalf("se esperaba ErrLeadNotQualified, obtenido: %v", err)
	}
}

func TestGetLeadNotFound(t *testing.T) {
	salesSvc := sales.NewSalesService()
	svc := NewMarketingService(salesSvc)

	_, err := svc.GetLead("no-existe")
	if err != types.ErrLeadNotFound {
		t.Fatalf("se esperaba ErrLeadNotFound, obtenido: %v", err)
	}
}

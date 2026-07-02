package types

import (
	"errors"
	"time"
)

type Lead struct {
	ID      string
	Name    string
	Email   string
	Status  LeadStatus
	Created time.Time
}

type LeadStatus string

const (
	LeadStatusNew        LeadStatus = "Nuevo"
	LeadStatusQualified  LeadStatus = "Calificado"
	LeadStatusConverted  LeadStatus = "Convertido"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	LeadID    string
	CreatedAt time.Time
}

type Opportunity struct {
	ID        string
	ClientID  string
	Amount    float64
	Status    OpportunityStatus
	CreatedAt time.Time
}

type OpportunityStatus string

const (
	OpportunityStatusOpen  OpportunityStatus = "Abierta"
	OpportunityStatusWon   OpportunityStatus = "Ganada"
	OpportunityStatusLost OpportunityStatus = "Perdida"
)

type Product struct {
	ID       string
	Name     string
	Price    float64
	Stock    int
	Category string
}

type Invoice struct {
	ID            string
	ClientID      string
	OpportunityID string
	Items         []InvoiceItem
	Subtotal      float64
	TaxRate       float64
	TaxAmount     float64
	Total         float64
	Status        InvoiceStatus
	CreatedAt     time.Time
}

type InvoiceItem struct {
	ProductID string
	Product   *Product
	Quantity  int
	UnitPrice float64
	Total     float64
}

type InvoiceStatus string

const (
	InvoiceStatusPending   InvoiceStatus = "Pendiente"
	InvoiceStatusPaid      InvoiceStatus = "Pagada"
	InvoiceStatusCancelled InvoiceStatus = "Cancelada"
)

type Employee struct {
	ID           string
	Name         string
	Email        string
	Department   string
	BaseSalary   float64
	HourlyRate   float64
	WorkSchedule []time.Time
}

type Payroll struct {
	ID             string
	EmployeeID     string
	BaseSalary     float64
	ExtraHours     float64
	ExtraHoursPay  float64
	Total          float64
	ProcessedAt    time.Time
}

type SupportTicket struct {
	ID          string
	ClientID    string
	ProductID   string
	Description string
	Priority    TicketPriority
	Status      TicketStatus
	CreatedAt   time.Time
}

type TicketPriority string

const (
	TicketPriorityLow    TicketPriority = "Baja"
	TicketPriorityMedium TicketPriority = "Media"
	TicketPriorityHigh   TicketPriority = "Alta"
)

type TicketStatus string

const (
	TicketStatusOpen     TicketStatus = "Abierto"
	TicketStatusProgress TicketStatus = "En Progreso"
	TicketStatusClosed   TicketStatus = "Cerrado"
)

type Role string

const (
	RoleAdmin     Role = "Administrador"
	RoleSales     Role = "Ventas"
	RoleMarketing Role = "Marketing"
	RoleHR        Role = "RRHH"
	RoleSupport   Role = "Soporte"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     Role
}

var (
	ErrLeadNotFound      = errors.New("lead no encontrado")
	ErrLeadNotQualified  = errors.New("lead no está en estado calificado")
	ErrClientNotFound    = errors.New("cliente no encontrado")
	ErrProductNotFound   = errors.New("producto no encontrado")
	ErrProductOutOfStock = errors.New("producto sin stock")
	ErrUnauthorized      = errors.New("acceso no autorizado")
	ErrTicketNotFound    = errors.New("ticket no encontrado")
	ErrEmployeeNotFound  = errors.New("empleado no encontrado")
)

package erp

import "erp/types"

type Lead = types.Lead
type LeadStatus = types.LeadStatus
type Client = types.Client
type Opportunity = types.Opportunity
type OpportunityStatus = types.OpportunityStatus
type Product = types.Product
type Invoice = types.Invoice
type InvoiceItem = types.InvoiceItem
type InvoiceStatus = types.InvoiceStatus
type Employee = types.Employee
type Payroll = types.Payroll
type SupportTicket = types.SupportTicket
type TicketPriority = types.TicketPriority
type TicketStatus = types.TicketStatus
type Role = types.Role
type User = types.User

const (
	LeadStatusNew        = types.LeadStatusNew
	LeadStatusQualified  = types.LeadStatusQualified
	LeadStatusConverted  = types.LeadStatusConverted
	OpportunityStatusOpen = types.OpportunityStatusOpen
	OpportunityStatusWon  = types.OpportunityStatusWon
	OpportunityStatusLost = types.OpportunityStatusLost
	InvoiceStatusPending   = types.InvoiceStatusPending
	InvoiceStatusPaid      = types.InvoiceStatusPaid
	InvoiceStatusCancelled = types.InvoiceStatusCancelled
	TicketPriorityLow    = types.TicketPriorityLow
	TicketPriorityMedium = types.TicketPriorityMedium
	TicketPriorityHigh   = types.TicketPriorityHigh
	TicketStatusOpen     = types.TicketStatusOpen
	TicketStatusProgress = types.TicketStatusProgress
	TicketStatusClosed   = types.TicketStatusClosed
	RoleAdmin     = types.RoleAdmin
	RoleSales     = types.RoleSales
	RoleMarketing = types.RoleMarketing
	RoleHR        = types.RoleHR
	RoleSupport   = types.RoleSupport
)

var (
	ErrLeadNotFound      = types.ErrLeadNotFound
	ErrLeadNotQualified  = types.ErrLeadNotQualified
	ErrClientNotFound    = types.ErrClientNotFound
	ErrProductNotFound   = types.ErrProductNotFound
	ErrProductOutOfStock = types.ErrProductOutOfStock
	ErrUnauthorized      = types.ErrUnauthorized
	ErrTicketNotFound    = types.ErrTicketNotFound
	ErrEmployeeNotFound  = types.ErrEmployeeNotFound
)

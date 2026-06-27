package refund

import (
	"errors"
	"sync"
	"time"

	"tdd/src/payment/processor"
)

const refundWindowDays = 30

var (
	ErrPaymentNotFound     = errors.New("refund: payment not found")
	ErrPaymentNotRefundable = errors.New("refund: payment is not refundable")
	ErrRefundWindowExpired = errors.New("refund: refund window expired (>30 days)")
	ErrRefundExceedsPayment = errors.New("refund: amount exceeds payment total")
	ErrInvalidRefundAmount = errors.New("refund: invalid refund amount")
)

type Status string

const (
	StatusProcessed Status = "processed"
	StatusRejected  Status = "rejected"
)

type RefundRequest struct {
	PaymentID string
	Amount    float64
	Reason    string
}

type Refund struct {
	ID          string
	PaymentID   string
	Amount      float64
	RefundedTax float64
	Reason      string
	Status      Status
	CreatedAt   time.Time
}

type Repository interface {
	FindPayment(id string) (*processor.Payment, error)
	TotalRefundedFor(paymentID string) (float64, error)
	SaveRefund(r *Refund) error
}

type InMemoryRepository struct {
	mu       sync.Mutex
	payments map[string]*processor.Payment
	refunds  []*Refund
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{payments: map[string]*processor.Payment{}}
}

func (r *InMemoryRepository) Save(p *processor.Payment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.payments[p.ID] = p
	return nil
}

func (r *InMemoryRepository) FindPayment(id string) (*processor.Payment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.payments[id], nil
}

func (r *InMemoryRepository) TotalRefundedFor(paymentID string) (float64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var total float64
	for _, rf := range r.refunds {
		if rf.PaymentID == paymentID && rf.Status == StatusProcessed {
			total += rf.Amount
		}
	}
	return total, nil
}

func (r *InMemoryRepository) SaveRefund(rf *Refund) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.refunds = append(r.refunds, rf)
	return nil
}

type Processor struct {
	repo Repository
	now  func() time.Time
}

func NewProcessor(repo Repository, now time.Time) *Processor {
	return &Processor{repo: repo, now: func() time.Time { return now }}
}

func (p *Processor) Refund(req RefundRequest) (*Refund, error) {
	if req.Amount <= 0 {
		return nil, ErrInvalidRefundAmount
	}

	payment, err := p.repo.FindPayment(req.PaymentID)
	if err != nil {
		return nil, err
	}
	if payment == nil {
		return nil, ErrPaymentNotFound
	}
	if payment.Status != processor.StatusCompleted {
		return nil, ErrPaymentNotRefundable
	}

	now := p.now()
	if now.Sub(payment.CreatedAt) > time.Duration(refundWindowDays)*24*time.Hour {
		return nil, ErrRefundWindowExpired
	}

	alreadyRefunded, err := p.repo.TotalRefundedFor(req.PaymentID)
	if err != nil {
		return nil, err
	}
	if alreadyRefunded+req.Amount > payment.Amount {
		return nil, ErrRefundExceedsPayment
	}

	refundedTax := payment.Tax * (req.Amount / payment.Amount)
	refund := &Refund{
		PaymentID:   req.PaymentID,
		Amount:      req.Amount,
		RefundedTax: refundedTax,
		Reason:      req.Reason,
		Status:      StatusProcessed,
		CreatedAt:   now,
	}
	if err := p.repo.SaveRefund(refund); err != nil {
		return nil, err
	}
	return refund, nil
}
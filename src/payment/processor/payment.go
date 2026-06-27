package processor

import (
	"errors"
	"sync"
	"time"

	"tdd/src/payment/tax"
)

const (
	minAmount   = 10.0
	dailyLimit  = 5000.0
)

var (
	ErrBelowMinimum      = errors.New("payment: amount below minimum")
	ErrDailyLimitExceeded = errors.New("payment: daily limit exceeded")
	ErrMissingUser       = errors.New("payment: user id required")
	ErrInvalidAmount     = errors.New("payment: invalid amount")
)

type Status string

const (
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type PaymentRequest struct {
	UserID string
	Amount float64
}

type Payment struct {
	ID        string
	UserID    string
	Amount    float64
	Tax       float64
	Total     float64
	Status    Status
	CreatedAt time.Time
}

type Repository interface {
	Save(p *Payment) error
	FindByUserAndDate(userID string, date time.Time) ([]*Payment, error)
}

type InMemoryRepository struct {
	mu       sync.Mutex
	payments []*Payment
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (r *InMemoryRepository) Save(p *Payment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.payments = append(r.payments, p)
	return nil
}

func (r *InMemoryRepository) FindByUserAndDate(userID string, date time.Time) ([]*Payment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var out []*Payment
	for _, p := range r.payments {
		if p.UserID == userID && sameDay(p.CreatedAt, date) {
			out = append(out, p)
		}
	}
	return out, nil
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}

type Processor struct {
	repo    Repository
	tax     *tax.Calculator
	now     func() time.Time
}

func NewProcessor(repo Repository, now time.Time) *Processor {
	return &Processor{
		repo: repo,
		tax:  tax.NewCalculator(),
		now:  func() time.Time { return now },
	}
}

func (p *Processor) Process(req PaymentRequest) (*Payment, error) {
	if req.UserID == "" {
		return nil, ErrMissingUser
	}
	if req.Amount <= 0 {
		return nil, ErrInvalidAmount
	}
	if req.Amount < minAmount {
		return nil, ErrBelowMinimum
	}

	today := p.now()
	existing, err := p.repo.FindByUserAndDate(req.UserID, today)
	if err != nil {
		return nil, err
	}
	var todayTotal float64
	for _, e := range existing {
		todayTotal += e.Total
	}
	if todayTotal+req.Amount > dailyLimit {
		return nil, ErrDailyLimitExceeded
	}

	taxAmount, err := p.tax.Calculate(req.Amount)
	if err != nil {
		return nil, err
	}

	payment := &Payment{
		UserID:    req.UserID,
		Amount:    req.Amount,
		Tax:       taxAmount,
		Total:     req.Amount + taxAmount,
		Status:    StatusCompleted,
		CreatedAt: today,
	}
	if err := p.repo.Save(payment); err != nil {
		return nil, err
	}
	return payment, nil
}
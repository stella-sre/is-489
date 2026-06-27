package refund

import (
	"tdd/src/payment/processor"
)

type mockRepo struct {
	payments map[string]*processor.Payment
	refunds  []*Refund
}

func newMockRepo() *mockRepo {
	return &mockRepo{payments: map[string]*processor.Payment{}}
}

func (m *mockRepo) FindPayment(id string) (*processor.Payment, error) {
	return m.payments[id], nil
}

func (m *mockRepo) TotalRefundedFor(paymentID string) (float64, error) {
	var total float64
	for _, r := range m.refunds {
		if r.PaymentID == paymentID && r.Status == StatusProcessed {
			total += r.Amount
		}
	}
	return total, nil
}

func (m *mockRepo) SaveRefund(r *Refund) error {
	m.refunds = append(m.refunds, r)
	return nil
}
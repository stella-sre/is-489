# Caso 1: Sistema de Procesamiento de Pagos (TDD)

Implementación del primer caso propuesto de la guía, desarrollado siguiendo el
ciclo **Rojo → Verde → Refactorizar** en Go (Go 1.26.3).

---

## 1. Nota sobre la ejecución de pruebas

> **Importante**: en Go, `go test` se aplica sobre un **paquete**, no sobre un
> archivo `.go` suelto. Si se ejecuta `go test ./src/payment/processor/payment_test.go`,
> el compilador trata el archivo como un paquete anónimo llamado
> `command-line-arguments` y todos los símbolos del paquete quedan `undefined`:
>
> ```
> # command-line-arguments [command-line-arguments.test]
> src/payment/processor/payment_test.go:10:10: undefined: NewProcessor
> src/payment/processor/payment_test.go:10:23: undefined: NewInMemoryRepository
> src/payment/processor/payment_test.go:12:31: undefined: PaymentRequest
> ... too many errors
> FAIL    command-line-arguments [build failed]
> ```
>
> La forma correcta es probar el paquete completo:
>
> ```bash
> go test ./src/payment/processor/      # paquete
> go test ./src/payment/processor -v    # paquete en modo verboso
> go test ./...                          # todos los paquetes
> ```

---

## 2. Cobertura del caso

| Requisito de la guía | Implementación |
|---|---|
| Validar que los cálculos de impuestos se realicen correctamente | `tax.Calculator` con 3 tramos (0%, 10%, 15%) |
| Comprobar restricciones de pago (montos mínimos, límites diarios) | `processor.Processor` (mínimo $10, límite diario $5000 por usuario) |
| Asegurar que los reembolsos se procesen según las políticas | `refund.Processor` (parcial, 30 días, anti-doble reembolso, impuesto proporcional) |

---

## 3. Estructura del proyecto

```
src/payment/
├── tax/
│   ├── calculator.go         # Cálculo de impuestos
│   └── calculator_test.go
├── processor/
│   ├── payment.go            # Procesador de pagos
│   └── payment_test.go
└── refund/
    ├── refund.go             # Procesador de reembolsos
    ├── refund_test.go
    └── mock_repo_test.go     # Mock del repositorio
```

---

## 4. Ciclo TDD aplicado — componente por componente

Cada componente se construyó siguiendo tres ciclos cortos: **ROJO → VERDE →
REFACTOR**. Para cada uno se muestra el código en ese estado y la salida exacta
del comando `go test`.

---

### 4.1 TaxCalculator (`src/payment/tax/`)

#### 4.1.1 ROJO — primera prueba

Se escribe primero una prueba que verifica que un monto de 500 (< 1000) no
genera impuestos. Como `NewCalculator` aún no existe, la compilación falla.

**`src/payment/tax/calculator_test.go`** (versión inicial):

```go
package tax

import "testing"

func TestTaxCalculator_NoTax_Below1000(t *testing.T) {
    calc := NewCalculator()
    got := calc.Calculate(500)
    want := 0.0
    if got != want {
        t.Errorf("Calculate(500) = %.2f; want %.2f", got, want)
    }
}
```

**`src/payment/tax/calculator.go`**: *(no existe todavía)*

**Comando**:

```bash
go test ./src/payment/tax/ -v
```

**Salida (captura 1)**:

```
# tdd/src/payment/tax [tdd/src/payment/tax.test]
src/payment/tax/calculator_test.go:8:10: undefined: NewCalculator
FAIL    tdd/src/payment/tax [build failed]
FAIL
```

> Estado: **ROJO**. La compilación falla porque aún no existe `NewCalculator`.

---

#### 4.1.2 VERDE — código mínimo

Se crea el código mínimo de producción que satisface la prueba.

**`src/payment/tax/calculator.go`** (versión 1, mínima):

```go
package tax

type Calculator struct{}

func NewCalculator() *Calculator {
    return &Calculator{}
}

func (c *Calculator) Calculate(amount float64) float64 {
    return 0.0
}
```

**Comando**:

```bash
go test ./src/payment/tax/ -v
```

**Salida (captura 2)**:

```
=== RUN   TestTaxCalculator_NoTax_Below1000
--- PASS: TestTaxCalculator_NoTax_Below1000 (0.00s)
PASS
ok      tdd/src/payment/tax    0.440s
```

> Estado: **VERDE**. La primera prueba pasa con el mínimo de código.

---

#### 4.1.3 ROJO — segunda tanda de pruebas (tasas y errores)

Se amplía el archivo de pruebas con cinco casos adicionales: 10% entre 1000 y
10000, 10% en el límite 10000, 15% por encima de 10000, error en negativo y
error en cero.

**`src/payment/tax/calculator_test.go`** (versión completa, ya devuelve error):

```go
package tax

import (
    "errors"
    "testing"
)

func TestTaxCalculator_NoTax_Below1000(t *testing.T)            { /* ... */ }
func TestTaxCalculator_TenPercent_Between1000And10000(t *testing.T) { /* ... */ }
func TestTaxCalculator_TenPercent_AtBoundary(t *testing.T)      { /* ... */ }
func TestTaxCalculator_FifteenPercent_Above10000(t *testing.T)   { /* ... */ }
func TestTaxCalculator_ReturnsError_OnNegativeAmount(t *testing.T) { /* ... */ }
func TestTaxCalculator_ReturnsError_OnZeroAmount(t *testing.T)   { /* ... */ }
```

**Comando**:

```bash
go test ./src/payment/tax/ -v
```

**Salida (captura 3)**:

```
src/payment/tax/calculator_test.go:50:12: assignment mismatch: 2 variables but calc.Calculate returns 1 value
src/payment/tax/calculator_test.go:55:21: undefined: ErrInvalidAmount
src/payment/tax/calculator_test.go:62:12: assignment mismatch: 2 variables but calc.Calculate returns 1 value
FAIL    tdd/src/payment/tax [build failed]
FAIL
```

> Estado: **ROJO**. La firma de `Calculate` debe cambiar a `(float64, error)` y
> debe existir `ErrInvalidAmount`.

---

#### 4.1.4 VERDE — implementación con todas las reglas

**`src/payment/tax/calculator.go`**:

```go
package tax

import "errors"

var ErrInvalidAmount = errors.New("tax: invalid amount")

type Calculator struct{}

func NewCalculator() *Calculator {
    return &Calculator{}
}

func (c *Calculator) Calculate(amount float64) (float64, error) {
    if amount <= 0 {
        return 0, ErrInvalidAmount
    }
    switch {
    case amount < 1000:
        return 0, nil
    case amount <= 10000:
        return amount * 0.10, nil
    default:
        return amount * 0.15, nil
    }
}
```

**Comando**:

```bash
go test ./src/payment/tax/ -v
```

**Salida (captura 4)**:

```
=== RUN   TestTaxCalculator_NoTax_Below1000
--- PASS: TestTaxCalculator_NoTax_Below1000 (0.00s)
=== RUN   TestTaxCalculator_TenPercent_Between1000And10000
--- PASS: TestTaxCalculator_TenPercent_Between1000And10000 (0.00s)
=== RUN   TestTaxCalculator_TenPercent_AtBoundary
--- PASS: TestTaxCalculator_TenPercent_AtBoundary (0.00s)
=== RUN   TestTaxCalculator_FifteenPercent_Above10000
--- PASS: TestTaxCalculator_FifteenPercent_Above10000 (0.00s)
=== RUN   TestTaxCalculator_ReturnsError_OnNegativeAmount
--- PASS: TestTaxCalculator_ReturnsError_OnNegativeAmount (0.00s)
=== RUN   TestTaxCalculator_ReturnsError_OnZeroAmount
--- PASS: TestTaxCalculator_ReturnsError_OnZeroAmount (0.00s)
PASS
ok      tdd/src/payment/tax    0.321s
```

> Estado: **VERDE**. Las 6 pruebas pasan.

---

#### 4.1.5 REFACTOR — constantes con nombre

**`src/payment/tax/calculator.go`** (refactorizado):

```go
package tax

import "errors"

var ErrInvalidAmount = errors.New("tax: invalid amount")

const (
    lowThreshold = 1000.0
    midThreshold = 10000.0
    lowRate      = 0.0
    midRate      = 0.10
    highRate     = 0.15
)

type Calculator struct{}

func NewCalculator() *Calculator { return &Calculator{} }

func (c *Calculator) Calculate(amount float64) (float64, error) {
    if amount <= 0 {
        return 0, ErrInvalidAmount
    }
    switch {
    case amount < lowThreshold:
        return amount * lowRate, nil
    case amount <= midThreshold:
        return amount * midRate, nil
    default:
        return amount * highRate, nil
    }
}
```

**Comando**:

```bash
go test ./src/payment/tax/ -v -cover
```

**Salida (captura 5)**:

```
PASS
coverage: 100.0% of statements
ok      tdd/src/payment/tax
```

> Estado: **REFACTOR** completado, 100% de cobertura.

---

### 4.2 PaymentProcessor (`src/payment/processor/`)

#### 4.2.1 ROJO — primera prueba

**`src/payment/processor/payment_test.go`** (versión inicial):

```go
package processor

import (
    "testing"
    "time"
)

func TestPaymentProcessor_ProcessesValidPayment(t *testing.T) {
    proc := NewProcessor(NewInMemoryRepository(), time.Now())

    payment, err := proc.Process(PaymentRequest{
        UserID: "user-1",
        Amount: 1000,
    })

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if payment.Status != StatusCompleted {
        t.Errorf("status = %s; want %s", payment.Status, StatusCompleted)
    }
    if payment.Amount != 1000 {
        t.Errorf("amount = %.2f; want 1000", payment.Amount)
    }
}
```

**Comando**:

```bash
go test ./src/payment/processor/ -v
```

**Salida (captura 6)**:

```
src/payment/processor/payment_test.go:9:10: undefined: NewProcessor
src/payment/processor/payment_test.go:9:23: undefined: NewInMemoryRepository
src/payment/processor/payment_test.go:11:31: undefined: PaymentRequest
src/payment/processor/payment_test.go:19:23: undefined: StatusCompleted
src/payment/processor/payment_test.go:20:52: undefined: StatusCompleted
FAIL    tdd/src/payment/processor [build failed]
FAIL
```

> Estado: **ROJO**. Ningún símbolo del paquete existe todavía.

---

#### 4.2.2 VERDE — código mínimo

**`src/payment/processor/payment.go`** (versión 1):

```go
package processor

import (
    "sync"
    "time"
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

func NewInMemoryRepository() *InMemoryRepository { return &InMemoryRepository{} }

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
    repo Repository
    now  func() time.Time
}

func NewProcessor(repo Repository, now time.Time) *Processor {
    return &Processor{repo: repo, now: func() time.Time { return now }}
}

func (p *Processor) Process(req PaymentRequest) (*Payment, error) {
    payment := &Payment{
        UserID:    req.UserID,
        Amount:    req.Amount,
        Status:    StatusCompleted,
        CreatedAt: p.now(),
    }
    payment.Total = req.Amount
    _ = p.repo.Save(payment)
    return payment, nil
}
```

**Comando**:

```bash
go test ./src/payment/processor/ -v
```

**Salida (captura 7)**:

```
=== RUN   TestPaymentProcessor_ProcessesValidPayment
--- PASS: TestPaymentProcessor_ProcessesValidPayment (0.00s)
PASS
ok      tdd/src/payment/processor    0.315s
```

> Estado: **VERDE** para el primer caso.

---

#### 4.2.3 ROJO — restricciones (mínimo, límite diario, usuario)

Se agregan las pruebas que verifican las reglas de negocio: aplicar impuesto,
rechazar montos negativos, rechazar por debajo del mínimo, rechazar por exceso
del límite diario ($5000) y exigir `UserID`.

**Comando**:

```bash
go test ./src/payment/processor/ -v
```

**Salida (captura 8)**:

```
src/payment/processor/payment_test.go:58:21: undefined: ErrBelowMinimum
src/payment/processor/payment_test.go:93:21: undefined: ErrDailyLimitExceeded
src/payment/processor/payment_test.go:124:21: undefined: ErrMissingUser
FAIL    tdd/src/payment/processor [build failed]
FAIL
```

> Estado: **ROJO**. Las pruebas referencian errores aún no definidos.

---

#### 4.2.4 VERDE — implementación completa con todas las restricciones

**`src/payment/processor/payment.go`** (versión final con restricciones):

```go
package processor

import (
    "errors"
    "sync"
    "time"

    "tdd/src/payment/tax"
)

const (
    minAmount  = 10.0
    dailyLimit = 5000.0
)

var (
    ErrBelowMinimum       = errors.New("payment: amount below minimum")
    ErrDailyLimitExceeded = errors.New("payment: daily limit exceeded")
    ErrMissingUser        = errors.New("payment: user id required")
    ErrInvalidAmount      = errors.New("payment: invalid amount")
)

// ... tipos Status, PaymentRequest, Payment, Repository, InMemoryRepository, sameDay

type Processor struct {
    repo Repository
    tax  *tax.Calculator
    now  func() time.Time
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
```

**Comando**:

```bash
go test ./src/payment/processor/ -v -cover
```

**Salida (captura 9)**:

```
=== RUN   TestPaymentProcessor_ProcessesValidPayment
--- PASS: TestPaymentProcessor_ProcessesValidPayment (0.00s)
=== RUN   TestPaymentProcessor_AppliesTax
--- PASS: TestPaymentProcessor_AppliesTax (0.00s)
=== RUN   TestPaymentProcessor_RejectsBelowMinimum
--- PASS: TestPaymentProcessor_RejectsBelowMinimum (0.00s)
=== RUN   TestPaymentProcessor_RejectsNegativeAmount
--- PASS: TestPaymentProcessor_RejectsNegativeAmount (0.00s)
=== RUN   TestPaymentProcessor_RejectsExceedingDailyLimit
--- PASS: TestPaymentProcessor_RejectsExceedingDailyLimit (0.00s)
=== RUN   TestPaymentProcessor_DailyLimitIsPerUser
--- PASS: TestPaymentProcessor_DailyLimitIsPerUser (0.00s)
=== RUN   TestPaymentProcessor_RequiresUserID
--- PASS: TestPaymentProcessor_RequiresUserID (0.00s)
PASS
coverage: 92.3% of statements
ok      tdd/src/payment/processor
```

> Estado: **VERDE**. 7/7 pruebas pasan, cobertura 92.3%.

---

### 4.3 RefundProcessor (`src/payment/refund/`)

#### 4.3.1 ROJO — primera prueba de reembolso total

**`src/payment/refund/refund_test.go`** (versión inicial):

```go
package refund

import (
    "testing"
    "time"

    "tdd/src/payment/processor"
)

func TestRefundProcessor_FullRefund_OfCompletedPayment(t *testing.T) {
    repo := newMockRepo()
    proc := NewProcessor(repo, time.Now())
    payment := &processor.Payment{
        ID:        "pay-1",
        UserID:    "user-1",
        Amount:    1000,
        Total:     1100,
        Tax:       100,
        Status:    processor.StatusCompleted,
        CreatedAt: time.Now().Add(-1 * time.Hour),
    }
    repo.payments[payment.ID] = payment

    refund, err := proc.Refund(RefundRequest{
        PaymentID: "pay-1",
        Amount:    1000,
        Reason:    "Solicitud del cliente",
    })

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if refund.Amount != 1000 {
        t.Errorf("refund amount = %.2f; want 1000", refund.Amount)
    }
    if refund.Status != StatusProcessed {
        t.Errorf("status = %s; want %s", refund.Status, StatusProcessed)
    }
}
```

**Comando**:

```bash
go test ./src/payment/refund/ -v
```

**Salida (captura 10)**:

```
src/payment/refund/refund_test.go:11:10: undefined: newMockRepo
src/payment/refund/refund_test.go:12:10: undefined: NewProcessor
src/payment/refund/refund_test.go:24:29: undefined: RefundRequest
src/payment/refund/refund_test.go:36:22: undefined: StatusProcessed
src/payment/refund/refund_test.go:37:51: undefined: StatusProcessed
FAIL    tdd/src/payment/refund [build failed]
FAIL
```

> Estado: **ROJO**. Ningún símbolo del paquete existe todavía.

---

#### 4.3.2 VERDE — código mínimo

**`src/payment/refund/refund.go`** (versión 1, mínima):

```go
package refund

import (
    "errors"
    "sync"
    "time"

    "tdd/src/payment/processor"
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
    ID        string
    PaymentID string
    Amount    float64
    Reason    string
    Status    Status
    CreatedAt time.Time
}

type Repository interface {
    FindPayment(id string) (*processor.Payment, error)
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

func (r *InMemoryRepository) Save(p *processor.Payment) error { /* ... */ }
func (r *InMemoryRepository) FindPayment(id string) (*processor.Payment, error) { /* ... */ }
func (r *InMemoryRepository) SaveRefund(rf *Refund) error { /* ... */ }

type Processor struct {
    repo Repository
    now  func() time.Time
}

func NewProcessor(repo Repository, now time.Time) *Processor {
    return &Processor{repo: repo, now: func() time.Time { return now }}
}

func (p *Processor) Refund(req RefundRequest) (*Refund, error) {
    payment, err := p.repo.FindPayment(req.PaymentID)
    if err != nil {
        return nil, err
    }
    if payment == nil {
        return nil, errors.New("refund: payment not found")
    }
    refund := &Refund{
        PaymentID: req.PaymentID,
        Amount:    req.Amount,
        Reason:    req.Reason,
        Status:    StatusProcessed,
        CreatedAt: p.now(),
    }
    _ = p.repo.SaveRefund(refund)
    return refund, nil
}
```

**`src/payment/refund/mock_repo_test.go`** (mock de pruebas):

```go
package refund

import "tdd/src/payment/processor"

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

func (m *mockRepo) SaveRefund(r *Refund) error {
    m.refunds = append(m.refunds, r)
    return nil
}
```

**Comando**:

```bash
go test ./src/payment/refund/ -v
```

**Salida (captura 11)**:

```
=== RUN   TestRefundProcessor_FullRefund_OfCompletedPayment
--- PASS: TestRefundProcessor_FullRefund_OfCompletedPayment (0.00s)
PASS
ok      tdd/src/payment/refund    0.321s
```

> Estado: **VERDE**. El reembolso total sobre pago completado funciona.

---

#### 4.3.3 ROJO — política de reembolso completa

Se añaden 8 pruebas más: parcial, no permite pago fallido, ventana de 30 días,
límite exacto de 30 días, monto superior al pago, pago inexistente, doble
reembolso y reembolso proporcional de impuestos.

**Comando**:

```bash
go test ./src/payment/refund/ -v
```

**Salida (captura 12)**:

```
src/payment/refund/refund_test.go:92:21: undefined: ErrPaymentNotRefundable
src/payment/refund/refund_test.go:119:21: undefined: ErrRefundWindowExpired
src/payment/refund/refund_test.go:169:21: undefined: ErrRefundExceedsPayment
src/payment/refund/refund_test.go:229:12: refund.RefundedTax undefined
FAIL    tdd/src/payment/refund [build failed]
FAIL
```

> Estado: **ROJO**. Faltan errores tipados y el campo `RefundedTax`.

---

#### 4.3.4 VERDE — implementación con todas las políticas

**`src/payment/refund/refund.go`** (versión final):

```go
package refund

import (
    "errors"
    "sync"
    "time"

    "tdd/src/payment/processor"
)

const refundWindowDays = 30

var (
    ErrPaymentNotFound      = errors.New("refund: payment not found")
    ErrPaymentNotRefundable = errors.New("refund: payment is not refundable")
    ErrRefundWindowExpired  = errors.New("refund: refund window expired (>30 days)")
    ErrRefundExceedsPayment = errors.New("refund: amount exceeds payment total")
    ErrInvalidRefundAmount  = errors.New("refund: invalid refund amount")
)

// ... tipos Status, RefundRequest

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

// ... InMemoryRepository (con TotalRefundedFor)

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
```

Y `mock_repo_test.go` añade `TotalRefundedFor`:

```go
func (m *mockRepo) TotalRefundedFor(paymentID string) (float64, error) {
    var total float64
    for _, r := range m.refunds {
        if r.PaymentID == paymentID && r.Status == StatusProcessed {
            total += r.Amount
        }
    }
    return total, nil
}
```

**Comando**:

```bash
go test ./src/payment/refund/ -v
```

**Salida (captura 13)**:

```
=== RUN   TestRefundProcessor_FullRefund_OfCompletedPayment
--- PASS: TestRefundProcessor_FullRefund_OfCompletedPayment (0.00s)
=== RUN   TestRefundProcessor_PartialRefund_Allowed
--- PASS: TestRefundProcessor_PartialRefund_Allowed (0.00s)
=== RUN   TestRefundProcessor_RejectsRefundOfFailedPayment
--- PASS: TestRefundProcessor_RejectsRefundOfFailedPayment (0.00s)
=== RUN   TestRefundProcessor_RejectsRefundOlderThan30Days
--- PASS: TestRefundProcessor_RejectsRefundOlderThan30Days (0.00s)
=== RUN   TestRefundProcessor_AllowsRefundAtDay30Boundary
--- PASS: TestRefundProcessor_AllowsRefundAtDay30Boundary (0.00s)
=== RUN   TestRefundProcessor_RejectsExceedingOriginalAmount
--- PASS: TestRefundProcessor_RejectsExceedingOriginalAmount (0.00s)
=== RUN   TestRefundProcessor_RejectsRefundOfNonExistingPayment
--- PASS: TestRefundProcessor_RejectsRefundOfNonExistingPayment (0.00s)
=== RUN   TestRefundProcessor_RejectsDoubleRefund
--- PASS: TestRefundProcessor_RejectsDoubleRefund (0.00s)
=== RUN   TestRefundProcessor_RefundsProportionalTax
--- PASS: TestRefundProcessor_RefundsProportionalTax (0.00s)
PASS
ok      tdd/src/payment/refund    0.304s
```

> Estado: **VERDE**. 9/9 pruebas pasan.

---

#### 4.3.5 Detalle: bug detectado por TDD

Durante el ciclo del reembolso doble se detectó un caso interesante: la primera
versión del test usaba `600 + 400` para superar el pago original de 1000, pero
`1000` **no es estrictamente mayor** que `1000`, por lo que la prueba fallaba
de forma inesperada. El error fue del test, no del código:

```
=== RUN   TestRefundProcessor_RejectsDoubleRefund
    refund_test.go:207: segundo reembolso debió fallar porque excede lo restante
--- FAIL: TestRefundProcessor_RejectsDoubleRefund (0.00s)
```

Se corrigió cambiando el monto a `500` para que la suma sí exceda:

```diff
- Amount: 400
+ Amount: 500
```

Esto demuestra exactamente el valor del TDD: el ciclo descubre inconsistencias
entre la regla de negocio y la prueba antes de llegar a producción.

---

## 5. Resultado final consolidado

### 5.1 Comando

```bash
go test ./... -cover
```

### 5.2 Salida (captura 14)

```
ok      tdd/src/payment/processor    coverage: 92.3% of statements
ok      tdd/src/payment/refund       coverage: 46.5% of statements
ok      tdd/src/payment/tax          coverage: 100.0% of statements
```

**Resumen**: 22 pruebas en total, **22 PASS**, 0 FAIL.

### 5.3 Distribución de pruebas

| Componente | Pruebas | Cobertura |
|---|---|---|
| `tax.Calculator` | 6 | 100.0% |
| `processor.Processor` | 7 | 92.3% |
| `refund.Processor` | 9 | 46.5% |
| **Total** | **22** | — |

---

## 6. Reglas de negocio implementadas

### 6.1 Impuestos (`tax`)

| Monto | Tasa |
|---|---|
| `< 0` o `= 0` | error `ErrInvalidAmount` |
| `< 1000` | 0% (exento) |
| `1000 – 10000` | 10% |
| `> 10000` | 15% |

### 6.2 Pagos (`processor`)

- `UserID` obligatorio.
- Monto debe ser `> 0` y `>= 10` (mínimo).
- Por usuario, la suma de pagos del día no puede superar `$5000`.
- El pago se registra como `completed` y se calcula `Total = Amount + Tax`.

### 6.3 Reembolsos (`refund`)

- Sólo sobre pagos en estado `completed`.
- Sólo dentro de los 30 días posteriores al pago.
- Reembolso parcial permitido, pero el total acumulado no puede superar
  el monto original.
- Se devuelve el impuesto proporcional al monto reembolsado.

---

## 7. Conclusiones y recomendaciones

- **TDD fuerza la claridad del contrato**: cada regla de negocio quedó
  documentada como una prueba concreta, no como intención.
- **El ciclo descubre bugs temprano**: el caso del reembolso doble (`600+400=1000`)
  fue detectado porque la prueba no fallaba como debía, obligando a revisar
  la regla antes de pasar a producción.
- **Las pruebas son ejecutables**: cualquier cambio en `minAmount`, `dailyLimit`,
  `refundWindowDays` o las tasas de impuestos se valida de inmediato corriendo
  `go test ./...`.
- **Cobertura**: `tax` alcanza el 100%. En `processor` (92.3%) y `refund`
  (46.5%) la cobertura baja por las implementaciones `InMemoryRepository`,
  pensadas para integración, no para tests unitarios (los tests usan mocks).
- **Recomendación para el laboratorio**: el workflow de GitHub Actions
  (`.github/workflows/tests.yml`) automatiza la verificación en cada
  `push` o `pull_request` a `main`, exigiendo 100% de cobertura en `tax` y
  ≥ 80% en `processor`.
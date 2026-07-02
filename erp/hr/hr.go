package hr

import (
	"erp/types"
	"time"
)

type HRService struct {
	employees map[string]*types.Employee
	payrolls  map[string]*types.Payroll
}

func NewHRService() *HRService {
	return &HRService{
		employees: make(map[string]*types.Employee),
		payrolls:  make(map[string]*types.Payroll),
	}
}

func (s *HRService) CreateEmployee(name, email, department string, baseSalary, hourlyRate float64) *types.Employee {
	employee := &types.Employee{
		ID:           generateHRID("emp"),
		Name:         name,
		Email:        email,
		Department:   department,
		BaseSalary:   baseSalary,
		HourlyRate:   hourlyRate,
		WorkSchedule: make([]time.Time, 0),
	}
	s.employees[employee.ID] = employee
	return employee
}

func (s *HRService) GetEmployee(id string) (*types.Employee, error) {
	employee, ok := s.employees[id]
	if !ok {
		return nil, types.ErrEmployeeNotFound
	}
	return employee, nil
}

func (s *HRService) RecordAttendance(employeeID string, checkIn time.Time) error {
	employee, err := s.GetEmployee(employeeID)
	if err != nil {
		return err
	}
	employee.WorkSchedule = append(employee.WorkSchedule, checkIn)
	return nil
}

func (s *HRService) ProcessPayroll(employeeID string, extraHours float64) (*types.Payroll, error) {
	employee, err := s.GetEmployee(employeeID)
	if err != nil {
		return nil, err
	}

	extraHoursPay := extraHours * employee.HourlyRate
	total := employee.BaseSalary + extraHoursPay

	payroll := &types.Payroll{
		ID:            generateHRID("pay"),
		EmployeeID:    employee.ID,
		BaseSalary:    employee.BaseSalary,
		ExtraHours:    extraHours,
		ExtraHoursPay: extraHoursPay,
		Total:         total,
		ProcessedAt:   time.Now(),
	}
	s.payrolls[payroll.ID] = payroll
	return payroll, nil
}

func (s *HRService) GetPayroll(id string) (*types.Payroll, error) {
	payroll, ok := s.payrolls[id]
	if !ok {
		return nil, types.ErrLeadNotFound
	}
	return payroll, nil
}

func generateHRID(prefix string) string {
	return prefix + "-" + time.Now().Format("20060102150405.000000")
}

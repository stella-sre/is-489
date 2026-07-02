package hr

import (
	"erp/types"
	"testing"
	"time"
)

func TestCreateEmployee(t *testing.T) {
	svc := NewHRService()

	employee := svc.CreateEmployee("Ana Martinez", "ana@empresa.com", "Ventas", 2500.00, 15.00)
	if employee == nil {
		t.Fatal("se esperaba un empleado creado")
	}
	if employee.Name != "Ana Martinez" {
		t.Errorf("nombre esperado: Ana Martinez, obtenido: %s", employee.Name)
	}
	if employee.BaseSalary != 2500.00 {
		t.Errorf("salario base esperado: 2500.00, obtenido: %f", employee.BaseSalary)
	}
	if employee.Department != "Ventas" {
		t.Errorf("departamento esperado: Ventas, obtenido: %s", employee.Department)
	}
}

func TestGetEmployee(t *testing.T) {
	svc := NewHRService()

	employee := svc.CreateEmployee("Luis Rodriguez", "luis@empresa.com", "RRHH", 3000.00, 20.00)
	retrieved, err := svc.GetEmployee(employee.ID)
	if err != nil {
		t.Fatalf("error obteniendo empleado: %v", err)
	}
	if retrieved.Name != "Luis Rodriguez" {
		t.Errorf("nombre esperado: Luis Rodriguez, obtenido: %s", retrieved.Name)
	}
}

func TestGetEmployeeNotFound(t *testing.T) {
	svc := NewHRService()

	_, err := svc.GetEmployee("no-existe")
	if err != types.ErrEmployeeNotFound {
		t.Fatalf("se esperaba ErrEmployeeNotFound, obtenido: %v", err)
	}
}

func TestRecordAttendance(t *testing.T) {
	svc := NewHRService()

	employee := svc.CreateEmployee("Rosa Castro", "rosa@empresa.com", "Marketing", 2200.00, 12.00)
	checkIn := time.Date(2026, 7, 1, 9, 0, 0, 0, time.UTC)

	err := svc.RecordAttendance(employee.ID, checkIn)
	if err != nil {
		t.Fatalf("error registrando asistencia: %v", err)
	}

	retrieved, _ := svc.GetEmployee(employee.ID)
	if len(retrieved.WorkSchedule) != 1 {
		t.Errorf("se esperaba 1 registro de asistencia, obtenido: %d", len(retrieved.WorkSchedule))
	}
}

func TestProcessPayroll(t *testing.T) {
	svc := NewHRService()

	employee := svc.CreateEmployee("Pedro Sanchez", "pedro@empresa.com", "Ventas", 3000.00, 18.00)
	payroll, err := svc.ProcessPayroll(employee.ID, 10.0)
	if err != nil {
		t.Fatalf("error procesando nomina: %v", err)
	}
	if payroll.BaseSalary != 3000.00 {
		t.Errorf("salario base esperado: 3000.00, obtenido: %f", payroll.BaseSalary)
	}
	if payroll.ExtraHours != 10.0 {
		t.Errorf("horas extra esperadas: 10.0, obtenido: %f", payroll.ExtraHours)
	}
	if payroll.ExtraHoursPay != 180.00 {
		t.Errorf("pago horas extra esperado: 180.00, obtenido: %f", payroll.ExtraHoursPay)
	}
	if payroll.Total != 3180.00 {
		t.Errorf("total esperado: 3180.00, obtenido: %f", payroll.Total)
	}
}

func TestProcessPayrollWithNoExtraHours(t *testing.T) {
	svc := NewHRService()

	employee := svc.CreateEmployee("Elena Ruiz", "elena@empresa.com", "RRHH", 2500.00, 15.00)
	payroll, err := svc.ProcessPayroll(employee.ID, 0)
	if err != nil {
		t.Fatalf("error procesando nomina: %v", err)
	}
	if payroll.Total != 2500.00 {
		t.Errorf("total esperado: 2500.00, obtenido: %f", payroll.Total)
	}
}

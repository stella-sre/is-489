package main

import (
	"strings"
	"testing"
)

func TestConvert1_ContainsLoginAndEmail(t *testing.T) {
	result := convert1(INPUT_DATA)
	if !strings.Contains(result, "peterson ==> peterson@outlook.com") {
		t.Errorf("expected 'peterson ==> peterson@outlook.com' in result")
	}
}

func TestConvert1_AllFourRecords(t *testing.T) {
	result := convert1(INPUT_DATA)
	expected := []string{"peterson", "james", "jackson", "gregory"}
	for _, login := range expected {
		if !strings.Contains(result, login+" ==>") {
			t.Errorf("expected login '%s' in convert1 result", login)
		}
	}
}

func TestConvert1_SkipsHeader(t *testing.T) {
	result := convert1(INPUT_DATA)
	if strings.Contains(result, "Login") {
		t.Error("convert1 should not include the header line")
	}
}

func TestConvert2_ContainsNameAndEmail(t *testing.T) {
	result := convert2(INPUT_DATA)
	if !strings.Contains(result, "Chris Peterson (email: peterson@outlook.com)") {
		t.Errorf("expected 'Chris Peterson (email: peterson@outlook.com)' in result")
	}
}

func TestConvert2_AllFourRecords(t *testing.T) {
	result := convert2(INPUT_DATA)
	expected := []string{"Chris Peterson", "Derek James", "Walter Jackson", "Mike Gregory"}
	for _, name := range expected {
		if !strings.Contains(result, name) {
			t.Errorf("expected name '%s' in convert2 result", name)
		}
	}
}

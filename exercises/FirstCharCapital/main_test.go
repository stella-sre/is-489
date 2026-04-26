package main

import "testing"

func TestTitleCase_NormalSentence(t *testing.T) {
	result := firstCharToTitleCase("hello world from go")
	if result != "Hello World From Go" {
		t.Errorf("expected 'Hello World From Go', got '%s'", result)
	}
}

func TestTitleCase_AllUppercase(t *testing.T) {
	result := firstCharToTitleCase("HELLO WORLD")
	if result != "Hello World" {
		t.Errorf("expected 'Hello World', got '%s'", result)
	}
}

func TestTitleCase_SingleWord(t *testing.T) {
	result := firstCharToTitleCase("GOLANG")
	if result != "Golang" {
		t.Errorf("expected 'Golang', got '%s'", result)
	}
}

func TestTitleCase_MixedCase(t *testing.T) {
	result := firstCharToTitleCase("sItE rElIaBiLiTy EnGiNeEr")
	if result != "Site Reliability Engineer" {
		t.Errorf("expected 'Site Reliability Engineer', got '%s'", result)
	}
}

func TestTitleCase_ExtraSpaces(t *testing.T) {
	result := firstCharToTitleCase("stella   sre")
	if result != "Stella   Sre" {
		t.Errorf("expected 'Stella   Sre', got '%s'", result)
	}
}

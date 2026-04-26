package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestRectangle_TopRowAllStars(t *testing.T) {
	out := captureOutput(func() { drawRectangle(4, 6) })
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if lines[0] != "******" {
		t.Errorf("expected '******', got '%s'", lines[0])
	}
}

func TestRectangle_BottomRowAllStars(t *testing.T) {
	out := captureOutput(func() { drawRectangle(4, 6) })
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	last := lines[len(lines)-1]
	if last != "******" {
		t.Errorf("expected '******', got '%s'", last)
	}
}

func TestRectangle_MiddleRowsEdgesOnly(t *testing.T) {
	out := captureOutput(func() { drawRectangle(4, 6) })
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	for i := 1; i < len(lines)-1; i++ {
		if lines[i][0] != '*' || lines[i][len(lines[i])-1] != '*' {
			t.Errorf("row %d must start and end with '*', got '%s'", i, lines[i])
		}
		if lines[i][1] != ' ' {
			t.Errorf("row %d inner chars must be spaces, got '%s'", i, lines[i])
		}
	}
}

func TestRectangle_SingleCell(t *testing.T) {
	out := captureOutput(func() { drawRectangle(1, 1) })
	if strings.TrimSpace(out) != "*" {
		t.Errorf("expected '*', got '%s'", out)
	}
}

func TestRectangle_CorrectNumberOfRows(t *testing.T) {
	out := captureOutput(func() { drawRectangle(5, 4) })
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(lines) != 5 {
		t.Errorf("expected 5 rows, got %d", len(lines))
	}
}

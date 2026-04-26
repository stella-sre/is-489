package main

import (
	"reflect"
	"testing"
)

func TestFilter_KeepsLongWords(t *testing.T) {
	words := []string{"hi", "hello", "world", "ok", "java"}
	result := filterWordsByLength(4, words)
	expected := []string{"hello", "world", "java"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFilter_NoWordsQualify(t *testing.T) {
	words := []string{"hi", "ok", "no"}
	result := filterWordsByLength(5, words)
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func TestFilter_AllWordsQualify(t *testing.T) {
	words := []string{"hello", "world", "java"}
	result := filterWordsByLength(1, words)
	if !reflect.DeepEqual(result, words) {
		t.Errorf("expected %v, got %v", words, result)
	}
}

func TestFilter_ExactLengthMatch(t *testing.T) {
	words := []string{"go", "java", "rust"}
	result := filterWordsByLength(4, words)
	expected := []string{"java", "rust"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFilter_SingleWordPasses(t *testing.T) {
	words := []string{"engineering"}
	result := filterWordsByLength(5, words)
	if len(result) != 1 || result[0] != "engineering" {
		t.Errorf("expected ['engineering'], got %v", result)
	}
}

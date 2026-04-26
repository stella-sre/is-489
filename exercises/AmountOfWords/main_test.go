package main

import "testing"

func TestAmountOfWords_NormalSentence(t *testing.T) {
	result := getWordsAmount("Hello world, how are you?")
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestAmountOfWords_SingleWord(t *testing.T) {
	result := getWordsAmount("engineer")
	if result != 1 {
		t.Errorf("expected 1, got %d", result)
	}
}

func TestAmountOfWords_ExtraSpaces(t *testing.T) {
	result := getWordsAmount("one   two   three")
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func TestAmountOfWords_WithPunctuation(t *testing.T) {
	result := getWordsAmount("site reliability engineer")
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func TestAmountOfWords_MultipleSpecialChars(t *testing.T) {
	result := getWordsAmount("hello, world! how... are you?")
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

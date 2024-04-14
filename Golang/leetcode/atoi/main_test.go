package main

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	testCases := []struct {
		input    byte
		expected bool
	}{
		{'1', true},
		{'4', true},
		{'0', true},
		{'a', false},
		{'-', false},
		{'+', false},
		{',', false},
		{'.', false},
		{' ', false},
		{'9', true},
		//{'\n', true}, //will fail
	}

	for _, tc := range testCases {
		result := IsDigit(tc.input)
		if result != tc.expected {
			t.Errorf("IsDigit(%d) expected: %t, got: %t", tc.input, tc.expected, result)
		}
	}

}

func TestAtoi(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{" ", 0},
		{"", 0},
		{"a", 0},
		{"aa", 0},
		{"aaa", 0},
		{"aaa1", 0},
		{"2", 2},
		{"   2", 2},
		{"-22", -22},
		{"42", 42},
		{"    42", 42},
		{"2147483647", 2147483647},  // Max int32 value3
		{"99997483648", 2147483647}, // Greater than max int32, should clamp to max int32
		{"-42", -42},
		{"-99997483648", -2147483648}, // Greater than max int32, should clamp to max int32
		{"42      ", 42},
		{"    42      ", 42},
		{"-42", -42},
		{"- 42", 0}, //doublecheck
		{"    -42", -42},
		{"-42      ", -42},
		{"    -42      ", -42},
		{"-4-2", -4},
		{"4-2", 4},
		{"4 - 2", 4},
		{"1+1", 1},
		{"-123", -123},
		{"-123-123", -123},
		{"+123+123", 123},
		{"123-123+123", 123},
		{"123+123", 123},
		{"-2147483648", -2147483648}, // Min int32 value
		{"-2147483649", -2147483648}, // Less than min int32, should clamp to -2147483647
		{"1 2 3", 1},                 // Spaces should be ignored
		{"+42", 42},                  // Leading '+' sign should be allowed
		{"- 42", 0},                  // Invalid leading '-' sign should return 0
		{"-+42", 0},                  // Invalid signs should return 0
		{"words and 987", 0},
		{"words and -987", 0},
		{"4193 with words", 4193},
		{"-4193 with words", -4193},
	}

	for _, tc := range testCases {
		result := Atoi(tc.input)
		if result != tc.expected {
			t.Errorf("Atoi(%q) expected: %d, got: %d", tc.input, tc.expected, result)
		}
	}
}

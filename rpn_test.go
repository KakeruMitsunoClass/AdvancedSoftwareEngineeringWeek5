package main

import "testing"

func TestRpn01(t *testing.T) {
	expected := "42"
	input := "40 2 +"
	result := rpn(input)
	if expected != result {
		t.Errorf("Test01 fail expected: %s, result: %s\n", expected, result)
	}
}

func TestRpn02(t *testing.T) {
	expected := "1"
	input := "3 2 -"
	result := rpn(input)
	if expected != result {
		t.Errorf("Test02 fail expected: %s, result: %s\n", expected, result)
	}
}

func TestRpn03(t *testing.T) {
	expected := "6"
	input := "2 3 *"
	result := rpn(input)
	if expected != result {
		t.Errorf("Test03 fail expected: %s, result: %s\n", expected, result)
	}
}

func TestRpn04(t *testing.T) {
	expected := "2"
	input := "4 2 /"
	result := rpn(input)
	if expected != result {
		t.Errorf("Test04 fail expected: %s, result: %s\n", expected, result)
	}
}

func TestRpn05(t *testing.T) {
	expected := "4"
	input := "2 3 + 3 - 4 * 2 /"
	result := rpn(input)
	if expected != result {
		t.Errorf("Test05 fail expected: %s, result: %s\n", expected, result)
	}
}

func TestRpn06(t *testing.T) {
	expected := "8"
	input := "2 3 3 + +"
	result := rpn(input)
	if expected != result {
		t.Errorf("Test06 fail expected: %s, result: %s\n", expected, result)
	}
}

func TestRpn07(t *testing.T) {
	expected := "-4"
	input := "2 3 3 + -"
	result := rpn(input)
	if expected != result {
		t.Errorf("Test07 fail expected: %s, result: %s\n", expected, result)
	}
}

func TestRpn08(t *testing.T) {
	expected := "11"
	input := "2 3 3 * +"
	result := rpn(input)
	if expected != result {
		t.Errorf("Test08 fail expected: %s, result: %s\n", expected, result)
	}
}

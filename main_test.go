package main_test

import (
	"github.com/huytd/go-fizzbuzz"
	"testing"
)

func TestFizzBuzzReturnFizz(t *testing.T) {
	result := main.FizzBuzz(3)
	expect := "Fizz"

	if result != expect {
		t.Fail()
	}
}

func TestFizzBuzzReturnBuzz(t *testing.T) {
	result := main.FizzBuzz(5)
	expect := "Buzz"

	if result != expect {
		t.Fail()
	}
}

func TestFizzBuzzReturnFizzBuzz(t *testing.T) {
	result := main.FizzBuzz(15)
	expect := "FizzBuzz"

	if result != expect {
		t.Fail()
	}
}

func TestFizzBuzzReturnNumber(t *testing.T) {
	result := main.FizzBuzz(1)
	expect := "1"

	if result != expect {
		t.Fail()
	}
}

func TestFizzBuzzTableDriven(t *testing.T) {
	var testCases = []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{6, "Fizz"},
		{9, "Fizz"},
		{8, "8"},
		{10, "Buzz"},
		{12, "Fizz"},
		{15, "FizzBuzz"},
		{25, "Buzz"},
		{30, "FizzBuzz"},
	}

	for _, test := range testCases {
		result := main.FizzBuzz(test.in)
		expect := test.out

		if result != expect {
			t.Errorf("Test failed at input: %v \n  Actual result: %v \n  Expected: %v", test.in, result, expect)
		}
	}
}

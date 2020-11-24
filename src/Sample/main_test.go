package main

import (
	"strconv"
	"testing"
)

func Test_main(t *testing.T) {
	got := sayHello("Ala")
	expect := "Hello Ala:)"

	if got != expect {
		t.Errorf("Wanted:" + expect + "but got: " + got)
	}
}

func Test_main2(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
		number int
	}{
		{input: "Ala", expect: "Hello Ala:)", number: 1},
		{input: "", expect: "Hello :)", number: 2},
	}

	for _, s := range scenarios {
		got := sayHello(s.input)
		if got != s.expect {
			t.Errorf("Test: " + strconv.Itoa(s.number) + " -> Wanted: " + s.expect + ", but got: " + got)
		}
	}
}

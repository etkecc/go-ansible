package ansible

import (
	"strconv"
	"testing"
)

func TestUniq(t *testing.T) {
	input := []string{"1", "2", "2", "3", "3", "3"}
	expected := []string{"1", "2", "3"}

	actual := Uniq(input)

	if len(expected) != len(actual) {
		t.Error("length is not expected. Expected:", len(expected), "actual:", len(actual))
	}
	for i, e := range expected {
		if e != actual[i] {
			t.Error("actual#"+strconv.Itoa(i), " is invalid. Expected: '"+e+"', actual: '"+actual[i]+"'")
		}
	}
}

func TestMapKeys(t *testing.T) {
	input := map[string]string{"1": "1", "2": "2", "3": "3"}
	expected := []string{"1", "2", "3"}

	actual := MapKeys(input)

	if len(expected) != len(actual) {
		t.Error("length is not expected. Expected:", len(expected), "actual:", len(actual))
	}
	for i, e := range expected {
		if e != actual[i] {
			t.Error("actual#"+strconv.Itoa(i), " is invalid. Expected: '"+e+"', actual: '"+actual[i]+"'")
		}
	}
}

func TestUnquote(t *testing.T) {
	tests := map[string]string{
		"not quoted": "not quoted",
		`"quoted"`:   "quoted",
	}
	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			actual := Unquote(input)

			if expected != actual {
				t.Error("actual is invalid. Expected: '" + expected + "', actual: '" + actual + "'")
			}
		})
	}
}

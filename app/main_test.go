package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	t.Run("base case", func(t *testing.T) {
		input := "echo 'hello'"
		got := parseInput(input)
		want := []string{"echo", "hello"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("command w/o args", func(t *testing.T) {
		input := "date"
		got := parseInput(input)
		want := []string{"date"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("one arg", func(t *testing.T) {
		input := "echo hello"
		got := parseInput(input)
		want := []string{"echo", "hello"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("two args", func(t *testing.T) {
		input := "echo hello world"
		got := parseInput(input)
		want := []string{"echo", "hello", "world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("single quoted arg", func(t *testing.T) {
		input := "echo 'hello world'"
		got := parseInput(input)
		want := []string{"echo", "hello world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("single quoted arg with multiple spaces", func(t *testing.T) {
		input := "echo 'hello   world'"
		got := parseInput(input)
		want := []string{"echo", "hello   world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("complex single quoted args cae", func(t *testing.T) {
		input := "echo 'hello   world' my''super hero"
		got := parseInput(input)
		want := []string{"echo", "hello   world", "mysuper", "hero"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("double quoted arg", func(t *testing.T) {
		input := "echo \"hello test\""
		got := parseInput(input)
		want := []string{"echo", "hello test"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("double quoted arg with multiple spaces", func(t *testing.T) {
		input := "echo \"hello   world\""
		got := parseInput(input)
		want := []string{"echo", "hello   world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("single quote in double quotes", func(t *testing.T) {
		input := "echo \"example\"  \"hello's\"  test\"\"shell"
		got := parseInput(input)
		want := []string{"echo", "example", "hello's", "testshell"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("backslash", func(t *testing.T) {
		input := "echo hello\\ \\ test"
		got := parseInput(input)
		want := []string{"echo", "hello  test"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("backslash in quotes", func(t *testing.T) {
		input := "echo \"hello \\\" world\""
		got := parseInput(input)
		want := []string{"echo", "hello \" world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("escape chars in paths", func(t *testing.T) {
		input := "echo \"/tmp/quz/'f 15'\""
		got := parseInput(input)
		want := []string{"echo", "/tmp/quz/'f 15'"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("escape chars in paths 2", func(t *testing.T) {
		input := "echo \"/tmp/ant/'f  \\43'\""
		got := parseInput(input)
		want := []string{"echo", "/tmp/ant/'f  \\43'"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("input: %v, result not match, got %v, want %v", input, got, want)
		}
	})

	t.Run("parse full input", func(t *testing.T) {
		input := "echo test"
		got := parseInput(input)
		want := []string{"echo", "test"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("input: %v, result not match, got %v, want %v", input, got, want)
		}
	})

	t.Run("parse full input", func(t *testing.T) {
		input := "echo 'test 123'"
		got := parseInput(input)
		want := []string{"echo", "test 123"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("input: %v, result not match, got %v, want %v", input, got, want)
		}
	})

	t.Run("parse full input", func(t *testing.T) {
		input := "'echo ko' 'test 123'"
		got := parseInput(input)
		want := []string{"echo ko", "test 123"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("input: %v, result not match, got %v, want %v", input, got, want)
		}
	})
}

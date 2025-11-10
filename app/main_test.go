package main

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		input := "echo 'hello'"
		got := parseCommand(input)
		want := "echo"
		if got != want {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("command w/o args", func(t *testing.T) {
		input := "date"
		got := parseCommand(input)
		want := "date"
		if got != want {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})
}

func TestParseArgs(t *testing.T) {
	t.Run("one arg", func(t *testing.T) {
		input := "echo hello"
		got := parseArgs(input)
		want := []string{"hello"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("two args", func(t *testing.T) {
		input := "echo hello world"
		got := parseArgs(input)
		want := []string{"hello", "world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("single quoted arg", func(t *testing.T) {
		input := "echo 'hello world'"
		got := parseArgs(input)
		want := []string{"hello world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("single quoted arg with multiple spaces", func(t *testing.T) {
		input := "echo 'hello   world'"
		got := parseArgs(input)
		want := []string{"hello   world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("complex single quoted args cae", func(t *testing.T) {
		input := "echo 'hello   world' my''super hero"
		got := parseArgs(input)
		want := []string{"hello   world", "mysuper", "hero"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("double quoted arg", func(t *testing.T) {
		input := "echo \"hello test\""
		got := parseArgs(input)
		want := []string{"hello test"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("double quoted arg with multiple spaces", func(t *testing.T) {
		input := "echo \"hello   world\""
		got := parseArgs(input)
		want := []string{"hello   world"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("single quote in double quotes", func(t *testing.T) {
		input := "echo \"example\"  \"hello's\"  test\"\"shell"
		got := parseArgs(input)
		want := []string{"example", "hello's", "testshell"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("backslash", func(t *testing.T) {
		input := "echo hello\\ \\ test"
		got := parseArgs(input)
		want := []string{"hello  test"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})

	t.Run("backslash in quotes", func(t *testing.T) {
		input := "echo 'hello\\test'"
		got := parseArgs(input)
		want := []string{"hello\\test"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("result not match, got %v, want %v", got, want)
		}
	})
}

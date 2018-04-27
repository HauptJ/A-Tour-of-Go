package main

import (
  "fmt"
  "testing"
)

/*
DESC: Test driver for tail.go
*/
func TestStringArray(t *testing.T) {
  // test variables normally fed in as CLI args
  file_path := "test.txt"
  num_lines := 1

  // Expected test result
  test_expected := "}\n"

  lines := read_file(&file_path)
  tail := get_tail(num_lines, len(lines), lines)
  tail_string := slice_tail(tail)

  // Check if we get the expected string
  switch {
  case tail_string == test_expected:
    fmt.Println("PASS")
  default:
    t.Error("Expected: ", test_expected ," got: ", tail_string)
  }
}

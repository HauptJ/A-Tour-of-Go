/*
FILE: tail.go
DESC: Prints last n lines in a text file
DATE: 26 APR 18
*/

package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func read_file(file_path *string) []string {
  file_handle, _ := os.Open(*file_path)
  defer file_handle.Close()
  file_scanner := bufio.NewScanner(file_handle)
  var lines []string
  for i := 0; file_scanner.Scan(); i++ {
    lines = append(lines, file_scanner.Text())
  }
  return lines
}

/*
DESC: Get the last n lines in an array
*/
func get_tail(num_lines int, file_len int, lines []string) []string {
  // if num_lines >= file_len -> return all lines
  if num_lines >= file_len {
    tail := lines
    return tail
  } else { // print last num_lines in file
    tail := lines[(len(lines)-num_lines):]
    return tail
  }
}

/*
DESC: Join the last n lines in the tail as a single string
*/
func slice_tail(tail []string) string {
  var line string
  for _, i := range tail {
    line += fmt.Sprintf("%v\n", i)
  }
  return line
}

func main() {
  file_path := os.Args[1]
  num_lines := os.Args[2]
  lines := read_file(&file_path)
  num_lines_int, _ := strconv.Atoi(num_lines)
  tail := get_tail(num_lines_int, len(lines), lines)
  tail_string := slice_tail(tail)
  fmt.Print(tail_string)
}

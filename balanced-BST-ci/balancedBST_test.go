package main

import (
  //"strings"
  "fmt"
  "reflect"
  "testing"
)

type test_data struct {
  // unprocessed input
  values_in, data_in []string
  // processed output
  values_out, data_out []string
}

var tests = []test_data{
  { []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"},
  []string{"delta", "bravo", "golang", "golf", "charlie", "echo", "alpha", "hotel", "foxtrot", "india", "juliett", "lima", "kilo"},
  []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"},
  []string{"alpha", "bravo", "charlie", "delta", "echo", "foxtrot", "golf", "hotel", "india", "juliett", "kilo", "lima"} },
}

func test_bst(t *testing.T) {
  for _, test := range tests {
    fmt.Printf("%v", test)
    v, d := bst(test.values_in, test.data_in)
    print_balanced_BST(&v, &d)
    if (reflect.DeepEqual(v, test.values_out) != true && reflect.DeepEqual(d, test.data_out) != true) {
      t.Error(
      "For values: ", test.values_in, " data: ", test.data_in,
      " Expected", test.values_out, " data: ", test.data_out,
      " Got values: ", v, " data: ", d,
      )
    } else {
      fmt.Println("PASS")
    }
  }
}

/*
FILE: stringer-test.go
DESC: Test driver for stringer.go
DATE: 22 APR 18
*/


package main

import (
	"fmt"
	"strings"
	"testing"
)


/*
DESC: Test driver for stringer.go
*/
func TestStringArray(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	
	// Get the hostname IP combo as an [2]array of strings from format_map()
	hostnames := format_map(hosts)
	fmt.Printf("%v\n", hostnames)
	
	// Make a new slice referencing hostnames [2]array
	host_map := hostnames[:]
	fmt.Printf("%v\n", host_map)
	
	// Join strings in host_map slice
	var test_string string = strings.Join(host_map, " , ")
	fmt.Printf("%v\n", test_string)
	
	// Check if we get one of the two expected strings
	//if test_string != "loopback: 127.0.0.1 , googleDNS: 8.8.8.8"  "googleDNS: 8.8.8.8 , loopback: 127.0.0.1" {
	switch {
	case test_string == "loopback: 127.0.0.1 , googleDNS: 8.8.8.8":
		fmt.Println("PASS, case 1")
	case test_string == "googleDNS: 8.8.8.8 , loopback: 127.0.0.1":
		fmt.Println("PASS, case 2")
	default: 
		t.Error("Expected: loopback: 127.0.0.1 , googleDNS: 8.8.8.8, got ", test_string)
	}
}
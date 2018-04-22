/*
FILE: stringer.go
DESC: A simple Go program to demonstrate CI testing
DATE: 22 APR 18
*/


package main

import (
	"fmt"
)


type IPAddr [4]byte


/*
DESC: STRINGER, joins IP octets into a single string
IN: IPAddr from hosts map
OUT: ip_string string
*/
func (octet IPAddr) String() string {
	ip_string := fmt.Sprintf("%d.%d.%d.%d", octet[0], octet[1], octet[2], octet[3])
	return ip_string
}


/*
DESC: FUNCTION, formats map key and values into a string
IN: hosts map[string]IPAddr
OUT: hostnames [2]string
*/
func format_map(hosts map[string]IPAddr) [2]string {
	var hostnames [2]string
	i := 0
	for name, ip := range hosts {
		hostnames[i] = fmt.Sprintf("%v: %v", name, ip)
		i++
	}
	return hostnames
}
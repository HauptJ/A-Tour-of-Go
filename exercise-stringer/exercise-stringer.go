package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (octet IPAddr) String() string {
	ip_string := fmt.Sprintf("%d.%d.%d.%d", octet[0], octet[1], octet[2], octet[3])
	return ip_string
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
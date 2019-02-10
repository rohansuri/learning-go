package interfaces

import (
	"fmt"
)

type IPAddr [4]byte

/*
	Exercise: Add a "String() string" method to IPAddr.

	We have to add each byte value to the string.

	Note there's a difference between interpreting the byte value itself as a string
	vs
	simply taking the byte value as a string i.e. putting quotes around the value and making it
	a string (not interpreting)
	Our task here is to make it a string.

	IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
*/
func (ipAddr IPAddr) String() (s string) {
	s = ""
	for i, b := range ipAddr {
		// s += string(b)  // This would add 127 as the first byte in the string

		s += fmt.Sprintf("%v", b)
		// this is the valueOf we need
		// this would add 49, 50, 55 as per the ASCII table to form a string of "127"
		// fmt.Printf("%v\n", []byte(fmt.Sprintf("%v", b)))

		if i+1 != len(ipAddr) {
			s += "."
		}
	}

	return
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":   {127, 0, 0, 1},
		"googleDNS":  {8, 8, 8, 8},
		"A in ASCII": {65, 65, 65, 65},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

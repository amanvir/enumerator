package subdomain

import (
	"fmt"
	"net"
)

func Bruteforce(domain string) {
	subdomains := []string{"blog", "ir", "support", "does-not-exist"}

	for _, name := range subdomains {
		host := name + "." + domain

		ips, err := net.LookupIP(host)

		if err != nil {
			fmt.Println(host)
			fmt.Println(err)
		} else {
			fmt.Println(host)
			fmt.Println(ips)
		}
	}
}

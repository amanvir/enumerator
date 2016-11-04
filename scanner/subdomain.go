package subdomain

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"net"
)

func Bruteforce(domain string) {
	subdomains := []string{"blog", "ir", "support"}

	//from golang docs
	netaddr := func(ip net.IP) (net.Addr, error) {
		switch c.LocalAddr().(type) {
		case *net.UDPAddr:
			return &net.UDPAddr{IP: ip}, nil
		case *net.IPAddr:
			return &net.IPAddr{IP: ip}, nil
		default:
			return nil, errors.New("neither UDPAddr nor IPAddr")
		}
	}

	for _, name := range subdomains {
		host := name + "." + domain

		ips, err := net.LookupIP(host)

		if err != nil {
			fmt.Println(err)
		}

		for _, ip := range ips {
			switch protocol {
			case iana.ProtocolICMP:
				if ip.To4() != nil {
					fmt.Println("found ip " + netaddr(ip) + "for " + host)
				}
			case iana.ProtocolIPv6ICMP:
				if ip.To16() != nil && ip.To4() == nil {
					fmt.Println("found ip" + netaddr(ip) + "for" + host)
				}
			}
		}

	}
}

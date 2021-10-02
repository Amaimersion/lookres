package info

import (
	"net"
)

type Domain struct {
	IPv4 string
	IPv6 string
	IP []string
	CNAME string
	TXT []string
	MX []string
	NS []string
}

func HandleDomain(domain string) (Domain, error) {
	info := Domain{}

	if ip, err := net.ResolveIPAddr("ip4", domain); err == nil {
		info.IPv4 = ip.String()
	}

	if ip, err := net.ResolveIPAddr("ip6", domain); err == nil {
		info.IPv6 = ip.String()
	}

	// `LookupIP()` is preferred over `LookupHost()`,
	// but there is no difference actually
	if ip, err := net.LookupIP(domain); err == nil {
		info.IP = make([]string, len(ip))

		for i, value := range ip {
			info.IP[i] = value.String()
		}
	}

	if cname, err := net.LookupCNAME(domain); err == nil {
		info.CNAME = cname
	}

	if txt, err := net.LookupTXT(domain); err == nil {
		info.TXT = txt
	}

	if mx, err := net.LookupMX(domain); err == nil {
		info.MX = make([]string, len(mx))

		for i, value := range mx {
			info.MX[i] = value.Host
		}
	}

	if ns, err := net.LookupNS(domain); err == nil {
		info.NS = make([]string, len(ns))

		for i, value := range ns {
			info.NS[i] = value.Host
		}
	}

	return info, nil
}

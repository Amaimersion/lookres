package sio

import (
	"strings"
	"lookres/info"
)

func IPInfoString(ip info.IP) string {
	s := ""

	if len(ip.Str) > 0 {
		s +=
			"IP: " +
			ip.Str +
			lineSeparator
	}

	if len(ip.Domains) > 0 {
		s +=
			"Domains: " +
			strings.Join(ip.Domains, arraySeparator) +
			lineSeparator
	}

	if ip.IsLoopback {
		s +=
			"Loopback: " +
			boolToStr(ip.IsLoopback) +
			lineSeparator
	}

	if ip.IsMulticast {
		s +=
			"Multicast: " +
			boolToStr(ip.IsMulticast) +
			lineSeparator
	}

	if ip.IsPrivate {
		s +=
			"Private: " +
			boolToStr(ip.IsPrivate) +
			lineSeparator
	}

	if ip.IsUnspecified {
		s +=
			"Unspecified: " +
			boolToStr(ip.IsUnspecified) +
			lineSeparator
	}

	if ip.IsGlobalUnicast {
		s +=
			"Global unicast: " +
			boolToStr(ip.IsGlobalUnicast) +
			lineSeparator
	}

	if ip.IsInterfaceLocalMulticast {
		s +=
			"Interface local multicast: " +
			boolToStr(ip.IsInterfaceLocalMulticast) +
			lineSeparator
	}

	if ip.IsLinkLocalMulticast {
		s +=
			"Link local multicast: " +
			boolToStr(ip.IsLinkLocalMulticast) +
			lineSeparator
	}

	if ip.IsLinkLocalUnicast {
		s +=
			"Link local unicast: " +
			boolToStr(ip.IsLinkLocalUnicast) +
			lineSeparator
	}

	if len(s) > 0 {
		s = s[:len(s) - len(lineSeparator)]
	} else {
		s = "No info about IP"
	}

	return s
}

func DomainInfoString(domain info.Domain) string {
	s := ""

	if len(domain.IPv4) > 0 {
		s +=
			"IPv4: " +
			domain.IPv4 +
			lineSeparator
	}

	if len(domain.IPv6) > 0 {
		s +=
			"IPv6: " +
			domain.IPv6 +
			lineSeparator
	}

	if len(domain.IP) > 0 {
		s +=
			"IP addresses: " +
			strings.Join(domain.IP, arraySeparator) +
			lineSeparator
	}

	if len(domain.CNAME) > 0 {
		s +=
			"CNAME: " +
			domain.CNAME +
			lineSeparator
	}

	if len(domain.MX) > 0 {
		s +=
			"MX: " +
			strings.Join(domain.MX, arraySeparator) +
			lineSeparator
	}

	if len(domain.NS) > 0 {
		s +=
			"NS: " +
			strings.Join(domain.NS, arraySeparator) +
			lineSeparator
	}

	if len(domain.TXT) > 0 {
		s +=
			"TXT: " +
			strings.Join(domain.TXT, arraySeparator) +
			lineSeparator
	}

	if len(s) > 0 {
		s = s[:len(s) - len(lineSeparator)]
	} else {
		s = "No info about domain"
	}

	return s
}

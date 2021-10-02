package info

import (
	"net"
	"strings"
	"errors"
)

type IP struct {
	Str string
	Version string
	Domains []string
	IsGlobalUnicast bool
	IsInterfaceLocalMulticast bool
	IsLinkLocalMulticast bool
	IsLinkLocalUnicast bool
	IsLoopback bool
	IsMulticast bool
	IsPrivate bool
	IsUnspecified bool
}

func HandleIP(s string) (IP, error) {
	ip := net.ParseIP(s)
	info := IP{}

	if ip == nil {
		return info, errors.New("invalid value")
	}

	info.Str = ip.String()
	info.IsGlobalUnicast = ip.IsGlobalUnicast()
	info.IsInterfaceLocalMulticast = ip.IsInterfaceLocalMulticast()
	info.IsLinkLocalMulticast = ip.IsLinkLocalMulticast()
	info.IsLinkLocalUnicast = ip.IsLinkLocalUnicast()
	info.IsLoopback = ip.IsLoopback()
	info.IsMulticast = ip.IsMulticast()
	info.IsPrivate = ip.IsPrivate()
	info.IsUnspecified = ip.IsUnspecified()
	info.Version = "4"

	if strings.Contains(info.Str, ":") {
		info.Version = "6"
	}

	if names, err := net.LookupAddr(info.Str); err == nil {
		info.Domains = names
	}

	return info, nil
}

func IsIP(s string) bool {
	ip := net.ParseIP(s)
	isIp := ip != nil

	return isIp
}

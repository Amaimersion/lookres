package main

import (
	"fmt"
	"os"
	"lookres/info"
	"lookres/sio"
)

func main() {
	input, err := sio.Read()

	if err != nil {
		fmt.Println("input error:", err)
		os.Exit(1)
	}

	isIp := info.IsIP(input.Arg1)

	var ipInfo info.IP
	var domainInfo info.Domain

	if isIp {
		ipInfo, _ = info.HandleIP(input.Arg1)

		if len(ipInfo.Domains) > 0 {
			domainInfo, _ = info.HandleDomain(ipInfo.Domains[0])
		}
	} else {
		domainInfo, _ = info.HandleDomain(input.Arg1)

		if len(domainInfo.IPv4) > 0 {
			ipInfo, _ = info.HandleIP(domainInfo.IPv4)
		} else if len(domainInfo.IPv6) > 0 {
			ipInfo, _ = info.HandleIP(domainInfo.IPv6)
		}
	}

	var output sio.Output = make([]string, 2)

	output[0] = sio.IPInfoString(ipInfo)
	output[1] = sio.DomainInfoString(domainInfo)

	if err := sio.Write(output); err != nil {
		fmt.Println("output error:", err)
		os.Exit(1)
	}
}

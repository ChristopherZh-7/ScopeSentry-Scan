// utils-------------------------------------
// @file      : dns.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2024/9/17 17:25
// -------------------------------------------

package utils

import (
	"fmt"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/types"
	miekgdns "github.com/miekg/dns"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/retryabledns"
	"math"
)

type DnsTools struct {
	Clinet *dnsx.DNSX
}

var DNS *DnsTools

var DefaultResolvers = []string{
	"udp:1.1.1.1:53",         // Cloudflare
	"udp:1.0.0.1:53",         // Cloudflare
	"udp:8.8.8.8:53",         // Google
	"udp:8.8.4.4:53",         // Google
	"udp:9.9.9.9:53",         // Quad9
	"udp:149.112.112.112:53", // Quad9
	"udp:208.67.222.222:53",  // Open DNS
	"udp:208.67.220.220:53",  // Open DNS
}

func InitializeDnsTools() {
	var DefaultOptions = dnsx.Options{
		BaseResolvers:     DefaultResolvers,
		MaxRetries:        3,
		QuestionTypes:     []uint16{miekgdns.TypeA},
		TraceMaxRecursion: math.MaxUint16,
		Hostsfile:         true,
	}
	dnsClient, err := dnsx.New(DefaultOptions)
	if err != nil {
		gologger.Error().Msg(fmt.Sprintf("DNS initialize error: %v", err))
		return
	}
	DNS = &DnsTools{
		Clinet: dnsClient,
	}
}

func (d *DnsTools) QueryOne(hostname string) *retryabledns.DNSData {
	rawResp, err := d.Clinet.QueryOne(hostname)
	if err != nil {
		gologger.Error().Msg(fmt.Sprintf("Dns QueryOne error: %v", err))
		return &retryabledns.DNSData{}
	}
	return rawResp
}

func (d *DnsTools) DNSdataToSubdomainResult(dnsData *retryabledns.DNSData) types.SubdomainResult {
	var recordType string
	switch {
	case len(dnsData.A) > 0:
		recordType = "A"
	case len(dnsData.AAAA) > 0:
		recordType = "AAAA"
	case len(dnsData.CNAME) > 0:
		recordType = "CNAME"
	case len(dnsData.MX) > 0:
		recordType = "MX"
	case len(dnsData.NS) > 0:
		recordType = "NS"
	case len(dnsData.TXT) > 0:
		recordType = "TXT"
	default:
		recordType = "UNKNOWN"
	}
	if recordType == "UNKNOWN" {
		return types.SubdomainResult{}
	}
	var value []string
	value = append(value, dnsData.CNAME...)
	value = append(value, dnsData.MX...)
	value = append(value, dnsData.PTR...)
	value = append(value, dnsData.NS...)
	value = append(value, dnsData.TXT...)
	value = append(value, dnsData.SRV...)
	value = append(value, dnsData.CAA...)
	value = append(value, dnsData.AllRecords...)
	var ip []string
	ip = append(ip, dnsData.A...)
	ip = append(ip, dnsData.AAAA...)
	return types.SubdomainResult{
		Host:  dnsData.Host,
		Type:  recordType,
		Value: value,
		IP:    ip,
	}
}

func (d *DnsTools) KsubdomainVerify(target []string) []types.SubdomainResult {
	//randomString := Tools.GenerateRandomString(6)
	//if len(target) == 0 {
	//	return []types.SubdomainResult{}
	//}
	//filename := util.CalculateMD5(target[0] + randomString)
	//targetPath := filepath.Join(system.KsubdomainPath, "target", filename)
	//resultPath := filepath.Join(system.KsubdomainPath, "result", filename)
	//defer Tools.DeleteFile(targetPath)
	//defer Tools.DeleteFile(resultPath)
	//
	//SubdomainWriteTarget(targetPath, target)
	//args := []string{"v", "-f", targetPath, "-o", resultPath}
	//cmd := exec.Command(system.KsubdomainExecPath, args...)
	//system.SlogInfoLocal(fmt.Sprintf("%v", cmd))
	//output, err := cmd.CombinedOutput()
	//if err != nil {
	//	system.SlogError(fmt.Sprintf("ksubdomain verify 执行命令时出错：%s %s %v\n", err, output, cmd))
	//	return []types.SubdomainResult{}
	//}
	//result := GetSubdomainResult(resultPath)
	//if len(result) == 0 {
	//	system.SlogInfo(fmt.Sprintf("verify target[0] %v get dns result 0", target[0]))
	//	return []types.SubdomainResult{}
	//}
	//return result
	return nil
}
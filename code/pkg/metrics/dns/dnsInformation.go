package dns

/**
 *
 * @Author: zhangxiaohu
 * @File: dnsInformation.go
 * @Version: 1.0.0
 * @Time: 2020/1/14
 */

type DnsEntry struct {
	server string
	address string
}

//func getDNSInformation() DnsEntry {
//	addr, _ := net.InterfaceAddrs()
//	for _,ans := range addr {
//
//	}
//	return DnsEntry{server:net.LookupHost("127.0.0.1")}
//}
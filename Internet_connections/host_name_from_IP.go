package Internet_connections

import (
	"net"
)

//func getIPAdress(r *http.Request) string {
//	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
//		addresses := strings.Split(r.Header.Get(h), ",")
//		// march from right to left until we get a public address
//		// that will be the address right before our proxy.
//		for i := len(addresses) -1 ; i >= 0; i-- {
//			ip := strings.TrimSpace(addresses[i])
//			// header can contain spaces too, strip those out.
//			realIP := net.ParseIP(ip)
//			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
//				// bad address, go to next
//				continue
//			}
//			return ip
//		}
//	}
//	return ""
//}

func getHostNameByIP(address string) (string, error) {
	nets, err := net.LookupAddr(address)

	return nets[0][:len(nets[0])-1], err
}
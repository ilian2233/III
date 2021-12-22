package Internet_connections

import (
	"net"
)

func getHostNameByIP(address string) (string, error) {
	nets, err := net.LookupAddr(address)

	return nets[0][:len(nets[0])-1], err
}

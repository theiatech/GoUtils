package ipUtil

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
)

func IPv4ToLong(ipv4 string) uint32 {
	var long uint32
	if nil == binary.Read(bytes.NewBuffer(net.ParseIP(ipv4).To4()), binary.BigEndian, &long) {
		return long
	}
	return 0
}
func LongToIPv4(ipInt int64) string {
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

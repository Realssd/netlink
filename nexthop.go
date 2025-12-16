package netlink

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type NexthopGroupItem struct {
	ID        uint32
	Weight    uint8
	WeighHigh uint8
	dummy     uint16
}

type Nexthop struct {
	ID        uint32
	Blackhole bool
	Group     []*NexthopGroupItem
	OIF       uint32
	Gateway   net.IP
	Protocol  RouteProtocol
}

func group2String(grp []*NexthopGroupItem) string {
	if grp == nil {
		return "nil"
	}

	builder := new(strings.Builder)
	for i, item := range grp {
		builder.WriteString(strconv.FormatUint(uint64(item.ID), 10))
		builder.WriteByte(',')
		builder.WriteString(strconv.FormatUint(uint64(uint16(item.WeighHigh)<<8|uint16(item.Weight)), 10))
		if i < len(grp)-1 {
			builder.WriteByte('/')
		}
	}
	return builder.String()
}

func (h *Nexthop) String() string {
	elems := []string{
		"ID: " + strconv.FormatUint(uint64(h.ID), 10),
		"Blackhole: " + strconv.FormatBool(h.Blackhole),
		"Group:" + group2String(h.Group),
		"OIF: " + strconv.FormatUint(uint64(h.OIF), 10),
		"Gateway: " + h.Gateway.String(),
		"Protocol: " + h.Protocol.String(),
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, " "))
}

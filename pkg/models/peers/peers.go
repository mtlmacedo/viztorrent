package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"

	"github.com/mtlmacedo/viztorrent/pkg/utils/constants"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

type PeersBin struct {
	Content []byte
	Lenth   int
}

func Unmarshal(peersBin PeersBin) ([]Peer, error) {
	peerCount := peersBin.Lenth / constants.PeerSize
	err := ValidatePeers(peersBin)
	peers := make([]Peer, peerCount)

	for i := 0; i < peerCount; i++ {
		offset := i * constants.PeerSize
		peers[i].IP = net.IP(peersBin.Content[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16([]byte(peersBin.Content[offset+4 : offset+6]))
	}
	return peers, err
}

func ValidatePeers(peersBin PeersBin) error {
	var err error = nil
	if peersBin.Lenth%constants.PeerSize != 0 {
		err = fmt.Errorf(constants.InvalidPeerSize)
	}
	return err
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}

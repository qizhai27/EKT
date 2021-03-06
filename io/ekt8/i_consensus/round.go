package i_consensus

import (
	"encoding/json"
	"fmt"

	"github.com/EducationEKT/EKT/io/ekt8/p2p"
	"github.com/EducationEKT/EKT/io/ekt8/util"
)

type Round struct {
	CurrentIndex int // default -1
	Peers        []p2p.Peer
	Random       int
}

func NextRound(round *Round, CurrentHash []byte) *Round {
	if round.CurrentIndex == len(round.Peers)-1 {
		bytes := CurrentHash[22:]
		Random := util.BytesToInt(bytes)
		round = &Round{
			CurrentIndex: -1,
			Peers:        round.Peers,
			Random:       Random,
		}
	} else {
		round.CurrentIndex++
	}
	return round
}

func (round Round) Len() int {
	return len(round.Peers)
}

func (round Round) Swap(i, j int) {
	round.Peers[i], round.Peers[j] = round.Peers[j], round.Peers[i]
}

func (round Round) Less(i, j int) bool {
	return round.Random%(i+j)%2 == 1
}

func (round Round) String() string {
	peers, _ := json.Marshal(round.Peers)
	return fmt.Sprintf(`{"peers": %s, "random": %d}`, string(peers), round.Random)
}

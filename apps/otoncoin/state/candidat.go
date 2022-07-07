package state

import (
	"encoding/gob"
	"time"

	"github.com/tendermint/tendermint/crypto"
)

const (
	CoinPower    = "Power"
	CoinDepo     = "Deposit"
	CoinContract = "ContractNode"
)

func init() {
	gob.Register(&Candidat{})
}

type Candidat struct {
	PubKey crypto.PubKey  `json:"pub_key"`
	Holder crypto.Address `json:"holder"`

	Contract Coins `json:"contract"`
	Delta    Coins `json:"delta"`
	Power    int64 `json:"power"`

	TimePay time.Time `json:"timepay"`
	Priz    Coins     `json:"priz"`
	Step    int64     `json:"step"`
	Next    int64     `json:"next"`
}

func CandidatKey(addr []byte) []byte {
	return append([]byte("base/cnd/"), addr...)
}

type CandidatProcessor struct {
	PullFee crypto.Address `json:"pull"`
	Ouner   crypto.Address `json:"ouner"`
	Gift    Coins          `json:"gift"`

	Minter crypto.Address `json:"minter"`

	ValMap map[string]Candidat `json:"all"`
}

func (cp *CandidatProcessor) NextPass(next int64) {

}

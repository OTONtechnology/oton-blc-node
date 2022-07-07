package tx

import "github.com/tendermint/tendermint/abci/apps/otoncoin/state"

type Coins []Coin



func (cs Coins) ToStateCoins() state.Coins {
 stateCoins := state.Coins{}
 for _, c := range cs {
  stateCoins = append(stateCoins, state.Coin{Name: c.Name, Amount: c.Amount})
 }

 return stateCoins
}


func FromStateCoins(sc state.Coins) Coins {
 txCoins := Coins{}
 for _, c := range sc {
  txCoins = append(txCoins, Coin{Name: c.Name, Amount: c.Amount})
 }
 return txCoins
}


type CoinsPr []CoinPr

func FromStateCoinsPr(sc state.Coins) CoinsPr {
 txCoins := CoinsPr{}
 for _, c := range sc {
  txCoins = append(txCoins, CoinPr{Name: c.Name, Amount: c.Amount, Expired: c.Expired})
 }
 return txCoins
}

package state

import (
 "encoding/gob"
 "fmt"
 "reflect"
 "strings"

 "github.com/tendermint/tendermint/crypto"
 "github.com/tendermint/tendermint/libs/json"
)

func init() {

 gob.Register(&LeaderMarketing{})
 gob.Register(&RankAssigning{})
 gob.Register(&PoolMarketing{})
 gob.Register(&AMC{})


 gob.Register(&CareerMarketing{})
 gob.Register(&LineMarketing{})
}

const (


 AMCCoinRank = "Rank"
 AMCCoinPacket = "Packet"
 AMCCoinCareerPoint = "CareerPoint"
 AMCCoinEasyBizziPoint = "EasyBizziPoint"
 AMCCoinTarifBbon = "TarifBitBone"
)

type PaymentMaker interface {
 MakePayment(refChain RefChain, delta Coins) (Coins, RefChain)
 MakeRank(refChain RefChain, refRank RefRank, tarif Coins) (RefChain, RefChain)

}


type AMC struct {




 Address crypto.Address `json:"address"`
 Name string `json:"name"`
 Creator crypto.Address `json:"master"`

 Balance Coins `json:"balance"`



 PaymentMakers []PaymentMaker `json:"pay_maker"`
}

type AMCParam struct {
 Key string
 Value int64
}

func NewAMC(adr crypto.Address, name string, creator crypto.Address, balance Coins, params []AMCParam) (*AMC, error) {
 var line *LineMarketing
 var career *CareerMarketing
 var leader *LeaderMarketing
 var pool *PoolMarketing

 var pm []PaymentMaker



 rank := &RankAssigning{}


 for _, p := range params {
  switch {
  case strings.HasPrefix(p.Key, "line."):
   if line == nil {
    line = &LineMarketing{}
   }
   if err := line.SetParam(p.Key[5:], p.Value); err != nil {
    return nil, fmt.Errorf("set line param: %v", err)
   }
  case strings.HasPrefix(p.Key, "career."):
   if career == nil {
    career = &CareerMarketing{}
   }
   if err := career.SetParam(p.Key[7:], p.Value); err != nil {
    return nil, fmt.Errorf("set career param: %v", err)
   }
  case strings.HasPrefix(p.Key, "leader."):
   if leader == nil {
    leader = &LeaderMarketing{}
   }
   if err := leader.SetParam(p.Key[7:], p.Value); err != nil {
    return nil, fmt.Errorf("set leader param: %v", err)
   }
  case strings.HasPrefix(p.Key, "pool."):
   if pool == nil {
    pool = &PoolMarketing{}
   }
   if err := pool.SetParam(p.Key[5:], p.Value); err != nil {
    return nil, fmt.Errorf("set pool param: %v", err)
   }
  case strings.HasPrefix(p.Key, "rank."):
   if err := rank.SetParam(p.Key[5:], p.Value); err != nil {
    return nil, fmt.Errorf("set rank param: %v", err)
   }
  default:
   return nil, fmt.Errorf("unknown param '%s'", p.Key)
  }
 }


 if line != nil {
  pm = append(pm, line)


  if err := line.ValidateLevel(); err != nil {
   return nil, err
  }
 }

 return &AMC{
  Address: adr,
  Name: name,
  Creator: creator,
  Balance: balance,
  PaymentMakers: pm,
 }, nil

}




func (amc *AMC) MakePayment(refChain RefChain, delta Coins) (RefChain, error) {
 totalPayments := Coins{}
 total := Coins{}
 event := RefChain{}



 for _, pm := range amc.PaymentMakers {
  if pm != nil && !reflect.ValueOf(pm).IsNil() {
   total, event = pm.MakePayment(refChain, delta)
   totalPayments = totalPayments.Plus(total)
  }
 }


 if !delta.IsGTE(totalPayments) {

  diff := totalPayments.Minus(delta)
  if !amc.Balance.IsGTE(diff) {
   return event, fmt.Errorf("insufficient balance")
  }

  amc.Balance = amc.Balance.Minus(diff)
 } else {
  diff := delta.Minus(totalPayments)
  amc.Balance = amc.Balance.Plus(diff)
 }
 return event, nil
}


func (amc *AMC) MakeRank(refChain RefChain, refRank RefRank, value Coins) (RefChain, RefChain, error) {



 rf := RefChain{}
 rl := RefChain{}


 for _, pm := range amc.PaymentMakers {
  if pm != nil && !reflect.ValueOf(pm).IsNil() {
   rf, rl = pm.MakeRank(refChain, refRank, value)
  }
 }
 return rf, rl, nil
}


func (amc *AMC) Validate(value Coins, delta Coins) (Coins, error) {
 return nil, nil
}

func (amc *AMC) String() string {
 if amc == nil {
  return "{}"
 }
 bz, _ := json.Marshal(amc.Balance)
 return fmt.Sprintf("{\"address\":\"%v\", \"name\":\"%v\",\"master\":\"%v\",\"balance\": %s }", amc.Address, amc.Name, amc.Creator, string(bz))
}

func AMCKey(address crypto.Address) []byte {
 return append([]byte("base/m/"), address...)
}

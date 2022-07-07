package state

import (
 "fmt"
 "strconv"
 "strings"
)

type CareerMarketing struct {
 Packets map[int64]int64




 Ranks map[int64]int64
 maxRank int64
}

func (cm *CareerMarketing) SetParam(key string, value int64) error {
 p := strings.Split(key, ".")


 if len(p) != 2 {
  return fmt.Errorf("key must have 2 parts")
 }
 switch p[0] {
 case "packet":
  packID, err := strconv.Atoi(p[1])
  if err != nil {
   return err
  }
  if cm.Packets == nil {
   cm.Packets = map[int64]int64{}
  }
  cm.Packets[int64(packID)] = value
 case "rank":
  r, err := strconv.Atoi(p[1])
  if err != nil {
   return err
  }
  rank := int64(r)
  if cm.Ranks == nil {
   cm.Ranks = map[int64]int64{}
  }
  cm.Ranks[rank] = value
  if cm.maxRank < rank {
   cm.maxRank = rank
  }
 default:
  return fmt.Errorf("unknown key part '%s'", p[0])
 }
 return nil
}

func (cm *CareerMarketing) MakePayment(refChain RefChain, delta Coins) (Coins, RefChain) {
 totalPayments := Coins{}
 var maxRank, deductible int64
 for i, acc := range refChain {
  if maxRank >= cm.maxRank {
   break
  }

  if i == 0 {
   continue
  }

  rank := acc.Balance.Coin(AMCCoinRank)
  if rank == nil {
   continue
  }
  paymentRank := rank.Amount

  if cm.Packets != nil {


   pack := acc.Balance.Coin(AMCCoinPacket)
   if pack == nil {
    continue
   }
   if packRank := cm.Packets[pack.Amount]; packRank < paymentRank {
    paymentRank = packRank
   }
  }
  if paymentRank <= maxRank {
   continue
  }

  base := cm.Ranks[rank.Amount]
  percent := base - deductible
  if percent > 0 {
   payment := delta.Percent(percent)
   acc.Balance = acc.Balance.Plus(delta.Percent(percent))
   totalPayments = totalPayments.Plus(payment)
  }
  maxRank = paymentRank
  deductible = base
 }
 return totalPayments, nil
}


func (cm *CareerMarketing) MakeRank(refChain RefChain, refRank RefRank, tarif Coins) (RefChain, RefChain) {
 return nil, nil
}

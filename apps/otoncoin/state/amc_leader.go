package state

import (
 "fmt"
 "strconv"
 "strings"
)

type LeaderMarketing struct {

 Levels int




 Ranks map[int64]int64
}


func (lm *LeaderMarketing) SetParam(key string, value int64) error {
 switch {
 case strings.HasPrefix(key, "level"):
  lm.Levels = int(value)
 case strings.HasPrefix(key, "rank."):
  rank, err := strconv.Atoi(key[5:])
  if err != nil {
   return err
  }
  if lm.Ranks == nil {
   lm.Ranks = map[int64]int64{}
  }
  lm.Ranks[int64(rank)] = value
 default:
  return fmt.Errorf("unknown key '%s'", key)
 }
 return nil
}



func (lm *LeaderMarketing) MakePayment(refChain RefChain, delta Coins) Coins {
 totalPayments := Coins{}
 if lm == nil {
  return totalPayments
 }
 lvl := 0
 for i, acc := range refChain {
  if lvl > lm.Levels {
   break
  }


  if i == 0 {
   lvl++
   continue
  }

  rank := acc.Balance.Coin(AMCCoinRank)
  if rank == nil {
   lvl++
   continue
  }

  percent := lm.Ranks[rank.Amount]
  if percent <= 0 {
   lvl++
   continue
  }
  payment := delta.Percent(percent)
  acc.Balance = acc.Balance.Plus(payment)
  totalPayments = totalPayments.Plus(payment)
  lvl++
 }
 return totalPayments
}

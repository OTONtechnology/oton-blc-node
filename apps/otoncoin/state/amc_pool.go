package state

import (
 "fmt"
 "strconv"
 "strings"
)

type Pool struct {
 balance Coins
 Percent int64
 Shares map[int64]int64
 Career int64
}

type PoolMarketing struct {
 Pools []*Pool
}

func (pm *PoolMarketing) SetParam(key string, value int64) error {

 i := strings.Index(key, ".")
 if i <= 0 {
  return fmt.Errorf("pool id not presented in key '%s'", key)
 }
 poolID, err := strconv.Atoi(key[:i])
 if err != nil {
  return fmt.Errorf("pool id must be number: %v", err)
 }
 if poolID > len(pm.Pools) {
  return fmt.Errorf("pool's settings must be set sequntualy by id")
 }
 if poolID == len(pm.Pools) {
  pm.Pools = append(pm.Pools, &Pool{})
 }

 key = key[i+1:]
 switch {
 case strings.HasPrefix(key, "percent"):
  pm.Pools[poolID].Percent = value
 case strings.HasPrefix(key, "shares."):
  rank, err := strconv.Atoi(key[7:])
  if err != nil {
   return err
  }
  if pm.Pools[poolID].Shares == nil {
   pm.Pools[poolID].Shares = map[int64]int64{}
  }
  pm.Pools[poolID].Shares[int64(rank)] = value
 default:
  return fmt.Errorf("unknown key '%s'", key)
 }
 return nil
}

func (pm *PoolMarketing) MakePayment(_ RefChain, delta Coins) Coins {
 totalPayments := Coins{}
 for _, pool := range pm.Pools {
  payment := delta.Percent(pool.Percent)
  pool.balance = pool.balance.Plus(payment)
  totalPayments = totalPayments.Plus(payment)
 }
 return totalPayments
}


func (pm *PoolMarketing) MakePoolsPayment(refTree RefTree) {
 for _, p := range pm.Pools {
  p.MakePayment(refTree)
 }
}


func (p *Pool) MakePayment(refTree RefTree) {
 var totalShares int64
 var toPay []*Account
 for _, acc := range refTree {
  if p.Career > 0 {
   career := acc.Balance.CoinNoExpired(AMCCoinEasyBizziPoint)
   if career.Amount < p.Career {
    continue
   }
  }
  rank := acc.Balance.Coin(AMCCoinRank)
  if shares, ok := p.Shares[rank.Amount]; ok {
   totalShares += shares
   toPay = append(toPay, acc)
  }
 }

 for _, acc := range toPay {
  rank := acc.Balance.Coin(AMCCoinRank)
  shares := p.Shares[rank.Amount]
  payment := p.balance.Fraction(float64(shares) / float64(totalShares))
  p.balance = p.balance.Minus(payment)
  acc.Balance = acc.Balance.Plus(payment)
 }
}

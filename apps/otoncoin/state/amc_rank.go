package state

import (
 "fmt"
 "strconv"
 "strings"
 "time"
)


type RankAssigning struct {
 Ranks []int64
 Packets map[int64]int64
}

func (ra *RankAssigning) SetParam(key string, value int64) error {
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
  if ra.Packets == nil {
   ra.Packets = map[int64]int64{}
  }
  ra.Packets[int64(packID)] = value
 case "rank":
  r, err := strconv.Atoi(p[1])
  if err != nil {
   return err
  }

  last := len(ra.Ranks) - 1
  if r != last+1 {
   return fmt.Errorf("ranks must sets sequentially, next rank '%d', received '%d'", last+1, r)
  }
  if ra.Ranks[last] >= value {
   return fmt.Errorf("turnover '%d' must be larger than the last one '%d'", value, ra.Ranks[last])
  }
  ra.Ranks = append(ra.Ranks, value)
 default:
  return fmt.Errorf("unknown key part '%s'", p[0])
 }
 return nil
}

func (ra *RankAssigning) MakePayment(refChain RefChain, delta Coins) Coins {
 for i, acc := range refChain {
  if i == 0 {
   continue
  }

  packet := acc.Balance.Coin(AMCCoinPacket)
  if packet == nil || packet.Amount == 0 {
   continue
  }


  if i == 1 {
   acc.Balance = acc.Balance.Plus(Coins{Coin{
    Name: AMCCoinEasyBizziPoint,
    Amount: delta[0].Amount,
    Expired: time.Now().Add(28 * 24 * time.Hour),
   }})
  }


  acc.Balance = acc.Balance.Plus(Coins{Coin{
   Name: AMCCoinCareerPoint,
   Amount: delta[0].Amount,
  }})


  rank := acc.Balance.Coin(AMCCoinRank)
  if rank == nil {
   rank = &Coin{Name: AMCCoinRank}
  }
  if rank.Amount >= ra.Packets[packet.Amount] {
   continue
  }

  careerPoints := acc.Balance.Coin(AMCCoinCareerPoint)
  ranksToAdd := Coin{Name: AMCCoinRank}
  for r, t := range ra.Ranks {
   if t > careerPoints.Amount {
    break
   }
   if int64(r) > rank.Amount {
    ranksToAdd.Amount++
   }
  }
  acc.Balance = acc.Balance.Plus(Coins{ranksToAdd})
 }
 return Coins{}
}

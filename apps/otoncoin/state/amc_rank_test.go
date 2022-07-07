package state

import (
 "github.com/stretchr/testify/assert"
 "testing"
)

func TestRank_MakePayment(t *testing.T) {
 rank := &RankAssigning{
  Ranks: []int64{0, 800, 2500, 7000, 20000, 50000, 150_000, 500_000, 1_000_000},
  Packets: map[int64]int64{1: 1, 2: 2, 3: 4, 4: 8},
 }
 makeCoins := func(name string, value int64) Coins {
  return Coins{Coin{
   Name: name,
   Amount: value,
  }}
 }
 delta := makeCoins("test_coin", 100)

 t.Run("standard rank ", func(t *testing.T) {
  rc := RefChain{
   &Account{Balance: Coins{
    {Name: AMCCoinCareerPoint, Amount: 25000},
    {Name: AMCCoinPacket, Amount: 4},
    {Name: AMCCoinRank, Amount: 4},
   }},
   &Account{Balance: Coins{
    {Name: AMCCoinCareerPoint, Amount: 1000},
    {Name: AMCCoinPacket, Amount: 1},
    {Name: AMCCoinRank, Amount: 1},
   }},
   &Account{Balance: Coins{
    {Name: AMCCoinCareerPoint, Amount: 2450},
    {Name: AMCCoinPacket, Amount: 2},
    {Name: AMCCoinRank, Amount: 1},
   }},
   &Account{Balance: Coins{
    {Name: AMCCoinCareerPoint, Amount: 10000},
    {Name: AMCCoinPacket, Amount: 3},
    {Name: AMCCoinRank, Amount: 3},
   }},
  }
  payments := rank.MakePayment(rc, delta)
  assert.Equal(t, Coins{}, payments)
  assert.Equal(t, &makeCoins(AMCCoinCareerPoint, 25000)[0], rc[0].Balance.Coin(AMCCoinCareerPoint))
  assert.Equal(t, &makeCoins(AMCCoinRank, 4)[0], rc[0].Balance.Coin(AMCCoinRank))
  assert.Equal(t, &makeCoins(AMCCoinCareerPoint, 1100)[0], rc[1].Balance.Coin(AMCCoinCareerPoint))
  assert.Equal(t, &makeCoins(AMCCoinRank, 1)[0], rc[1].Balance.Coin(AMCCoinRank))
  assert.Equal(t, &makeCoins(AMCCoinCareerPoint, 2550)[0], rc[2].Balance.Coin(AMCCoinCareerPoint))
  assert.Equal(t, &makeCoins(AMCCoinRank, 2)[0], rc[2].Balance.Coin(AMCCoinRank))
  assert.Equal(t, &makeCoins(AMCCoinCareerPoint, 10100)[0], rc[3].Balance.Coin(AMCCoinCareerPoint))
  assert.Equal(t, &makeCoins(AMCCoinRank, 3)[0], rc[3].Balance.Coin(AMCCoinRank))
 })
}

package state

import (
 "github.com/stretchr/testify/assert"
 "testing"
)

func TestLeader_MakePayment(t *testing.T) {
 leader := &LeaderMarketing{
  Levels: 3,
  Ranks: map[int64]int64{6: 1, 7: 1, 8: 1, 9: 1, 10: 1, 11: 1},
 }
 makeCoins := func(value int64) Coins {
  return Coins{Coin{
   Name: "test_coin",
   Amount: value,
  }}
 }
 delta := makeCoins(100)

 t.Run("standard leader payments", func(t *testing.T) {
  rc := RefChain{
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 4,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 11,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 2,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 7,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 8,
   }}},
  }
  payments := leader.MakePayment(rc, delta)
  assert.Equal(t, makeCoins(2), payments)
  assert.Equal(t, (*Coin)(nil), rc[0].Balance.Coin("test_coin"))
  assert.Equal(t, &makeCoins(1)[0], rc[1].Balance.Coin("test_coin"))
  assert.Equal(t, (*Coin)(nil), rc[2].Balance.Coin("test_coin"))
  assert.Equal(t, &makeCoins(1)[0], rc[3].Balance.Coin("test_coin"))
  assert.Equal(t, (*Coin)(nil), rc[4].Balance.Coin("test_coin"))
 })
}

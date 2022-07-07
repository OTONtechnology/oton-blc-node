package state

import (
 "testing"

 "github.com/stretchr/testify/assert"
)

func TestLine_MakePayment(t *testing.T) {
 line := &LineMarketing{

  Levels: []int64{0, 10, 5, 4, 3, 2, 1, 1, 1},
  Packets: map[int64]int{1: 1, 2: 2, 3: 4, 4: 8},
 }
 makeCoins := func(value int64) Coins {
  return Coins{Coin{
   Name: "test_coin",
   Amount: value,
  }}
 }
# 32 "amc_line_test.go"
 delta := makeCoins(100)

 t.Run("standard line payments", func(t *testing.T) {
  rc := RefChain{
   &Account{Balance: Coins{Coin{
    Name: AMCCoinPacket,
    Amount: 4,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinPacket,
    Amount: 1,



   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinPacket,
    Amount: 2,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinPacket,
    Amount: 3,
   }}},
  }




  payments, _ := line.MakePayment(rc, delta)

  assert.Equal(t, makeCoins(19), payments)
  assert.Equal(t, (*Coin)(nil), rc[0].Balance.Coin("test_coin"))
  assert.Equal(t, &makeCoins(10)[0], rc[1].Balance.Coin("test_coin"))
  assert.Equal(t, &makeCoins(5)[0], rc[2].Balance.Coin("test_coin"))
  assert.Equal(t, &makeCoins(4)[0], rc[3].Balance.Coin("test_coin"))
 })

 t.Run("dynamic compression", func(t *testing.T) {
  t.Run("empty balance", func(t *testing.T) {
   rc := RefChain{
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 4,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 1,
    }}},
    &Account{},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 3,
    }}},
   }
   payments, _ := line.MakePayment(rc, delta)
   assert.Equal(t, makeCoins(15), payments)
   assert.Equal(t, (*Coin)(nil), rc[0].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(10)[0], rc[1].Balance.Coin("test_coin"))
   assert.Equal(t, (*Coin)(nil), rc[2].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(5)[0], rc[3].Balance.Coin("test_coin"))
  })
  t.Run("zero packets amount", func(t *testing.T) {
   rc := RefChain{
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 4,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 1,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 0,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 3,
    }}},
   }
   payments, _ := line.MakePayment(rc, delta)
   assert.Equal(t, makeCoins(15), payments)
   assert.Equal(t, (*Coin)(nil), rc[0].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(10)[0], rc[1].Balance.Coin("test_coin"))
   assert.Equal(t, (*Coin)(nil), rc[2].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(5)[0], rc[3].Balance.Coin("test_coin"))
  })
  t.Run("available levels", func(t *testing.T) {
   rc := RefChain{
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 4,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 1,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 1,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinPacket,
     Amount: 3,
    }}},
   }
   payments, _ := line.MakePayment(rc, delta)
   assert.Equal(t, makeCoins(15), payments)
   assert.Equal(t, (*Coin)(nil), rc[0].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(10)[0], rc[1].Balance.Coin("test_coin"))
   assert.Equal(t, (*Coin)(nil), rc[2].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(5)[0], rc[3].Balance.Coin("test_coin"))
  })
 })
}

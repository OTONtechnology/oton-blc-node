package state

import (
 "testing"

 "github.com/stretchr/testify/assert"
)

func TestCareer_MakePayment(t *testing.T) {
 career := &CareerMarketing{
  Packets: map[int64]int64{1: 1, 2: 4, 3: 8, 4: 11},
  Ranks: map[int64]int64{1: 10, 2: 12, 3: 14, 4: 16, 5: 20, 6: 24, 7: 28, 8: 31, 9: 33, 10: 34, 11: 35},
  maxRank: 11,
 }
 makeCoins := func(value int64) Coins {
  return Coins{Coin{
   Name: "test_coin",
   Amount: value,
  }}
 }
 delta := makeCoins(100)

 t.Run("standard career payments", func(t *testing.T) {
  rc := RefChain{
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 4,
   }, Coin{
    Name: AMCCoinPacket,
    Amount: 4,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 1,
   }, Coin{
    Name: AMCCoinPacket,
    Amount: 1,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 2,
   }, Coin{
    Name: AMCCoinPacket,
    Amount: 2,
   }}},
   &Account{Balance: Coins{Coin{
    Name: AMCCoinRank,
    Amount: 3,
   }, Coin{
    Name: AMCCoinPacket,
    Amount: 3,
   }}},
  }

  payments, _ := career.MakePayment(rc, delta)
  assert.Equal(t, makeCoins(14), payments)
  assert.Equal(t, (*Coin)(nil), rc[0].Balance.Coin("test_coin"))
  assert.Equal(t, &makeCoins(10)[0], rc[1].Balance.Coin("test_coin"))
  assert.Equal(t, &makeCoins(2)[0], rc[2].Balance.Coin("test_coin"))
  assert.Equal(t, &makeCoins(2)[0], rc[3].Balance.Coin("test_coin"))
 })

 t.Run("dynamic compression", func(t *testing.T) {
  t.Run("empty packages", func(t *testing.T) {
   rc := RefChain{
    &Account{Balance: Coins{Coin{
     Name: AMCCoinRank,
     Amount: 4,
    }, Coin{
     Name: AMCCoinPacket,
     Amount: 4,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinRank,
     Amount: 1,
    }, Coin{
     Name: AMCCoinPacket,
     Amount: 1,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinRank,
     Amount: 2,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinRank,
     Amount: 3,
    }, Coin{
     Name: AMCCoinPacket,
     Amount: 3,
    }}},
   }
   payments, _ := career.MakePayment(rc, delta)
   assert.Equal(t, makeCoins(14), payments)
   assert.Equal(t, (*Coin)(nil), rc[0].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(10)[0], rc[1].Balance.Coin("test_coin"))
   assert.Equal(t, (*Coin)(nil), rc[2].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(4)[0], rc[3].Balance.Coin("test_coin"))
  })
  t.Run("not enough rank", func(t *testing.T) {
   rc := RefChain{
    &Account{Balance: Coins{Coin{
     Name: AMCCoinRank,
     Amount: 4,
    }, Coin{
     Name: AMCCoinPacket,
     Amount: 4,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinRank,
     Amount: 1,
    }, Coin{
     Name: AMCCoinPacket,
     Amount: 1,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinRank,
     Amount: 1,
    }, Coin{
     Name: AMCCoinPacket,
     Amount: 2,
    }}},
    &Account{Balance: Coins{Coin{
     Name: AMCCoinRank,
     Amount: 3,
    }, Coin{
     Name: AMCCoinPacket,
     Amount: 3,
    }}},
   }
   payments, _ := career.MakePayment(rc, delta)
   assert.Equal(t, makeCoins(14), payments)
   assert.Equal(t, (*Coin)(nil), rc[0].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(10)[0], rc[1].Balance.Coin("test_coin"))
   assert.Equal(t, (*Coin)(nil), rc[2].Balance.Coin("test_coin"))
   assert.Equal(t, &makeCoins(4)[0], rc[3].Balance.Coin("test_coin"))
  })
 })
}

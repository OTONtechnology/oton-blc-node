package state


import (
 "testing"

 "github.com/stretchr/testify/assert"
 "github.com/tendermint/tendermint/crypto/ed25519"
)

func TestNewAMC(t *testing.T) {
 address := ed25519.GenPrivKey().PubKey().Address()
 amc, err := NewAMC(
  []byte("test"),
  "test",
  address,
  nil,
  []AMCParam{
   {Key: "line.level.2", Value: int64(2)},
   {Key: "line.level.3", Value: int64(3)},
   {Key: "line.packet.1", Value: int64(1)},





  })
 assert.NoError(t, err)
 assert.Equal(t, "test", amc.Name)
 assert.Equal(t, address, amc.Creator)
 assert.Nil(t, amc.Balance)

 assert.NotEmpty(t, amc.PaymentMakers)
 assert.NotNil(t, amc.PaymentMakers[0])


 assert.Equal(t, int64(0), amc.PaymentMakers[0].(*LineMarketing).Levels[1])
 assert.Equal(t, int64(2), amc.PaymentMakers[0].(*LineMarketing).Levels[2])
 assert.Equal(t, int64(3), amc.PaymentMakers[0].(*LineMarketing).Levels[3])
 assert.Equal(t, 1, amc.PaymentMakers[0].(*LineMarketing).Packets[1])






}

func TestAMC_MakePayment(t *testing.T) {
 t.Run("C1200", func(t *testing.T) {
  amc, err := NewAMC(
   []byte("test"),
   "test",
   ed25519.GenPrivKey().PubKey().Address(),
   nil,
   []AMCParam{})
  assert.NoError(t, err)

  t.Run("Step1", func(t *testing.T) {
   refChain := RefChain{
    &Account{Balance: Coins{{Name: AMCCoinPacket, Amount: 1}}},
    &Account{},
    &Account{Balance: Coins{{Name: AMCCoinPacket, Amount: 1}}},
    &Account{Balance: Coins{{Name: AMCCoinPacket, Amount: 4}}},
    &Account{Balance: Coins{{Name: AMCCoinPacket, Amount: 4}}},
   }

   _, err := amc.MakePayment(refChain, Coins{{Name: "coin", Amount: 100}})
   assert.NoError(t, err)

   assert.Equal(t,
    &Coin{Name: AMCCoinCareerPoint, Amount: 100},
    refChain[4].Balance.Coin(AMCCoinCareerPoint))
   assert.Equal(t,
    (*Coin)(nil),
    refChain[4].Balance.Coin(AMCCoinEasyBizziPoint))

   assert.Equal(t,
    &Coin{Name: AMCCoinCareerPoint, Amount: 100},
    refChain[3].Balance.Coin(AMCCoinCareerPoint))
   assert.Equal(t,
    (*Coin)(nil),
    refChain[3].Balance.Coin(AMCCoinEasyBizziPoint))

   assert.Equal(t,
    &Coin{Name: AMCCoinCareerPoint, Amount: 100},
    refChain[2].Balance.Coin(AMCCoinCareerPoint))
   assert.Equal(t,
    (*Coin)(nil),
    refChain[2].Balance.Coin(AMCCoinEasyBizziPoint))

   assert.Equal(t,
    (*Coin)(nil),
    refChain[1].Balance.Coin(AMCCoinCareerPoint))
   assert.Equal(t,
    (*Coin)(nil),
    refChain[1].Balance.Coin(AMCCoinEasyBizziPoint))
  })

  t.Run("Step2", func(t *testing.T) {
   refChain := RefChain{
    &Account{Balance: Coins{{Name: AMCCoinPacket, Amount: 1}}},
    &Account{Balance: Coins{{Name: AMCCoinPacket, Amount: 2}}},
    &Account{},
    &Account{Balance: Coins{{Name: AMCCoinPacket, Amount: 4}}},
    &Account{Balance: Coins{{Name: AMCCoinPacket, Amount: 4}}},
   }

   _, err := amc.MakePayment(refChain, Coins{{Name: "coin", Amount: 100}})
   assert.NoError(t, err)

   assert.Equal(t,
    &Coin{Name: AMCCoinCareerPoint, Amount: 100},
    refChain[4].Balance.Coin(AMCCoinCareerPoint))
   assert.Equal(t,
    (*Coin)(nil),
    refChain[4].Balance.Coin(AMCCoinEasyBizziPoint))

   assert.Equal(t,
    &Coin{Name: AMCCoinCareerPoint, Amount: 100},
    refChain[3].Balance.Coin(AMCCoinCareerPoint))
   assert.Equal(t,
    (*Coin)(nil),
    refChain[3].Balance.Coin(AMCCoinEasyBizziPoint))

   assert.Equal(t,
    (*Coin)(nil),
    refChain[2].Balance.Coin(AMCCoinCareerPoint))
   assert.Equal(t,
    (*Coin)(nil),
    refChain[2].Balance.Coin(AMCCoinEasyBizziPoint))

   assert.Equal(t,
    &Coin{Name: AMCCoinCareerPoint, Amount: 100},
    refChain[1].Balance.Coin(AMCCoinCareerPoint))
   ebp := refChain[1].Balance.Coin(AMCCoinEasyBizziPoint)
   assert.NotNil(t, ebp)
   assert.Equal(t, int64(100), ebp.Amount)
  })
 })
}

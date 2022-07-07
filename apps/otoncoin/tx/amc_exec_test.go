package tx

import (
 "fmt"
 "testing"

 "github.com/stretchr/testify/assert"
 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
 "github.com/tendermint/tendermint/abci/types"

 "github.com/tendermint/tendermint/crypto/ed25519"
 "github.com/tendermint/tendermint/libs/json"
)


func TestAmc_Exec(t *testing.T) {

 t.Run("ok amc create, set acc to amc", func(t *testing.T) {
  privKey1 := ed25519.GenPrivKey()
  pubKey1 := privKey1.PubKey()

  privKey2 := ed25519.GenPrivKey()
  pubKey2 := privKey2.PubKey()

  privKey21 := ed25519.GenPrivKey()
  pubKey21 := privKey21.PubKey()


  privKey3 := ed25519.GenPrivKey()
  pubKey3 := privKey3.PubKey()

  store := state.NewMemKVStore()
  s := state.NewState(store)
  s.SetChainID("test")

  acc1 := &state.Account{
   Balance: state.Coins{
    state.Coin{Name: "feecoin", Amount: 10000},
    state.Coin{Name: "testcoin", Amount: 25000},
   },
   PubKey: pubKey1,
  }
  err := s.SetAccount(pubKey1.Address(), acc1)
  assert.NoError(t, s.SetAccount(pubKey1.Address(), acc1))


  s.SetCoin(&state.Coin{Name: "feecoin", Amount: 6000, Movable: true})
  s.SetCoin(&state.Coin{Name: "testcoin", Amount: 6000, Movable: true})
  s.SetCoin(&state.Coin{Name: "mycoin", Amount: 6000, Movable: true})

  acc2 := &state.Account{
   Balance: state.Coins{
    state.Coin{Name: "mycoin", Amount: 15},
    state.Coin{Name: "testcoin", Amount: 25000},
   },
   PubKey: pubKey2,
  }


  err = s.SetAccount(pubKey2.Address(), acc2)
  assert.NoError(t, err)

  tx := &CreateAMC{
   Fee: Coin{"feecoin", 1},
   Inputs: Inputs{
    Input{
     Address: acc1.PubKey.Address(),
     Coins: Coins{Coin{Name: "feecoin", Amount: 1}},
     Sequence: 1,
     PubKey: acc1.PubKey.Bytes(),
    },
   },
   Address: pubKey3.Address(),
   Name: "new_amC",


   Params: []AMCParam{
    {Key: "line.level.1", Value: int64(50)},
    {Key: "line.level.2", Value: int64(20)},
    {Key: "line.level.3", Value: int64(10)},
    {Key: "line.packet.1", Value: int64(10)},

    {Key: "career.packet.3", Value: int64(10)},
    {Key: "career.rank.3", Value: int64(10)},

    {Key: "leader.rank.1", Value: int64(1)},
    {Key: "leader.level.3", Value: int64(10)},
    {Key: "pool.0.percent.3", Value: int64(10)},
   },
  }


  signBytes := SignBytes(tx, "test")
  sig, err := privKey1.Sign(signBytes)
  tx.Inputs[0].Signature = sig
  assert.NoError(t, err)


  bz2, _ := json.MarshalIndent(tx, "", "\t")
  fmt.Printf("tx CreateAMC %s \n \n", bz2)

  hash := []byte("test_hash")
  res := tx.Exec(s, hash, false)
  assert.Equal(t, types.ResponseDeliverTx{
   Code: TypeOK,
   Data: []byte("CreateAMC"),
   Log: "Ok",
  }, res)


  fmt.Printf("\nAdd referal to tree...\n")
  tx2 := &SetInAMC{
   Fee: Coin{"feecoin", 3},
   Inputs: Inputs{
    Input{
     Address: acc1.PubKey.Address(),
     Coins: Coins{Coin{Name: "feecoin", Amount: 3}},
     Sequence: 2,

    },
   },
   Address: pubKey3.Address(),

   Referal: acc2.PubKey.Address(),
   Sponsor: acc1.PubKey.Address(),
  }

  signBytes = SignBytes(tx2, "test")
  sig, err = privKey1.Sign(signBytes)
  tx2.Inputs[0].Signature = sig

  res = tx2.Exec(s, hash, false)
  assert.Equal(t, types.ResponseDeliverTx{
   Code: TypeOK,
   Data: []byte("SetInAMC"),
   Log: "Ok",
  }, res)

  spn, _ := s.GetAccount(acc1.PubKey.Address())
  fmt.Printf("get acc1 %v \n", spn)
  spn2, _ := s.GetAccount(acc2.PubKey.Address())
  fmt.Printf("get acc2 %v \n \n", spn2)


  amc, err := s.GetAMC(tx.Address)
  if err != nil || amc == nil {
   fmt.Printf("Error get AMC: %v", err)
  } else {
   amc.Balance = state.Coins{
    state.Coin{Name: "feecoin", Amount: 3000, Movable: true},

    state.Coin{Name: "Packet", Amount: 1, Movable: false}}


   s.SetAMC(amc)
   fmt.Printf("Result AMC, add coins/packet/tarif: \n%v \n \n \n", amc)
  }


  amc.Balance = amc.Balance.Plus(state.Coins{state.Coin{Name: "feecoin", Amount: 333, Movable: true}})
  fmt.Printf("amc.Balance(non fix) .... %v \n", amc.Balance)


  fmt.Printf("Buy tarif for referal...\n")

  tx3 := &BuyInAMC{
   Fee: Coin{"feecoin", 3},
   Inputs: Inputs{
    Input{
     Address: acc1.PubKey.Address(),




     Coins: Coins{Coin{Name: "feecoin", Amount: 3}, Coin{Name: "testcoin", Amount: 199}},
     Sequence: 3,
    },
   },
   Address: pubKey3.Address(),

   Referal: acc2.PubKey.Address(),
   Value: Coin{Name: "TarifBitBone", Amount: 199},
   Delta: Coin{Name: "testcoin", Amount: 99},
  }

  signBytes = SignBytes(tx3, "test")
  sig, err = privKey1.Sign(signBytes)
  tx3.Inputs[0].Signature = sig
  res = tx3.Exec(s, hash, false)
  bz2, _ = json.MarshalIndent(res, "", "\t")
  fmt.Printf("exec BuyInAmc %s \n \n", bz2)

  spn, _ = s.GetAccount(acc1.PubKey.Address())
  fmt.Printf("get acc1 %v \n", spn)
  spn2, _ = s.GetAccount(acc2.PubKey.Address())
  fmt.Printf("get acc2 %v \n \n \n", spn2)
  amc2, err := s.GetAMC(tx.Address)
  fmt.Printf("Result AMC.Balance: \n%v \n \n \n", amc2.Balance)


  fmt.Printf("Add next referal...\n")
  tx2.Referal = pubKey21.Address()
  tx2.Inputs[0].Sequence = 4
  tx2.Sponsor = acc2.PubKey.Address()
  signBytes = SignBytes(tx2, "test")
  sig, err = privKey1.Sign(signBytes)
  tx2.Inputs[0].Signature = sig
  res = tx2.Exec(s, hash, false)
  bz2, _ = json.MarshalIndent(res, "", "\t")
  fmt.Printf("exec SetInAMC %s \n \n", bz2)


  fmt.Printf("Buy for next referal...\n")
  tx3.Referal = pubKey21.Address()
  tx3.Inputs[0].Sequence = 5
  signBytes = SignBytes(tx3, "test")
  sig, err = privKey1.Sign(signBytes)
  tx3.Inputs[0].Signature = sig
  res = tx3.Exec(s, hash, false)

  bz2, _ = json.MarshalIndent(res, "", "\t")
  fmt.Printf("exec BuyInAmc %s \n \n", bz2)

  amc3, err := s.GetAMC(tx.Address)
  fmt.Printf("Result AMC.Balance: \n%v \n \n \n", amc3.Balance)

 })
}

func TestAmc_CurlTX(t *testing.T) {

}

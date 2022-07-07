package tx



import (
 "math/rand"

 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
 "github.com/tendermint/tendermint/crypto/ed25519"
)


func RandAccounts(num int, minAmount int64, maxAmount int64) ([]state.Account, []ed25519.PrivKey) {
 privAccs := make([]state.Account, num)
 privKeyAccs := make([]ed25519.PrivKey, num)

 for i := 0; i < num; i++ {

  balance := minAmount
  if maxAmount > minAmount {
   balance += int64(rand.Intn(100)) % (maxAmount - minAmount)
  }

  privKey1 := ed25519.GenPrivKey()
  pubKey1 := privKey1.PubKey()

  privKeyAccs[i] = privKey1
  privAccs[i] = state.Account{
   Address: pubKey1.Address(),
   Balance: state.Coins{state.Coin{Name: "mycoin", Amount: balance}},
   PubKey: pubKey1,
  }
 }
 return privAccs, privKeyAccs
}

func Accs2TxInputs(seq int, accs ...state.Account) []Input {
 var txs []Input

 for _, acc := range accs {
  tx := Input{
   Address: acc.PubKey.Address(),
   Coins: Coins{Coin{Name: "mycoin", Amount: 5}},
   Sequence: int64(seq),
   PubKey: acc.PubKey.Bytes(),
  }
  txs = append(txs, tx)
 }
 return txs
}


func Accs2TxOutputs(accs ...state.Account) []Output {
 var txs []Output

 for _, acc := range accs {
  tx := Output{
   Address: acc.PubKey.Address(),
   Coins: Coins{Coin{Name: "mycoin", Amount: 4}},
  }

  txs = append(txs, tx)
 }
 return txs
}

func MakeSendTx(seq int, accOut state.Account, accsIn ...state.Account) *SendCoins {
 tx := &SendCoins{
  Fee: Coin{"mycoin", 1},
  Inputs: Accs2TxInputs(seq, accsIn...),
  Outputs: Accs2TxOutputs(accOut),
 }
 return tx
}

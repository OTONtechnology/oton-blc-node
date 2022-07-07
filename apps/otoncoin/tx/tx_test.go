package tx

import (
 "testing"

 "github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
 tx := &SendCoins{
  Fee: Coin{Name: "testcoin", Amount: 1},
  Gas: 11,
  Inputs: Inputs{Input{
   Address: []byte("1234"),
   Signature: []byte("signature"),
   Sequence: 1,
   PubKey: []byte("pubkey"),
  }},
  Outputs: Outputs{},
 }
 txCopy := Copy(tx).(*SendCoins)
 assert.NotSame(t, tx.Inputs, txCopy.Inputs)
 assert.Empty(t, txCopy.Inputs[0].Signature)
 assert.Equal(t, &SendCoins{
  Fee: tx.Fee,
  Gas: tx.Gas,
  Inputs: Inputs{Input{
   Address: tx.Inputs[0].Address,
   Signature: Signature{},
   Sequence: tx.Inputs[0].Sequence,
   PubKey: tx.Inputs[0].PubKey,
  }},
  Outputs: tx.Outputs,
 }, txCopy)
}

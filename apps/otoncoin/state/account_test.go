package state

import (
 "testing"

 "github.com/stretchr/testify/assert"
 "github.com/tendermint/tendermint/crypto/ed25519"
)

func TestNilAccount(t *testing.T) {

 var acc Account
 var acc2 *Account

 privKey := ed25519.GenPrivKey()
 pubKey := privKey.PubKey()


 accCopy := acc.Copy()

 assert.True(t, &acc != accCopy, "Account Copy Error, acc1: %v, acc2: %v", &acc, accCopy)
 assert.Equal(t, acc.Sequence, accCopy.Sequence)

 acc.Sequence = 123
 acc.PubKey = pubKey
 accBytes, err := Serialize(&acc)
 assert.NoError(t, err)

 err = Deserialize(accBytes, &acc2)
 assert.NoError(t, err)
}

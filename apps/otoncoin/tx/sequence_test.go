package tx

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/abci/apps/otoncoin/state"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	prvKeyHex1 = ""
)

func TestSequence(t *testing.T) {
	var st string

	var Key1, _ = hex.DecodeString(prvKeyHex1)
	var privKey1 ed25519.PrivKey

	privKey1 = Key1

	pubKey1 := privKey1.PubKey()

	store := state.NewMemKVStore()
	s := state.NewState(store)

	chainID := ""
	s.SetChainID(chainID)

	acc1 := &state.Account{
		Address: pubKey1.Address(),
		Balance: state.Coins{
			state.Coin{Name: "mycoin", Amount: 1 << 53},
			state.Coin{Name: "testcoin", Amount: 1 << 53},
		},
		PubKey: pubKey1,
	}
	s.SetAccount(pubKey1.Address(), acc1)

	privAccounts, pKey := RandAccounts(30, 1000000, 0)
	privAccountSequences := make(map[string]int64)

	sequence := int64(91)
	for i := 0; i < len(privAccounts); i++ {
		privAccount := privAccounts[i]

		tx := &SendCoins{
			Fee: Coin{"mycoin", 2},
			Inputs: Inputs{
				Input{
					Address: acc1.PubKey.Address(),
					Coins: Coins{
						Coin{Name: "mycoin", Amount: 1000002},
						Coin{Name: "testcoin", Amount: 1000000}},
					Sequence: sequence,
					PubKey:   acc1.PubKey.Bytes(),
				},
			},
			Outputs: Outputs{
				Output{
					Address: privAccount.PubKey.Address(),
					Coins: Coins{
						Coin{Name: "mycoin", Amount: 1000000},
						Coin{Name: "testcoin", Amount: 1000000}},
				},
			},
		}

		sequence += 1
		if sequence != 2 {
			tx.Inputs[0].PubKey = nil
		}

		signBytes := SignBytes(tx, chainID)
		sig, err := privKey1.Sign(signBytes)
		tx.Inputs[0].Signature = sig
		assert.NoError(t, err)
		fmt.Printf("ADDR: %X -> %X , seq %d \n", tx.Inputs[0].Address, tx.Outputs[0].Address, tx.Inputs[0].Sequence)

		var upTxRaw Raw
		upTxRaw.Type = TypeSendCoins
		upTxRaw.Raw, _ = tx.Marshal()
		rw, _ := upTxRaw.Marshal()

		st = st + fmt.Sprintf("curl -s '127.0.0.1:26657/broadcast_tx_commit?tx=0x%X' \n", rw)

	}

	fmt.Printf("\n%s \n\n", st)

	t.Log("-------------------- RANDOM SENDS --------------------")
	st = " "

	for i := 0; i < 500; i++ {
		randA := int(rand.Intn(100)) % len(privAccounts)
		randB := int(rand.Intn(100)) % len(privAccounts)

		if randA == randB {
			continue
		}

		privAccountA := privAccounts[randA]
		privAccountASequence := privAccountSequences[string(privAccountA.PubKey.Address())]
		privAccountSequences[string(privAccountA.PubKey.Address())] = privAccountASequence + 1
		privAccountB := privAccounts[randB]

		privKeyA := pKey[randA]

		tx := &SendCoins{
			Fee: Coin{"mycoin", 2},
			Inputs: Inputs{
				Input{
					Address:  privAccountA.PubKey.Address(),
					Coins:    Coins{Coin{Name: "mycoin", Amount: 3}, Coin{Name: "testcoin", Amount: 1}},
					Sequence: privAccountASequence + 1,
					PubKey:   privAccountA.PubKey.Bytes(),
				},
			},
			Outputs: Outputs{
				Output{
					Address: privAccountB.PubKey.Address(),
					Coins:   Coins{Coin{Name: "mycoin", Amount: 1}, Coin{Name: "testcoin", Amount: 1}},
				},
			},
		}

		if privAccountASequence+1 != 1 {
			tx.Inputs[0].PubKey = nil
		}

		signBytes := SignBytes(tx, chainID)
		sig, _ := privKeyA.Sign(signBytes)
		tx.Inputs[0].Signature = sig
		fmt.Printf("ADDR: %X -> %X , seq %d \n", tx.Inputs[0].Address, tx.Outputs[0].Address, tx.Inputs[0].Sequence)

		var upTxRaw Raw
		upTxRaw.Type = TypeSendCoins
		upTxRaw.Raw, _ = tx.Marshal()
		rw, _ := upTxRaw.Marshal()

		st = st + fmt.Sprintf("curl -s '127.0.0.1:26657/broadcast_tx_commit?tx=0x%X' \n", rw)

	}

	fmt.Printf("\n%s  \n\n", st)
}

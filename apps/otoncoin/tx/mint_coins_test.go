package tx

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/abci/apps/otoncoin/state"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	prvKeyHex1 = ""

	prvKeyHex2 = ""
)

func TestMintCoins_Exec(t *testing.T) {
	var upTxRaw Raw

	t.Run("mint coin to one output", func(t *testing.T) {

		var Key1, _ = hex.DecodeString(prvKeyHex1)
		var Key2, _ = hex.DecodeString(prvKeyHex2)
		var privKey1 ed25519.PrivKey
		var privKey2 ed25519.PrivKey

		privKey1 = Key1
		pubKey1 := privKey1.PubKey()

		privKey2 = Key2
		pubKey2 := privKey2.PubKey()

		store := state.NewMemKVStore()
		s := state.NewState(store)
		s.SetChainID("")

		acc1 := &state.Account{
			Balance: state.Coins{
				state.Coin{Name: "feecoin", Amount: 150},
			},
			PubKey: pubKey1,
		}
		err := s.SetAccount(pubKey1.Address(), acc1)
		assert.NoError(t, s.SetAccount(pubKey1.Address(), acc1))

		acc2 := &state.Account{
			Balance: state.Coins{
				state.Coin{Name: "mycoin", Amount: 15},
				state.Coin{Name: "testcoin", Amount: 25},
			},
			PubKey: pubKey2,
		}
		err = s.SetAccount(pubKey2.Address(), acc2)
		assert.NoError(t, err)

		s.SetCoin(&state.Coin{Name: "feecoin", Amount: 300, Movable: true})

		s.ValProc.Minter = acc1.PubKey.Address()

		tx := &MintCoins{
			Fee: Coin{"feecoin", 1},
			Inputs: Inputs{
				Input{
					Address:  acc1.PubKey.Address(),
					Coins:    Coins{Coin{Name: "feecoin", Amount: 1}},
					Sequence: 1,
					PubKey:   acc1.PubKey.Bytes(),
				},
			},
			Outputs: Outputs{
				Output{
					Address: acc2.PubKey.Address(),
					Coins:   Coins{Coin{Name: "new_coin", Amount: 102}},
				},
			},
			Name:         "new_coin",
			Amount:       102,
			DecimalPoint: 2,
			Movable:      true,
		}

		signBytes := SignBytes(tx, "")
		sig, err := privKey1.Sign(signBytes)
		tx.Inputs[0].Signature = sig
		assert.NoError(t, err)

		fmt.Printf("Bytes for sign %x \n", signBytes)
		fmt.Printf("sig %x \n", sig)
		fmt.Printf("privKey1 %x \n \n", privKey1)

		upTxRaw.Type = TypeMintCoins
		upTxRaw.Raw, _ = tx.Marshal()
		rw, _ := upTxRaw.Marshal()

		se := fmt.Sprintf("curl -s '127.0.0.1:26647/broadcast_tx_commit?tx=0x%X' \n \n", rw)
		fmt.Printf(se)

		hash := []byte("test_hash")
		res := tx.Exec(s, hash, false)
		assert.Equal(t, types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("mint"),
			Log:  "Ok",
		}, res)

		resAcc, err := s.GetAccount(pubKey2.Address())
		fmt.Printf("GetAccount 2: %v \n", resAcc)

		resAcc1, err := s.GetAccount(pubKey1.Address())
		fmt.Printf("GetAccount 1: %v \n", resAcc1)

		resCn, err := s.GetCoin(tx.Name)
		fmt.Printf("Get new mint coin: %v \n", resCn)

		fmt.Printf("\nAddon Mint ...\n")
		tx.Inputs[0].Sequence = 2
		tx.Inputs[0].PubKey = nil
		signBytes = SignBytes(tx, "")
		sig, err = privKey1.Sign(signBytes)
		tx.Inputs[0].Signature = sig
		res = tx.Exec(s, hash, false)
		assert.Equal(t, types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("mint"),
			Log:  "Ok",
		}, res)

		resAcc, err = s.GetAccount(pubKey2.Address())
		fmt.Printf("GetAccount 2: %v \n", resAcc)

		resCn, err = s.GetCoin(tx.Name)
		fmt.Printf("Get add mint coin: %v \n \n", resCn)

		resAcc, err = s.GetAccount(pubKey2.Address())
		assert.NoError(t, err)
		assert.Equal(t, state.Coin{
			Name:   tx.Name,
			Amount: tx.Amount * 2,
		}, resAcc.Balance[1])

		coin, err := s.GetCoin(tx.Name)
		assert.NoError(t, err)
		assert.Equal(t, &state.Coin{
			Name:         tx.Name,
			Amount:       tx.Amount * 2,
			CreatorAddr:  acc1.PubKey.Address(),
			DecimalPoint: tx.DecimalPoint,
			Movable:      tx.Movable,
			Hash:         hash,
		}, coin)

		address := ed25519.GenPrivKey().PubKey().Address()
		amc, err := state.NewAMC(
			address,
			"amctest",
			acc1.PubKey.Address(),
			nil,
			[]state.AMCParam{
				{Key: "line.level.2", Value: int64(2)},
				{Key: "line.level.3", Value: int64(3)},
				{Key: "line.packet.1", Value: int64(1)},
			})

		err = s.SetAMC(amc)
		assert.NoError(t, err)

		fmt.Printf("Addon Mint АМС ...\n")
		tx.Inputs[0].Sequence = 3
		tx.Inputs[0].PubKey = nil
		tx.Outputs[0].Address = address
		signBytes = SignBytes(tx, "")
		sig, err = privKey1.Sign(signBytes)
		tx.Inputs[0].Signature = sig
		res = tx.Exec(s, hash, false)
		assert.Equal(t, types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("mint"),
			Log:  "Ok",
		}, res)

		amc21, err := s.GetAMC(address)
		fmt.Printf("Get amc: %v \n \n", amc21)
		assert.NoError(t, err)
		assert.Equal(t, state.Coin{
			Name:   tx.Name,
			Amount: tx.Amount,
		}, amc21.Balance[0])

	})
}

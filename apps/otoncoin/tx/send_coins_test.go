package tx

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/abci/apps/otoncoin/state"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/json"
)

const (
	prvKeyHex0 = ""
)

func TestSendCoinsTx_Exec(t *testing.T) {
	t.Run("empty transaction", func(t *testing.T) {
		tx := &SendCoins{}
		store := state.NewMemKVStore()
		s := state.NewState(store)
		res := tx.Exec(s, nil, false)
		assert.Equal(t, CTErrValidateInputsBasic, res.Code)
	})

	t.Run("ok one input, one output", func(t *testing.T) {

		var Key1, _ = hex.DecodeString(prvKeyHex0)
		var privKey1 ed25519.PrivKey
		privKey1 = Key1
		pubKey1 := privKey1.PubKey()

		privKey2 := ed25519.GenPrivKey()
		pubKey2 := privKey2.PubKey()

		store := state.NewMemKVStore()
		s := state.NewState(store)

		s.SetChainID("test-chain-V47")
		fmt.Printf("\nChainID:_%v_1\n", s.GetChainID())

		acc1 := &state.Account{
			Address: pubKey1.Address(),
			Balance: state.Coins{
				state.Coin{Name: "bitbonCoin", Amount: 100},
				state.Coin{Name: "mycoin", Amount: 10},
				state.Coin{Name: "testcoin", Amount: 200},
			},
			PubKey: pubKey1,
		}
		s.SetAccount(pubKey1.Address(), acc1)

		acc2 := &state.Account{
			Address: pubKey2.Address(),
			Balance: state.Coins{
				state.Coin{Name: "mycoin", Amount: 15},
				state.Coin{Name: "testcoin", Amount: 20},
			},
			PubKey: pubKey2,
		}
		s.SetAccount(pubKey2.Address(), acc2)

		s.SetCoin(&state.Coin{Name: "testcoin", Amount: 300, Movable: true})
		s.SetCoin(&state.Coin{Name: "mycoin", Amount: 200, Movable: false})
		s.SetCoin(&state.Coin{Name: "bitbonCoin", Amount: 3000, Movable: true})

		tx := &SendCoins{
			Fee: Coin{"testcoin", 2},
			Inputs: Inputs{
				Input{
					Address:  acc1.PubKey.Address(),
					Coins:    Coins{Coin{Name: "testcoin", Amount: 11}},
					Sequence: 1,
					PubKey:   acc1.PubKey.Bytes(),
				},
			},
			Outputs: Outputs{
				Output{
					Address: acc2.PubKey.Address(),
					Coins:   Coins{Coin{Name: "testcoin", Amount: 9}},
				},
			},
		}

		signBytes := SignBytes(tx, s.GetChainID())
		fmt.Printf("\nChainID:_%v_2\n", s.GetChainID())

		sig, err := privKey1.Sign(signBytes)
		tx.Inputs[0].Signature = sig
		assert.NoError(t, err)

		fmt.Printf("Bytes for sign %x \n", signBytes)
		fmt.Printf("sig %x \n", sig)
		fmt.Printf("privKey1 %x \n \n", privKey1)

		var upTxRaw Raw
		upTxRaw.Type = TypeSendCoins
		upTxRaw.Raw, _ = tx.Marshal()
		rw, _ := upTxRaw.Marshal()

		se := fmt.Sprintf("curl -s '127.0.0.1:26647/broadcast_tx_commit?tx=0x%X' \n \n", rw)
		fmt.Printf(se)

		res := tx.Exec(s, nil, false)
		assert.Equal(t, types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("send"),
			Log:  "Ok",
		}, res)

		resAcc, err := s.GetAccount(pubKey2.Address())
		assert.NoError(t, err)
		assert.Equal(t, acc2.Balance[1].Amount+9, resAcc.Balance[1].Amount)

		stx, _ := json.Marshal(tx.Outputs)
		fmt.Printf("SendCoins.Outputs %s \n", stx)
		bz, _ := json.Marshal(resAcc.Balance[1])
		fmt.Printf("Acc2 (20+(9) testcoin) %s \n", bz)

		resAcc1, err := s.GetAccount(pubKey1.Address())
		bz, _ = json.Marshal(resAcc1.Balance[1])
		fmt.Printf("Acc1 (200-(9+0) testcoin) %s \n", bz)

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

		amc2, _ := s.GetAMC(address)
		assert.NoError(t, err)
		fmt.Printf("\nGet amc: %v \n", amc2)

		tx.Outputs[0].Address = address
		tx.Inputs[0].Sequence = 2
		tx.Inputs[0].PubKey = nil

		signBytes = SignBytes(tx, s.GetChainID())
		sig, err = privKey1.Sign(signBytes)
		tx.Inputs[0].Signature = sig

		res = tx.Exec(s, nil, false)
		assert.Equal(t, types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("send"),
			Log:  "Ok",
		}, res)

		amc21, _ := s.GetAMC(address)

		fmt.Printf("Get amc: %v \n \n", amc21)

		resAcc, err = s.GetAccount(address)
		fmt.Printf("Get Account(amc_address): %v \n \n", resAcc)

		resAcc, err = s.GetAccount(s.ValProc.PullFee)
		fmt.Printf("Get Account(amc_address): %v \n \n", resAcc)

	})
}

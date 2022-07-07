package tx

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/tendermint/tendermint/abci/apps/otoncoin/state"
	"github.com/tendermint/tendermint/abci/types"
)

func (tx *MintCoins) Exec(s *state.State, hash []byte, check bool) types.ResponseDeliverTx {

	if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("validate inputs basic: %v", err)}
	}

	fees := Coins{tx.Fee}.ToStateCoins()
	if !fees.IsPositive() {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("Fee (%v) cannot be =< zero", fees)}
	}

	if !Inputs(tx.Inputs).TotalCoins().IsEqual(Coins{tx.Fee}.ToStateCoins()) {
		return types.ResponseDeliverTx{
			Code: CTErrInputsCoinsNotEqualFee,
			Log:  fmt.Sprintf("inputs total coins not equal to fee"),
		}
	}

	inTotal := Coins{tx.Fee}.ToStateCoins()
	if !inTotal.IsMovable(s) {
		return types.ResponseDeliverTx{
			Code: CTErrInputsNoMovable,
			Log:  fmt.Sprintf("in coin no movable")}
	}

	if err := Outputs(tx.Outputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateOutputsBasic,
			Log:  fmt.Sprintf("validate outputs basic: %v", err)}
	}

	if !(state.Coins{state.Coin{Name: tx.Name, Amount: tx.Amount}}.IsEqual(Outputs(tx.Outputs).TotalCoins())) {
		return types.ResponseDeliverTx{
			Code: CTErrValidateOutputsBasic,
			Log:  fmt.Sprintf("output total coins not equal to mint amount")}
	}

	accounts, err := Inputs(tx.Inputs).Accounts(s)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("get accounts, %v", err)}
	}

	signBytes := SignBytes(tx, s.GetChainID())
	if _, err := Inputs(tx.Inputs).ValidateAdvanced(accounts, signBytes); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsAdvanced,
			Log:  fmt.Sprintf("validate inputs advanced, %v", err),
			Info: fmt.Sprintf("%v", errors.Unwrap(err))}
	}

	accounts, err = Outputs(tx.Outputs).AddToAccountsMap(s, accounts)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetOrMakeOutputs,
			Log:  fmt.Sprintf("add to accounts map: %v", err)}
	}

	if err := tx.ValidateNewCoins(s, accounts); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateNewCoins,
			Log:  fmt.Sprintf("validate new coins: %v", err),
		}
	}

	if !bytes.Equal(s.ValProc.Minter, tx.Inputs[0].Address) {
		return types.ResponseDeliverTx{
			Code: TypeErrMinterOuner,
			Log:  fmt.Sprintf("Error minter ouner: %x", tx.Inputs[0].Address)}
	}

	if check {
		return types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("check"),
			Log:  fmt.Sprintf("Ok")}
	}

	if err := tx.AdjustCoins(s, accounts, hash); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrAdjustCoins,
			Log:  fmt.Sprintf("adjust coins: %v", err),
		}
	}

	if err := Inputs(tx.Inputs).AdjustAccounts(s, accounts); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("adjust input accounts: %v", err),
		}
	}

	if err := Outputs(tx.Outputs).AdjustAccounts(s, accounts); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("adjust output accounts: %v", err),
		}
	}

	if err := CoinToPullFee(s, fees); err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrOutput,
			Log:  fmt.Sprintf("adjust fee accounts: %v", err),
		}
	}

	return types.ResponseDeliverTx{
		Code: TypeOK,
		Data: []byte("mint"),
		Log:  "Ok",
	}
}

func (tx *MintCoins) ValidateNewCoins(s *state.State, accounts map[string]*state.Account) error {
	minter := accounts[string(tx.Inputs[0].Address)]
	newCoins := Outputs(tx.Outputs).TotalCoins()

	for _, c := range newCoins {
		coin, err := s.GetCoin(c.Name)

		if err != nil {
			return fmt.Errorf("get coin: %v", err)
		}
		if coin != nil && coin.CreatorAddr.String() != minter.PubKey.Address().String() {
			return fmt.Errorf("minter must be the same as coin creator in registry")
		}
	}
	return nil
}

func (tx *MintCoins) AdjustCoins(s *state.State, accounts map[string]*state.Account, hash []byte) error {
	minter := accounts[string(tx.Inputs[0].Address)]
	newCoins := Outputs(tx.Outputs).TotalCoins()
	for _, c := range newCoins {
		coin, err := s.GetCoin(c.Name)
		if err != nil {
			return fmt.Errorf("get coin: %v", err)
		}
		if coin != nil {
			coin.Amount += c.Amount
		} else {
			coin = &state.Coin{
				Name:         c.Name,
				Amount:       c.Amount,
				CreatorAddr:  minter.PubKey.Address(),
				DecimalPoint: tx.DecimalPoint,
				Movable:      tx.Movable,
				Hash:         hash,
				Delta:        tx.Delta,
			}
		}
		fmt.Printf("AdjustCoins: %v \n", coin)

		if err := s.SetCoin(coin); err != nil {
			return fmt.Errorf("set coin: %v", err)
		}
	}
	return nil
}

func (tx *SetSalePrice) Exec(s *state.State) types.ResponseDeliverTx {
	if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("validate inputs basic: %v", err)}
	}

	if !Inputs(tx.Inputs).TotalCoins().IsGTE(Coins{tx.Fee}.ToStateCoins()) {
		return types.ResponseDeliverTx{
			Code: CTErrInputsCoinsNotEqualFee,
			Log:  fmt.Sprintf("inputs total coins not equal to fee"),
		}
	}

	return types.ResponseDeliverTx{
		Code: TypeOK,
		Log:  "Ok",
	}
}

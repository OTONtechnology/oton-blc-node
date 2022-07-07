package tx

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/tendermint/tendermint/abci/apps/otoncoin/state"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	OnNewPrizHight = 12610
	OnFixDelVal    = 23422
)

func (tx *SetInteresNode) Exec(s *state.State, _ []byte, check bool) types.ResponseDeliverTx {
	var txe SetNodeReward

	txe.Fee = tx.Fee
	txe.Value = tx.Value
	txe.Inputs = tx.Inputs

	fmt.Printf("SetInteresNode(), %v \n", txe.Inputs)
	return txe.Exec(s, nil, check)
}

func (tx *SetNodeReward) Exec(s *state.State, _ []byte, check bool) types.ResponseDeliverTx {

	if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("in validateInputsBasic(): %v", err)}
	}

	accounts, err := Inputs(tx.Inputs).Accounts(s)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("in getInputs(), %v", err)}
	}

	if !Inputs(tx.Inputs).TotalCoins().IsEqual(Coins{tx.Fee}.ToStateCoins()) {
		return types.ResponseDeliverTx{
			Code: CTErrInputsCoinsNotEqualFee,
			Log:  fmt.Sprintf("inputs total coins not equal to fee"),
		}
	}

	signBytes := SignBytes(tx, s.GetChainID())
	inTotal, err := Inputs(tx.Inputs).ValidateAdvanced(accounts, signBytes)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsAdvanced,
			Log:  fmt.Sprintf("validate inputs advanced, %v", err),
			Info: fmt.Sprintf("%v", errors.Unwrap(err))}
	}

	if !inTotal.IsMovable(s) {
		return types.ResponseDeliverTx{
			Code: CTErrInputsNoMovable,
			Log:  fmt.Sprintf("in coin no movable()")}
	}

	if !bytes.Equal(s.ValProc.Ouner, tx.Inputs[0].Address) {
		return types.ResponseDeliverTx{
			Code: TypeErrValOuner,
			Log:  fmt.Sprintf("Error gift ouner: %x", tx.Inputs[0].Address)}
	}

	if check {
		return types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("check"),
			Log:  fmt.Sprintf("Ok")}
	}

	if err := Inputs(tx.Inputs).AdjustAccounts(s, accounts); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("adjust input accounts: %v", err),
		}
	}

	fees := Coins{tx.Fee}.ToStateCoins()
	if err := CoinToPullFee(s, fees); err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrOutput,
			Log:  fmt.Sprintf("adjust fee accounts: %v", err),
		}
	}

	s.ValProc.Gift = Coins{tx.Value}.ToStateCoins()

	return types.ResponseDeliverTx{
		Code: TypeOK,
		Data: []byte("dozefee"),
		Log:  fmt.Sprintf("Ok")}
}

func (tx *BuyNode) Exec(s *state.State, _ []byte, check bool) types.ResponseDeliverTx {

	fmt.Printf("Buy node TX...\n")
	if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("in validateInputsBasic(): %v", err)}
	}

	if err := Outputs(tx.Outputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateOutputsBasic,
			Log:  fmt.Sprintf("in validateOutputsBasic(): %v", err)}
	}

	accounts, err := Inputs(tx.Inputs).Accounts(s)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("in getInputs(), %v", err)}
	}

	signBytes := SignBytes(tx, s.GetChainID())

	inTotal, res := Inputs(tx.Inputs).ValidateAdvanced(accounts, signBytes)
	if res != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsAdvanced,
			Log:  fmt.Sprintf("in validateInputsAdvanced(), %v", res),
			Info: fmt.Sprintf("%v", errors.Unwrap(res))}
	}

	if !inTotal.IsMovable(s) {
		return types.ResponseDeliverTx{
			Code: CTErrInputsNoMovable,
			Log:  fmt.Sprintf("in coin no movable()")}
	}

	accounts, err = Outputs(tx.Outputs).AddToAccountsMap(s, accounts)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetOrMakeOutputs,
			Log:  fmt.Sprintf("in getOrMakeOutputs()")}
	}

	outTotal := Outputs(tx.Outputs).TotalCoins()
	outPlusFees := outTotal
	fees := Coins{tx.Fee}.ToStateCoins()

	if !fees.IsPositive() {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("Fee (%v) cannot be =< zero", fees)}
	}
	if fees.IsValid() {
		outPlusFees = outTotal.Plus(fees)
	}

	if !inTotal.IsEqual(outPlusFees) {
		return types.ResponseDeliverTx{
			Code: CTErrBaseInvalidOutput,
			Log:  fmt.Sprintf("Input total (%v) != output total + fees (%v)", inTotal, outPlusFees)}
	}

	if tx.Validator == nil || tx.Power < 0 || len(tx.Validator) != 32 {
		return types.ResponseDeliverTx{
			Code: TypeEncodingError,
			Log:  fmt.Sprintf("Validator error %v, power %d", tx.Validator, tx.Power)}
	}

	if !bytes.Equal(s.ValProc.Ouner, tx.Inputs[0].Address) {
		return types.ResponseDeliverTx{
			Code: TypeErrValOuner,
			Log:  fmt.Sprintf("Error reg val ouner: %x", tx.Inputs[0].Address)}
	}

	if check {
		return types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("check"),
			Log:  fmt.Sprintf("Ok")}
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

	var pubkey ed25519.PubKey
	pubkey = tx.Validator
	candidat := s.ValProc.ValMap[string(pubkey.Address())]

	fmt.Printf("candidat.Holder :  %s , len=%v\n", candidat.Holder, len(tx.Holder))

	if len(tx.Holder) == 20 {
		candidat.Holder = tx.Holder
	}

	if len(candidat.Holder) < 20 {
		candidat.Holder = pubkey.Address()
	}

	fmt.Printf("candidat.Holder :  %s , len=%v\n", candidat.Holder, len(tx.Holder))
	fmt.Printf("candidat.Address :  %v \n", pubkey.Address())

	if tx.Value.Name == state.CoinDepo && (state.HeightLine > OnNewPrizHight) {

		Deposit := Coins{tx.Value}.ToStateCoins()
		candidat.Contract = candidat.Contract.Plus(Deposit)

		fmt.Printf("add_depo:  %s \n", candidat.Contract)
	}

	if tx.Value.Name == state.CoinDepo && (state.HeightLine < OnNewPrizHight) && tx.Value.Amount != 2 {
		Deposit := Coins{tx.Value}.ToStateCoins()
		candidat.Contract = candidat.Contract.Plus(Deposit)
		fmt.Printf("old_add_depo:  %s \n", candidat.Contract)
	}

	if tx.Value.Name == state.CoinDepo && tx.Value.Amount == 2 {
		fmt.Printf("skip_depo:  %s \n", candidat.Contract)
	}

	if tx.Value.Name == state.CoinContract {
		Contract := Coins{tx.Value}.ToStateCoins()
		Deposit := candidat.Contract.Coin(state.CoinDepo)

		if candidat.Contract.Coin(state.CoinDepo) != nil {
			candidat.Contract = Contract
			candidat.Contract = append(candidat.Contract, *Deposit)

		} else {

			candidat.Contract = Contract

			candidat.Priz = state.Coins{state.NewCoin("oton", 1)}
		}
	}

	candidat.Power = tx.Power
	candidat.Delta = Coins{tx.Delta}.ToStateCoins()

	candidat.TimePay = state.TimeWorkingBlock

	ref := &state.Coin{}
	fraction := float64(100.0)

	if candidat.Contract.Coin(state.CoinContract) != nil {
		switch candidat.Contract.Coin(state.CoinContract).Amount {
		case 1:
			fraction = float64(0.25) / fraction
			ref = &state.Coin{Name: state.CoinContract, Amount: 0, Expired: state.TimeWorkingBlock.AddDate(5, 0, 0)}
		case 2:
			fraction = float64(1) / fraction
			ref = &state.Coin{Name: state.CoinContract, Amount: 0, Expired: state.TimeWorkingBlock.AddDate(1, 0, 0)}
		case 3:
			fraction = float64(1.5) / fraction
			ref = &state.Coin{Name: state.CoinContract, Amount: 0, Expired: state.TimeWorkingBlock.AddDate(3, 0, 0)}
		case 4:
			fraction = float64(2) / fraction
			ref = &state.Coin{Name: state.CoinContract, Amount: 0, Expired: state.TimeWorkingBlock.AddDate(5, 0, 0)}
		default:
			fraction = float64(2) / fraction
			ref = &state.Coin{Name: state.CoinContract, Amount: 0, Expired: state.TimeWorkingBlock.AddDate(5, 0, 0)}
		}

		candidat.Contract = candidat.Contract.Plus(state.Coins{*ref})
	} else {

		fraction = float64(2) / fraction
	}

	if candidat.Contract.Coin(state.CoinDepo) != nil {

		new_priz := state.Coin{Name: candidat.Delta[0].Name, Amount: candidat.Contract.Coin(state.CoinDepo).Fraction(fraction).Amount}
		fmt.Printf("new_priz: %v , old %v \n", new_priz, candidat.Priz)

		candidat.Priz = state.Coins{new_priz}
		fmt.Printf("candidat_priz:  %v \n", candidat.Priz)
		fmt.Printf("candidat_coin:  %s \n", candidat.Contract)
	}

	s.ValProc.ValMap[string(pubkey.Address())] = candidat

	return types.ResponseDeliverTx{
		Code:      TypeOK,
		Data:      pubkey.Bytes(),
		GasWanted: tx.Power,
		Log:       fmt.Sprintf("Ok")}
}

func (tx *SendCoins) Coinbase(s *state.State, check bool) types.ResponseDeliverTx {
	if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("in validateInputsBasic(): %v", err)}
	}

	if err := Outputs(tx.Outputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateOutputsBasic,
			Log:  fmt.Sprintf("in validateOutputsBasic(): %v", err)}
	}

	accounts, err := Inputs(tx.Inputs).Accounts(s)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("in getInputs(), %v", err)}
	}

	inTotal := Inputs(tx.Inputs).TotalCoins()
	if !inTotal.IsMovable(s) {
		return types.ResponseDeliverTx{
			Code: CTErrInputsNoMovable,
			Log:  fmt.Sprintf("in coin no movable()")}
	}

	accounts, err = Outputs(tx.Outputs).AddToAccountsMap(s, accounts)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetOrMakeOutputs,
			Log:  fmt.Sprintf("in getOrMakeOutputs()")}
	}

	outTotal := Outputs(tx.Outputs).TotalCoins()
	outPlusFees := outTotal
	fees := Coins{tx.Fee}.ToStateCoins()

	if fees.IsValid() {
		outPlusFees = outTotal.Plus(fees)
	}

	if !inTotal.IsEqual(outPlusFees) {
		return types.ResponseDeliverTx{
			Code: CTErrBaseInvalidOutput,
			Log:  fmt.Sprintf("Input total (%v) != output total + fees (%v)", inTotal, outPlusFees)}
	}

	if check {
		return types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("check"),
			Log:  fmt.Sprintf("Ok")}
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

	return types.ResponseDeliverTx{
		Code: TypeOK,
		Data: []byte("coinbase"),
		Log:  fmt.Sprintf("Ok")}
}

func CoinToPullFee(s *state.State, fee state.Coins) error {

	fmt.Printf("Coin to PullFee: %v \n", s.ValProc.PullFee)

	acc, err := s.GetAccount(s.ValProc.PullFee)
	if err != nil {
		return fmt.Errorf("get fee account: %v", err)
	}

	if acc == nil {
		acc = &state.Account{}
	}

	acc.Address = s.ValProc.PullFee
	acc.Balance = acc.Balance.Plus(fee)
	if err := s.SetAccount(s.ValProc.PullFee, acc); err != nil {
		return fmt.Errorf("set fee account: %v", err)
	}

	return nil
}

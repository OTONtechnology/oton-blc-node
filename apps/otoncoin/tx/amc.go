package tx

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/tendermint/tendermint/abci/apps/otoncoin/state"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/json"
)

func (tx *CreateAMC) Exec(s *state.State, _ []byte, check bool) types.ResponseDeliverTx {

	if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("validate inputs basic: %v", err)}
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

	if len(tx.Address) != 20 {
		return types.ResponseDeliverTx{
			Code: TypeErrGetAMC,
			Log:  fmt.Sprintf("Invalid AMC address length:, %x", tx.Address)}
	}

	fmt.Printf("Create AMC with address: %x \n\n", tx.Address)
	amc, err := s.GetAMC(tx.Address)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrGetAMC,
			Log:  fmt.Sprintf("Error get AMC:, %v", err)}
	}

	if amc != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrDubAMC,
			Log:  fmt.Sprintf("Error dublicate for create AMC")}
	}

	adr_amc, err := s.GetAccount(tx.Address)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrGetAMC,
			Log:  fmt.Sprintf("Error get address ACC-AMC:, %v", err)}
	}

	if adr_amc != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrDubAMC,
			Log:  fmt.Sprintf("Error dublicate for create AMC")}
	}

	inTotal = inTotal.Minus(Coins{tx.Fee}.ToStateCoins())
	amc, err = state.NewAMC(
		tx.Address, tx.Name, tx.Inputs[0].Address, inTotal, AMCParams(tx.Params).ToStateParams(),
	)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrSetAMC,
			Log:  fmt.Sprintf("Error new AMC: %v", err)}
	}

	if check {
		return types.ResponseDeliverTx{
			Code: TypeOK,
			Data: []byte("check"),
			Log:  fmt.Sprintf("Ok")}
	}

	if err := s.SetAMC(amc); err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrSetAMC,
			Log:  fmt.Sprintf("Error set AMC: %v", err)}
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

	return types.ResponseDeliverTx{
		Code: TypeOK,
		Data: []byte("CreateAMC"),
		Log:  fmt.Sprintf("Ok")}
}

func (tx *SetInAMC) Exec(s *state.State, _ []byte, check bool) types.ResponseDeliverTx {

	if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("validate inputs basic: %v", err)}
	}

	if (len(tx.Address) + len(tx.Referal) + len(tx.Sponsor)) != 60 {
		return types.ResponseDeliverTx{
			Code: TypeErrGetAMC,
			Log:  fmt.Sprintf("Invalid param address length: %x, %x, %x", tx.Address, tx.Referal, tx.Sponsor)}
	}

	accounts, err := Inputs(tx.Inputs).Accounts(s)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("get inputs accounts, %v", err)}
	}

	spn, err := s.GetAccount(tx.Sponsor)
	if spn == nil || err != nil {

		return types.ResponseDeliverTx{
			Code: TypeErrSponsorAMC,
			Log:  fmt.Sprintf("Error sponsor for AMC")}
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

	amc, err := s.GetAMC(tx.Address)
	if err != nil || amc == nil {
		return types.ResponseDeliverTx{
			Code: TypeErrGetAMC,
			Log:  fmt.Sprintf("Error get AMC: %v", err)}
	}

	acc, _ := s.GetAccount(tx.Referal)

	if acc == nil {
		acc = &state.Account{}

		amct, err := s.GetAMC(tx.Referal)
		if amct != nil {
			return types.ResponseDeliverTx{
				Code: TypeErrReferalAMC,
				Log:  fmt.Sprintf("Error referal for AMC: %v", err)}
		}

	}

	if acc.Up == nil {
		acc.Up = tx.Sponsor

		if !check {

			fmt.Printf("tx.Sponsor: %x \n", tx.Sponsor)
			fmt.Printf("tx.Referal: %x \n", tx.Referal)
			fmt.Printf("acc.Referal: %v \n", acc)
		}
	} else {

		return types.ResponseDeliverTx{
			Code: TypeErrSponsorAMC,
			Log:  fmt.Sprintf("Error sponsor AMC: %v", err)}
	}

	rc, err := s.RefChain(tx.Sponsor)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrRefChain,
			Log:  fmt.Sprintf("Error referal chain for AMC: %v", err)}
	}

	if !bytes.Equal(amc.Creator, rc.Root().PubKey.Address()) {
		return types.ResponseDeliverTx{
			Code: TypeErrTreeAMC,
			Log:  fmt.Sprintf("Error tree for AMC")}
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

	rc, _ = s.RefChain(tx.Sponsor)

	fmt.Printf("rc[0] : %v \n", rc[0])
	rc[0].Down = append(rc[0].Down, tx.Referal)
	acc.Up = tx.Sponsor
	acc.Address = tx.Referal

	var hk state.RefChain
	hk = append(hk, acc)
	rc = append(hk, rc...)

	fmt.Printf("rc[a] : %v \n", rc[0])
	fmt.Printf("rc[1] : %v \n \n", rc[1])

	if err := s.SetRefChain(rc); err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrRefChain,
			Log:  fmt.Sprintf("Error set referal chain: %v", err)}
	}

	fees := Coins{tx.Fee}.ToStateCoins()
	if err := CoinToPullFee(s, fees); err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrOutput,
			Log:  fmt.Sprintf("adjust fee accounts: %v", err),
		}
	}

	return types.ResponseDeliverTx{
		Code: TypeOK,
		Data: []byte("SetInAMC"),
		Log:  fmt.Sprintf("Ok")}
}

func (tx *BuyInAMC) Exec(s *state.State, _ []byte, check bool) types.ResponseDeliverTx {

	if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("validate inputs basic: %v", err)}
	}

	accounts, err := Inputs(tx.Inputs).Accounts(s)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: CTErrGetInputs,
			Log:  fmt.Sprintf("get inputs accounts, %v", err)}
	}

	value := Coins{tx.Value}.ToStateCoins()
	delta := Coins{tx.Delta}.ToStateCoins()
	fees := Coins{tx.Fee}.ToStateCoins()

	if !(fees.IsPositive() && value.IsPositive() && delta.IsPositive()) {
		return types.ResponseDeliverTx{
			Code: CTErrValidateInputsBasic,
			Log:  fmt.Sprintf("Fee (%v), Value (%v), Delta (%v)  cannot be =< zero", fees, value, value)}
	}

	if !value.IsMovable(s) {
		value = state.Coins{state.Coin{Name: tx.Delta.Name, Amount: tx.Value.Amount}}
		fmt.Printf("now tarif coin IsMovable %v \n", value)
	}

	valuePlusFees := value.Plus(Coins{tx.Fee}.ToStateCoins())
	if !Inputs(tx.Inputs).TotalCoins().IsEqual(valuePlusFees) {
		return types.ResponseDeliverTx{
			Code: CTErrInputsCoinsNotEqualFee,
			Log:  fmt.Sprintf("inputs total coins not equal to fee + value"),
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

	amc, err := s.GetAMC(tx.Address)
	if err != nil || amc == nil {
		return types.ResponseDeliverTx{
			Code: TypeErrGetAMC,
			Log:  fmt.Sprintf("Error get AMC: %v", err)}
	}

	amc_creator := accounts[string(tx.Inputs[0].Address)]
	if amc.Creator.String() != amc_creator.PubKey.Address().String() {
		return types.ResponseDeliverTx{
			Code: TypeErrMasterAMC,
			Log:  fmt.Sprintf("Error master AMC: %v", err)}
	}

	acc, _ := s.GetAccount(tx.Referal)
	if acc == nil {
		return types.ResponseDeliverTx{
			Code: TypeErrReferalAMC,
			Log:  fmt.Sprintf("Error referal AMC: %v", err)}
	}
	if acc.Up == nil {
		return types.ResponseDeliverTx{
			Code: TypeErrSponsorAMC,
			Log:  fmt.Sprintf("Error sponsor AMC: %v", err)}
	}

	rc, err := s.RefChain(tx.Referal)
	if err != nil {

		return types.ResponseDeliverTx{
			Code: TypeErrMasterAMC,
			Log:  fmt.Sprintf("Error chain referal in AMC: %v", err)}
	}

	if !bytes.Equal(amc.Creator, rc.Root().PubKey.Address()) {
		return types.ResponseDeliverTx{
			Code: TypeErrTreeAMC,
			Log:  fmt.Sprintf("Error tree for AMC")}
	}

	rf, err := s.RefRank(rc)
	if err != nil {
		fmt.Printf("RefRankErr = rf %v \n \n", err)
		return types.ResponseDeliverTx{
			Code: TypeErrMasterAMC,
			Log:  fmt.Sprintf("Error dw referal in AMC: %v", err)}
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

	rc, _ = s.RefChain(tx.Referal)
	rf, _ = s.RefRank(rc)
	rl := state.RefChain{}

	rc, rl, err = amc.MakeRank(rc, rf, Coins{tx.Value}.ToStateCoins())

	fmt.Printf("BuyInAMC after check_tx \n")
	fmt.Printf("RefChain = rf %v \n \n", rc)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrSetAMC,
			Log:  fmt.Sprintf("Error make rank:, %v", err)}
	}

	amc.Balance = amc.Balance.Plus(value)
	amc.Balance = amc.Balance.Minus(delta)
	payment, err := amc.MakePayment(rc, Coins{tx.Delta}.ToStateCoins())
	fmt.Printf("Payment:  %v \n \n", payment)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrSetAMC,
			Log:  fmt.Sprintf("Error make payments:, %v", err)}
	}

	if err := s.SetAMC(amc); err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrSetAMC,
			Log:  fmt.Sprintf("Error set AMC:, %v", err)}
	}

	if err := s.SetRefChain(rc); err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrRefChain,
			Log:  fmt.Sprintf("Error set referal chain: %v", err)}
	}

	amctx, _ := WrapEvents(payment, []byte(tx.Address))
	bz_pay, _ := json.Marshal(amctx)

	_, ranktx := WrapEvents(rl, []byte(tx.Address))
	bz_rank, _ := json.Marshal(ranktx)
	fmt.Printf("bzo: %s \n \n \n", bz_rank)

	events := []types.Event{
		{
			Type: "app",
			Attributes: []types.EventAttribute{
				{Key: "payment", Value: string(bz_pay), Index: true},
				{Key: "ranks", Value: string(bz_rank), Index: true},
			},
		},
	}

	if err := CoinToPullFee(s, fees); err != nil {
		return types.ResponseDeliverTx{
			Code: TypeErrOutput,
			Log:  fmt.Sprintf("adjust fee accounts: %v", err),
		}
	}

	return types.ResponseDeliverTx{
		Code:   TypeOK,
		Data:   []byte("BuyInAMC"),
		Events: events,
		Log:    fmt.Sprintf("Ok")}
}

func WrapEvents(refChain state.RefChain, address []byte) (SendCoins, OutputsPr) {
	var adr []byte
	var out Outputs
	var rank OutputsPr

	totalPayments := state.Coins{}

	for _, acc := range refChain {

		if acc.PubKey == nil {
			adr = acc.Address
		} else {
			adr = acc.PubKey.Address()
		}

		if adr != nil {
			coins := FromStateCoins(acc.Balance)
			ot := Output{Address: adr, Coins: coins}
			out = append(out, ot)

			coinsPr := FromStateCoinsPr(acc.Balance)
			otrank := OutputPr{Address: adr, Coins: coinsPr}
			rank = append(rank, otrank)

			totalPayments = totalPayments.Plus(acc.Balance)
		}
	}

	tc := FromStateCoins(totalPayments)
	tx := SendCoins{
		Fee:     Coin{"", 0},
		Inputs:  Inputs{Input{Address: address, Coins: tc}},
		Outputs: out,
	}

	if len(tc) == 0 {
		tx.Inputs = nil
	}
	return tx, rank
}

type AMCParams []AMCParam

func (params AMCParams) ToStateParams() []state.AMCParam {
	var res []state.AMCParam
	for _, p := range params {
		res = append(res, state.AMCParam{Key: p.Key, Value: p.Value})
	}
	return res
}

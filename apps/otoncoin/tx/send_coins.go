package tx

import (
 "errors"
 "fmt"

 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
 "github.com/tendermint/tendermint/abci/types"
)

func (tx *SendCoins) Exec(s *state.State, _ []byte, check bool) types.ResponseDeliverTx {
 if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrValidateInputsBasic,
   Log: fmt.Sprintf("in validateInputsBasic(): %v", err)}
 }

 if err := Outputs(tx.Outputs).ValidateBasic(); err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrValidateOutputsBasic,
   Log: fmt.Sprintf("in validateOutputsBasic(): %v", err)}
 }


 accounts, err := Inputs(tx.Inputs).Accounts(s)
 if err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrGetInputs,
   Log: fmt.Sprintf("in getInputs(), %v", err)}
 }


 signBytes := SignBytes(tx, s.GetChainID())

 inTotal, res := Inputs(tx.Inputs).ValidateAdvanced(accounts, signBytes)
 if res != nil {
  return types.ResponseDeliverTx{
   Code: CTErrValidateInputsAdvanced,
   Log: fmt.Sprintf("in validateInputsAdvanced(), %v", res),
   Info: fmt.Sprintf("%v", errors.Unwrap(res))}
 }


 if !inTotal.IsMovable(s) {
  return types.ResponseDeliverTx{
   Code: CTErrInputsNoMovable,
   Log: fmt.Sprintf("in coin no movable()")}
 }



 accounts, err = Outputs(tx.Outputs).AddToAccountsMap(s, accounts)
 if err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrGetOrMakeOutputs,
   Log: fmt.Sprintf("in getOrMakeOutputs()")}
 }


 outTotal := Outputs(tx.Outputs).TotalCoins()
 outPlusFees := outTotal
 fees := Coins{tx.Fee}.ToStateCoins()


 if !fees.IsPositive() {
  return types.ResponseDeliverTx{
   Code: CTErrValidateInputsBasic,
   Log: fmt.Sprintf("Fee (%v) cannot be =< zero", fees)}
 }
 if fees.IsValid() {
  outPlusFees = outTotal.Plus(fees)
 }

 if !inTotal.IsEqual(outPlusFees) {
  return types.ResponseDeliverTx{
   Code: CTErrBaseInvalidOutput,
   Log: fmt.Sprintf("Input total (%v) != output total + fees (%v)", inTotal, outPlusFees)}
 }


 if check {
  return types.ResponseDeliverTx{
   Code: TypeOK,
   Data: []byte("check"),
   Log: fmt.Sprintf("Ok")}
 }


 if err := Inputs(tx.Inputs).AdjustAccounts(s, accounts); err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrGetInputs,
   Log: fmt.Sprintf("adjust input accounts: %v", err),
  }
 }
 if err := Outputs(tx.Outputs).AdjustAccounts(s, accounts); err != nil {
  return types.ResponseDeliverTx{
   Code: TypeErrOutput,
   Log: fmt.Sprintf("adjust output accounts: %v", err),
  }
 }


 if err := CoinToPullFee(s, fees); err != nil {
  return types.ResponseDeliverTx{
   Code: TypeErrOutput,
   Log: fmt.Sprintf("adjust fee accounts: %v", err),
  }
 }

 return types.ResponseDeliverTx{
  Code: TypeOK,
  Data: []byte("send"),
  Log: fmt.Sprintf("Ok")}
}

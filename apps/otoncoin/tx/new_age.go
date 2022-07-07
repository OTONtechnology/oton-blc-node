package tx

import (
 "bytes"
 "errors"
 "fmt"

 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
 "github.com/tendermint/tendermint/abci/types"
)

func (tx *SetNewAge) Exec(s *state.State, hash []byte, check bool) types.ResponseDeliverTx {

 if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrValidateInputsBasic,
   Log: fmt.Sprintf("validate inputs basic: %v", err)}
 }

 fees := Coins{tx.Fee}.ToStateCoins()
 if !fees.IsPositive() {
  return types.ResponseDeliverTx{
   Code: CTErrValidateInputsBasic,
   Log: fmt.Sprintf("Fee (%v) cannot be =< zero", fees)}
 }



 if !Inputs(tx.Inputs).TotalCoins().IsEqual(Coins{tx.Fee}.ToStateCoins()) {
  return types.ResponseDeliverTx{
   Code: CTErrInputsCoinsNotEqualFee,
   Log: fmt.Sprintf("inputs total coins not equal to fee"),
  }
 }


 inTotal := Coins{tx.Fee}.ToStateCoins()
 if !inTotal.IsMovable(s) {
  return types.ResponseDeliverTx{
   Code: CTErrInputsNoMovable,
   Log: fmt.Sprintf("in coin no movable")}
 }


 accounts, err := Inputs(tx.Inputs).Accounts(s)
 if err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrGetInputs,
   Log: fmt.Sprintf("get accounts, %v", err)}
 }


 signBytes := SignBytes(tx, s.GetChainID())
 if _, err := Inputs(tx.Inputs).ValidateAdvanced(accounts, signBytes); err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrValidateInputsAdvanced,
   Log: fmt.Sprintf("validate inputs advanced, %v", err),
   Info: fmt.Sprintf("%v", errors.Unwrap(err))}
 }


 if !bytes.Equal(s.ValProc.Minter, tx.Inputs[0].Address) {
  return types.ResponseDeliverTx{
   Code: TypeErrAgeOuner,
   Log: fmt.Sprintf("Error age ouner: %x", tx.Inputs[0].Address)}
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


 if err := CoinToPullFee(s, fees); err != nil {
  return types.ResponseDeliverTx{
   Code: TypeErrOutput,
   Log: fmt.Sprintf("adjust fee accounts: %v", err),
  }
 }




 newage := AgeParams(tx.Params).ToStateAge()
 up_age := s.Age.UpdateAge(newage)


 err = s.Age.WorkAge(up_age)
 if err != nil {
  fmt.Printf("Age update load err: %e \n", err)
 }

 return types.ResponseDeliverTx{
  Code: TypeOK,
  Data: []byte("newage"),
  Log: "Ok",
 }
}


type AgeParams []AgeParam


func (params AgeParams) ToStateAge() []state.AgeParam {
 var res []state.AgeParam
 for _, p := range params {
  if p.Key != "" {
   res = append(res, state.AgeParam{Key: p.Key, Value: p.Value, Text: p.Svalue})
  }
 }
 return res
}

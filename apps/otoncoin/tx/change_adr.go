package tx

import (
 "bytes"
 "errors"
 "fmt"

 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
 "github.com/tendermint/tendermint/abci/types"
 "github.com/tendermint/tendermint/crypto"
)



func RepairKey(hash []byte) []byte {
 return append([]byte("base/repadr/"), hash...)
}



func (tx *ChangeAddressWant) Exec(s *state.State, hash []byte, check bool) types.ResponseDeliverTx {


 txs := &SendCoins{
  Fee: tx.Fee,
  Inputs: tx.Inputs,
  Outputs: tx.Outputs,
 }



 rdl := txs.Exec(s, hash, check)

 if (rdl.Code == TypeOK) && !check {


  s.SetValue(RepairKey(tx.Inputs[0].Address), tx)
  rdl.Data = []byte("replace_adr_want")
 }
 return rdl
}


func (tx *ChangeAddressApply) Exec(s *state.State, hash []byte, check bool) types.ResponseDeliverTx {

 if err := Inputs(tx.Inputs).ValidateBasic(); err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrValidateInputsBasic,
   Log: fmt.Sprintf("in validateInputsBasic(): %v", err)}
 }

 if err := Votes(tx.Votes).ValidateBasicVotes(); err != nil {
  return types.ResponseDeliverTx{
   Code: TypeErrVoteAdr,
   Log: fmt.Sprintf("in ValidateBasicVotes(): %v", err)}
 }


 accounts, err := Inputs(tx.Inputs).Accounts(s)
 if err != nil {
  return types.ResponseDeliverTx{
   Code: CTErrGetInputs,
   Log: fmt.Sprintf("in getInputs(), %v", err)}
 }



 if !Inputs(tx.Inputs).TotalCoins().IsEqual(Coins{tx.Fee}.ToStateCoins()) {
  return types.ResponseDeliverTx{
   Code: CTErrInputsCoinsNotEqualFee,
   Log: fmt.Sprintf("inputs total coins not equal to fee"),
  }
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


 old, err := s.GetAccount(tx.Address)
 if err != nil || old == nil {
  return types.ResponseDeliverTx{
   Code: TypeErrOldAddress,
   Log: fmt.Sprintf("get old account: %v", err)}
 }


 WantList := ChangeAddressWant{}
 err = s.GetValue(RepairKey(tx.Address), &WantList)
 if err != nil || len(WantList.Outputs) < 1 {

  return types.ResponseDeliverTx{
   Code: TypeErrWantList,
   Log: fmt.Sprintf("no list votes: %v", err)}
 }


 signBytesVotes := tx.SignBytesVotes(s.GetChainID())


 votes, err := Votes(tx.Votes).Accounts(s, signBytesVotes)
 if err != nil {
  return types.ResponseDeliverTx{
   Code: TypeErrVoteAdr,
   Log: fmt.Sprintf("in getVotes(), %v", err)}
 }


 if err := EqualVote(votes, WantList.Outputs); err != nil {
  return types.ResponseDeliverTx{
   Code: TypeErrVoteAdr,
   Log: fmt.Sprintf("vote: %v", err),
  }
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



 if err = SwapAddress(s, tx.Inputs[0].Address, tx.Address); err != nil {
  return types.ResponseDeliverTx{
   Code: TypeErrChangeAdr,
   Log: fmt.Sprintf("swap accounts: %v", err),
  }
 }


 fees := Coins{tx.Fee}.ToStateCoins()
 if err := CoinToPullFee(s, fees); err != nil {
  return types.ResponseDeliverTx{
   Code: TypeErrOutput,
   Log: fmt.Sprintf("adjust fee accounts: %v", err),
  }
 }

 return types.ResponseDeliverTx{
  Code: TypeOK,
  Data: []byte("replace_adr_apply"),
  Log: fmt.Sprintf("Ok")}
}


func EqualVote(vote Votes, list Outputs) error {
 vote_n := 0
 for _, l := range list {
  for _, v := range vote {
   if bytes.Compare(v.Address, l.Address) == 0 {
    vote_n++
    continue
   }
  }
 }

 if (len(list) / 2) <= vote_n {
  return nil
 } else {
  return fmt.Errorf("vote: %v", vote_n)
 }
}




func SwapAddress(s *state.State, newAdr crypto.Address, oldAdr crypto.Address) error {

 acc_new, _ := s.GetAccount(newAdr)
 acc_old, _ := s.GetAccount(oldAdr)


 acc_new.Balance = acc_new.Balance.Plus(acc_old.Balance)
 acc_old.Balance = state.Coins{}


 acc_new.Down = acc_old.Down
 acc_new.Up = acc_old.Up


 acc_old.Down = nil
 acc_old.Up = nil


 for _, adr := range acc_new.Down {
  acc_dw, err := s.GetAccount(adr)
  if err != nil || acc_dw == nil {
   return fmt.Errorf("get account dw: %v", err)
  }

  acc_dw.Up = newAdr

  if err := s.SetAccount(acc_dw.Address, acc_dw); err != nil {
   return fmt.Errorf("set account: %v", err)
  }
 }


 sponsor, err := s.GetAccount(acc_new.Up)
 if err != nil || sponsor == nil {
  return fmt.Errorf("get account sponsor: %v", err)
 }
 for i, adr := range sponsor.Down {
  if bytes.Compare(acc_old.Address, adr) == 0 {
   sponsor.Down[i] = acc_new.Address
  }
 }

 if err := s.SetAccount(sponsor.Address, sponsor); err != nil {
  return fmt.Errorf("set account: %v", err)
 }


 if err := s.SetAccount(acc_new.Address, acc_new); err != nil {
  return fmt.Errorf("set account: %v", err)
 }


 if err := s.SetAccount(acc_old.Address, acc_old); err != nil {
  return fmt.Errorf("set account: %v", err)
 }

 return nil
}



func (tx *ChangeAddressApply) SignBytesVotes(chainID string) []byte {

 insCopy := make(Inputs, len(tx.Inputs))
 copy(insCopy, tx.Inputs)

 voteCopy := make(Votes, len(tx.Votes))
 copy(voteCopy, tx.Votes)

 txo := &ChangeAddressApply{
  Fee: tx.Fee,
  Inputs: insCopy,
  Votes: voteCopy,
  Address: tx.Address,
 }


 for i := range txo.Inputs {
  txo.Inputs[i].Signature = Signature{}
 }


 for i := range txo.Votes {
  txo.Votes[i].Signature = Signature{}
  txo.Votes[i].PubKey = Signature{}
 }


 signBytesNext, _ := tx.Marshal()
 signBytes := []byte(chainID)
 signBytes = append(signBytes, signBytesNext...)


 return signBytes
}

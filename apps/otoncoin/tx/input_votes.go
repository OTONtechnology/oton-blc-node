package tx

import (
 "fmt"

 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
 "github.com/tendermint/tendermint/crypto/ed25519"
)



func (in Vote) ValidateBasicVote() error {
 if len(in.Address) != 20 {
  return fmt.Errorf("invalid address length")
 }
 if in.PubKey == nil {
  return fmt.Errorf("PubKey must be present ")
 }
 return nil
}


type Votes []Vote


func (ins Votes) ValidateBasicVotes() error {
 if len(ins) == 0 {
  return fmt.Errorf("should contain at least one input")
 }
 for _, in := range ins {

  if err := in.ValidateBasicVote(); err != nil {
   return err
  }
 }
 return nil
}


func (ins Votes) Accounts(s *state.State, signBytes []byte) (Votes, error) {
 res := map[string]*state.Account{}
 vt := Votes{}

 for _, in := range ins {

  if _, ok := res[string(in.Address)]; ok {
   return nil, fmt.Errorf("ErrBaseDuplicateAddress: %x", in.Address)
  }


  account, err := s.GetAccount(in.Address)
  if err != nil {
   continue

  }
  if account == nil {
   continue

  }

  if in.PubKey != nil {
   account.PubKey = ed25519.PubKey(in.PubKey[:32])
  }


  if account.PubKey.VerifySignature(signBytes, in.Signature) {
   vt = append(vt, in)
   res[string(in.Address)] = account
  }
 }
 return vt, nil
}

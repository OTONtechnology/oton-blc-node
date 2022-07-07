package tx

import (
 "encoding/hex"
 "encoding/json"

 "fmt"

 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
 "github.com/tendermint/tendermint/crypto/ed25519"
)


func (in Input) ValidateBasic() error {
 coins := Coins(in.Coins).ToStateCoins()
 if len(in.Address) != 20 {
  return fmt.Errorf("invalid address length")
 }
 if !coins.IsValid() {
  return fmt.Errorf("invalid coins: %v", in.Coins)
 }
 if coins.IsZero() {
  return fmt.Errorf("coins cannot be zero")
 }

 if !coins.IsPositive() {
  return fmt.Errorf("coins cannot be less than zero")
 }

 if in.Sequence <= 0 {
  return fmt.Errorf("sequence must be greater than 0")
 }
 if in.Sequence == 1 && in.PubKey == nil {
  return fmt.Errorf("PubKey must be present when Sequence == 1")
 }
 if in.Sequence > 1 && in.PubKey != nil {
  return fmt.Errorf("PubKey must be nil when Sequence > 1")
 }
 return nil
}

func (in Input) ValidateAdvanced(acc *state.Account, signBytes []byte) (res error) {


 seq, balance := acc.Sequence, acc.Balance
 if (seq + 1) != in.Sequence {

  return fmt.Errorf("got %v, expected %v. (acc.seq=%v), %w", in.Sequence, seq+1, acc.Sequence, fmt.Errorf("%v", ErrInvalidSequence))
 }

 if !balance.IsGTEs(Coins(in.Coins).ToStateCoins()) {
  return fmt.Errorf("balance is %v, tried to send %v, %w", balance, in.Coins, fmt.Errorf("%v", ErrInputsBalance))
 }

 if !acc.PubKey.VerifySignature(signBytes, in.Signature) {
  return fmt.Errorf("SignBytes: %X  %w", signBytes, fmt.Errorf("%v", ErrInvalidSignature))
 }
 return nil
}



type Inputs []Input

func (ins Inputs) ValidateBasic() error {
 if len(ins) == 0 {
  return fmt.Errorf("should contain at least one input")
 }
 for _, in := range ins {

  if err := in.ValidateBasic(); err != nil {
   return err
  }
 }
 return nil
}




func (ins Inputs) ValidateAdvanced(accounts map[string]*state.Account, signBytes []byte) (state.Coins, error) {
 total := state.Coins{}
 for _, in := range ins {
  acc := accounts[string(in.Address)]
  if acc == nil {
   return nil, fmt.Errorf("expects account in accounts")
  }
  if err := in.ValidateAdvanced(acc, signBytes); err != nil {
   return nil, err
  }

  total = total.Plus(Coins(in.Coins).ToStateCoins())
 }
 return total, nil
}

func (ins Inputs) Accounts(s *state.State) (map[string]*state.Account, error) {
 res := map[string]*state.Account{}
 for _, in := range ins {

  if _, ok := res[string(in.Address)]; ok {
   return nil, fmt.Errorf("ErrBaseDuplicateAddress: %x", in.Address)
  }

  account, err := s.GetAccount(in.Address)
  if err != nil {
   return nil, fmt.Errorf("get account: %v", err)
  }
  if account == nil {
   return nil, fmt.Errorf("ErrBaseUnknownAddress: %x", in.Address)
  }

  if in.PubKey != nil {
   account.PubKey = ed25519.PubKey(in.PubKey[:32])
  }
  res[string(in.Address)] = account
 }

 return res, nil
}




func (ins Inputs) AdjustAccounts(s *state.State, accounts map[string]*state.Account) error {
 for _, in := range ins {
  acc := accounts[string(in.Address)]
  if acc == nil {
   return fmt.Errorf("expects account in accounts")
  }
  coins := Coins(in.Coins).ToStateCoins()
  if !acc.Balance.IsGTEs(coins) {
   return fmt.Errorf("expects sufficient funds")
  }
  acc.Balance = acc.Balance.Minus(coins)
  acc.Sequence += 1
  if err := s.SetAccount(in.Address, acc); err != nil {
   return fmt.Errorf("set account: %v", err)
  }
 }
 return nil
}

func (ins Inputs) TotalCoins() state.Coins {
 total := state.Coins{}
 for _, out := range ins {
  total = total.Plus(Coins(out.Coins).ToStateCoins())
 }
 return total
}

type UnsignedCopier interface {
 UnsignedCopy() Inputs
}


func (ins Inputs) UnsignedCopy() Inputs {
 insCopy := make(Inputs, len(ins))
 copy(insCopy, ins)
 for i := range insCopy {
  insCopy[i].Signature = Signature{}
 }
 return insCopy
}



func (u *Input) MarshalJSON() ([]byte, error) {
 type Alias Input
 return json.Marshal(&struct {
  Address string `json:"address"`
  *Alias
 }{
  Address: hex.EncodeToString(u.Address),
  Alias: (*Alias)(u),
 })
}

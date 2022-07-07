package state

import (
 "encoding/gob"
 "fmt"

 "github.com/tendermint/tendermint/crypto"
 "github.com/tendermint/tendermint/crypto/ed25519"
)

func init() {

 gob.Register(ed25519.PubKey{})
}


type Account struct {
 Address crypto.Address `json:"adr"`
 PubKey crypto.PubKey `json:"pub_key"`
 Sequence int64 `json:"sequence"`
 Balance Coins `json:"coins"`

 Up crypto.Address `json:"up"`
 Down []crypto.Address `json:"dw"`
}

type GenesisDoc struct {
 PullFee crypto.Address `json:"pull_fee"`
 Ouner crypto.Address `json:"val_ouner"`
 Minter crypto.Address `json:"mint_ouner"`

 Accounts []Account `json:"accounts"`
}

func AccountKey(addr []byte) []byte {
 return append([]byte("base/a/"), addr...)
}

func (acc *Account) Copy() *Account {
 if acc == nil {
  return nil
 }
 accCopy := *acc
 return &accCopy
}

func (acc *Account) String() string {
 if acc == nil {
  return "nil-Account"
 }
 return fmt.Sprintf("Account{%v %v %v %v %v %v}",
  acc.Address, acc.PubKey, acc.Sequence, acc.Balance, acc.Up, acc.Down)
}



type PrivAccount struct {
 crypto.PrivKey
 Account
}

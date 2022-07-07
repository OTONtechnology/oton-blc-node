package tx

import (
 "bytes"
 "encoding/hex"
 "encoding/json"
 "fmt"

 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
)

func (out Output) ValidateBasic() error {
 coins := Coins(out.Coins).ToStateCoins()
 if _, _, err := out.ChainAndAddress(); err != nil {
  return err
 }
 if !coins.IsValid() {
  return fmt.Errorf("invalid coins %v", out.Coins)
 }
 if coins.IsZero() {
  return fmt.Errorf("coins cannot be zero")
 }


 if !coins.IsPositive() {
  return fmt.Errorf("coins cannot be less than zero")
 }

 return nil
}







func (out Output) ChainAndAddress() ([]byte, []byte, error) {
 var chainPrefix []byte
 address := out.Address
 if len(address) > 20 {
  spl := bytes.SplitN(address, []byte("/"), 2)
  if len(spl) != 2 {
   return nil, nil, fmt.Errorf("invalid address format")
  }
  chainPrefix = spl[0]
  address = spl[1]
 }


 if len(address) != 20 {
  return nil, nil, fmt.Errorf("invalid address length")
 }

 return chainPrefix, address, nil
}



type Outputs []Output

func (outs Outputs) ValidateBasic() error {
 for _, out := range outs {

  if err := out.ValidateBasic(); err != nil {
   return err
  }
 }
 return nil
}



func (outs Outputs) AddToAccountsMap(s *state.State, accounts map[string]*state.Account) (map[string]*state.Account, error) {
 var amc *state.AMC

 if accounts == nil {
  accounts = make(map[string]*state.Account)
 }

 for i, out := range outs {
  _, outAddress, _ := out.ChainAndAddress()






  if _, ok := accounts[string(outAddress)]; ok {
   return nil, fmt.Errorf("account duplicated")
  }

  acc, err := s.GetAccount(outAddress)
  if err != nil {
   return nil, fmt.Errorf("get account: %v", err)
  }


  if acc == nil {

   acc = &state.Account{}


   amc, err = s.GetAMC(outAddress)
   if err != nil {
    return nil, fmt.Errorf("get amc: %v", err)
   }


   if amc != nil {

    outAddress = append([]byte("amc/"), outAddress...)



    outs[i].Address = outAddress
    acc.Address = outAddress
    fmt.Printf("SendCoin to AMC(): %v \n", acc)
   }
  }
  accounts[string(outAddress)] = acc
 }
 return accounts, nil
}

func (outs Outputs) TotalCoins() state.Coins {
 total := state.Coins{}
 for _, out := range outs {
  total = total.Plus(Coins(out.Coins).ToStateCoins())
 }
 return total
}


func (outs Outputs) AdjustAccounts(s *state.State, accounts map[string]*state.Account) error {
 for _, out := range outs {
  destChain, outAddress, _ := out.ChainAndAddress()


  if destChain != nil {
   if bytes.Equal(destChain, []byte("amc")) {

    amc, err := s.GetAMC(outAddress)
    if err != nil || amc == nil {


     return fmt.Errorf("set amc: %v", err)
    }


    amc.Balance = amc.Balance.Plus(Coins(out.Coins).ToStateCoins())
    if err := s.SetAMC(amc); err != nil {
     return fmt.Errorf("set amc: %v", err)
    }
   }
   continue
  }

  acc := accounts[string(outAddress)]
  if acc == nil {
   return fmt.Errorf("expects account in accounts")
  }
  acc.Address = outAddress
  acc.Balance = acc.Balance.Plus(Coins(out.Coins).ToStateCoins())
  if err := s.SetAccount(outAddress, acc); err != nil {
   return fmt.Errorf("set account: %v", err)
  }
 }
 return nil
}


func (outs Outputs) MergeBalance() (nodub Outputs) {
 var nd Output

 accounts := make(map[string]*state.Account)

 for _, out := range outs {
  _, outAddress, _ := out.ChainAndAddress()

  if _, ok := accounts[string(outAddress)]; ok {
   accounts[string(outAddress)].Balance = accounts[string(outAddress)].Balance.Plus(Coins(out.Coins).ToStateCoins())
  } else {
   acc := &state.Account{}
   acc.Address = outAddress
   acc.Balance = Coins(out.Coins).ToStateCoins()

   accounts[string(outAddress)] = acc
  }
 }

 for _, acc := range accounts {
  nd = Output{acc.Address, FromStateCoins(acc.Balance)}
  nodub = append(nodub, nd)
 }
 return nodub
}



func (u *Output) MarshalJSON() ([]byte, error) {
 type Alias Output
 return json.Marshal(&struct {
  Address string `json:"address"`
  *Alias
 }{
  Address: hex.EncodeToString(u.Address),
  Alias: (*Alias)(u),
 })
}

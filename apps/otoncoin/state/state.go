package state

import (
 "bytes"
 "encoding/gob"
 "fmt"
 "time"

 "github.com/tendermint/tendermint/crypto"
 "github.com/tendermint/tendermint/libs/log"
)


const (

)





type State struct {
 chainID string
 store KVMem
 readCache map[string][]byte
 writeCache *KVCache
 logger log.Logger

 Size int64
 AppHash []byte
 Height int64


 ValProc CandidatProcessor


 Age AgeProcessor
}



var TimeWorkingBlock time.Time
var HeightLine int64

func NewState(store KVMem) *State {
 return &State{
  chainID: "",
  store: store,
  readCache: make(map[string][]byte),
  writeCache: nil,
  logger: log.NewNopLogger(),
 }
}

func (s *State) SetLogger(l log.Logger) {
 s.logger = l
}

func (s *State) SetChainID(chainID string) {
 s.chainID = chainID
 s.store.Set([]byte("base/chain_id"), []byte(chainID))
}

func (s *State) GetChainID() string {
 if s.chainID != "" {
  return s.chainID
 }
 s.chainID = string(s.store.Get([]byte("base/chain_id")))
 return s.chainID
}


func (s *State) SetAdrOuner(address crypto.Address) {
 s.ValProc.Ouner = address
 s.store.Set([]byte("base/ounerval"), []byte(address))
}

func (s *State) GetOuner() (*Account, error) {
 acc, err := s.GetAccount(s.ValProc.Ouner)
 return acc, err
}


func (s *State) SetAdrPullFee(address crypto.Address) {
 s.ValProc.PullFee = address
 s.store.Set([]byte("base/pullfee"), []byte(address))
}

func (s *State) GetPullFee() (*Account, error) {
 acc, err := s.GetAccount(s.ValProc.PullFee)
 return acc, err
}


func (s *State) SetAdrMinterOuner(address crypto.Address) {
 s.ValProc.Minter = address
 s.store.Set([]byte("base/minterouner"), []byte(address))
}
func (s *State) GetMinterOuner() (*Account, error) {
 acc, err := s.GetAccount(s.ValProc.Minter)
 return acc, err
}


func (s *State) Get(key []byte) (value []byte) {
 if s.readCache != nil {
  value, ok := s.readCache[string(key)]
  if ok {
   return value
  }
 }
 return s.store.Get(key)
}

func (s *State) Set(key []byte, value []byte) {
 if s.readCache != nil {
  s.readCache[string(key)] = value
 }
 s.store.Set(key, value)
}




func (s *State) GetAccount(address crypto.Address) (*Account, error) {
 var acc *Account
 err := s.GetValue(AccountKey(address), &acc)
 return acc, err
}


func (s *State) SetAccount(address crypto.Address, account *Account) error {
 return s.SetValue(AccountKey(address), account)
}




func (s *State) GetCoin(name string) (*Coin, error) {
 var coin *Coin
 err := s.GetValue(CoinKey(name), &coin)
 return coin, err
}


func (s *State) SetCoin(coin *Coin) error {
 return s.SetValue(CoinKey(coin.Name), coin)
}



func (s *State) GetAMC(address crypto.Address) (*AMC, error) {
 var amc *AMC
 err := s.GetValue(AMCKey(address), &amc)
 return amc, err
}

func (s *State) SetAMC(amc *AMC) error {
 return s.SetValue(AMCKey(amc.Address), amc)
}



type AccountDw struct {
 Acc Account
 Down []Account
}
type RefRank []*AccountDw



func (s *State) RefRank(refChain RefChain) (RefRank, error) {
 var address crypto.Address
 rf := RefRank{}


 cerr := 0


 for _, acc := range refChain {
  address = acc.Address

  if address != nil {
   rc := AccountDw{}
   rc.Acc = *acc
   for _, adr := range acc.Down {
    acc_dw, err := s.GetAccount(adr)





    if acc_dw != nil {
     rc.Down = append(rc.Down, *acc_dw)
    }

    if err != nil {
     cerr = cerr + 1
    }
   }


   rf = append(rf, &rc)
  }
 }

 if cerr != 0 {
  return rf, fmt.Errorf("get dw account: %v", cerr)
 } else {
  return rf, nil
 }
}



type RefChain []*Account

func (rt RefChain) Root() *Account {
 if len(rt) == 0 {
  return nil
 }
 return rt[len(rt)-1]
}

func (s *State) RefChain(address crypto.Address) (RefChain, error) {
 rc := RefChain{}
 for {
  acc, err := s.GetAccount(address)
  if err != nil {
   return nil, fmt.Errorf("get account: %v", err)
  }
  if acc == nil {
   return nil, fmt.Errorf("account not found")
  }


  acc.Address = address
  rc = append(rc, acc)

  if acc.Up == nil {
   return rc, nil
  }
  address = acc.Up
 }
}

func (s *State) SetRefChain(refChain RefChain) error {
 var address crypto.Address

 for _, acc := range refChain {

  if acc.PubKey == nil {
   address = acc.Address
  } else {
   address = acc.PubKey.Address()
   acc.Address = acc.PubKey.Address()
  }

  if address != nil {
   if err := s.SetAccount(address, acc); err != nil {
    return fmt.Errorf("set account: %v", err)
   }
  }
 }
 return nil
}




type RefTree []*Account



func (s *State) RefTree(address crypto.Address) (RefTree, error) {
 rt := RefTree{}
 added := map[string]bool{}
 queue := []crypto.Address{address}
 for len(queue) > 0 {
  next := queue[0]
  queue = queue[1:]

  if _, ok := added[next.String()]; ok {
   return nil, fmt.Errorf("account '%s' already added to RefTree", next.String())
  }

  acc, err := s.GetAccount(next)
  if err != nil {
   return nil, fmt.Errorf("get account: %v", err)
  }
  if acc == nil {
   return nil, fmt.Errorf("account not found")
  }
  rt = append(rt, acc)
  added[next.String()] = true
  queue = append(queue, acc.Down...)
 }
 return rt, nil
}

func (s *State) SetRefTree(refTree RefTree) error {
 for _, acc := range refTree {
  if err := s.SetAccount(acc.PubKey.Address(), acc); err != nil {
   return fmt.Errorf("set account: %v", err)
  }
 }
 return nil
}



func (s *State) CacheWrap() *State {
 cache := NewKVCache(s)
 return &State{
  chainID: s.chainID,
  store: cache,
  readCache: nil,
  writeCache: cache,
  logger: s.logger,
 }
}


func (s *State) CacheSync() {
 s.writeCache.Sync()
}

func (s *State) Commit() error {







 return nil
}

func (s *State) GetValue(key []byte, v interface{}) error {
 data := s.store.Get(key)
 if len(data) == 0 {
  return nil
 }
 err := Deserialize(data, v)
 if err != nil {
  return fmt.Errorf("reading account %X error: %v", data, err)
 }
 return nil
}

func (s *State) SetValue(key []byte, v interface{}) error {
 data, err := Serialize(v)
 if err != nil {
  return fmt.Errorf("serialize value: %v", err)
 }
 s.store.Set(key, data)
 return nil
}


func Serialize(v interface{}) ([]byte, error) {
 var buff bytes.Buffer
 enc := gob.NewEncoder(&buff)
 err := enc.Encode(v)
 return buff.Bytes(), err
}


func Deserialize(data []byte, v interface{}) error {
 dec := gob.NewDecoder(bytes.NewReader(data))
 err := dec.Decode(v)
 return err
}

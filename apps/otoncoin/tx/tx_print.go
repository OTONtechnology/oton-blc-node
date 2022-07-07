package tx

import (
 "encoding/hex"
 "encoding/json"
 "time"
)

type CoinPr struct {
 Name string `json:"name,omitempty"`
 Amount int64 `json:"amount,omitempty"`
 Expired time.Time `json:"expired"`
}

type OutputsPr []OutputPr
type OutputPr struct {
 Address []byte `json:"address,omitempty"`
 Coins []CoinPr `json:"coins"`
}

func (u *OutputPr) MarshalJSON() ([]byte, error) {
 type Alias OutputPr
 return json.Marshal(&struct {
  Address string `json:"address"`
  *Alias
 }{
  Address: hex.EncodeToString(u.Address),
  Alias: (*Alias)(u),
 })
}


func (u *CoinPr) MarshalJSON() ([]byte, error) {
 type Alias CoinPr

 if u.Expired.Unix() > int64(-62135596800) {
  return json.Marshal(&struct {
   Expired time.Time `json:"expired"`
   *Alias
  }{
   Expired: u.Expired,
   Alias: (*Alias)(u),
  })
 } else {
  return json.Marshal(&struct {
   Expired *byte `json:"expired"`
   *Alias
  }{
   Expired: nil,
   Alias: (*Alias)(u),
  })
 }
}

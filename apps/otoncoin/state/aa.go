package state


import (
 "github.com/tendermint/tendermint/crypto"
)

type AppMark interface {
 MakePayment() int64
}

type Mark struct {
 Sequence int64 `json:"sequence"`
 Balance Coins `json:"coins"`
}

type MakrOne struct {
 Mark
 up int64
}


func (pm *MakrOne) MakePayment() int64 {
 return 11
}

type MakrTwo struct {
 Mark
 up3 int64
}


func (pm *MakrTwo) MakePayment() int64 {
 return 15
}

type AMC21 struct {
 Pays []AppMark `json:"pay_maker"`
}

func NewAMC21() (*AMC21, error) {
 var line *MakrOne
 var career *MakrTwo

 career = &MakrTwo{}
 line = &MakrOne{}

 line.up = 63
 career.up3 = line.MakePayment()

 return &AMC21{
  Pays: []AppMark{line, career},
 }, nil
}


type AMC2 struct {
 Address crypto.Address
 PaymentMakers []PaymentMaker `json:"pay_maker"`
}

func NewAMC2(adr crypto.Address) (*AMC2, error) {
 var line *LineMarketing
 var career *CareerMarketing

 career = &CareerMarketing{}
 line = &LineMarketing{}

 line.Levels = []int64{98, 93, 77, 82, 83}

 return &AMC2{
  Address: adr,
  PaymentMakers: []PaymentMaker{line, career, line},
 }, nil
}


func (s *State) GetAMC2(address crypto.Address) (*AMC2, error) {
 var amc *AMC2
 err := s.GetValue(AMCKey(address), &amc)
 return amc, err
}

func (s *State) SetAMC2(amc *AMC2) error {
 return s.SetValue(AMCKey(amc.Address), amc)
}

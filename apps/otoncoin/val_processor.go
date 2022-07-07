package otoncoin


import (
 "bytes"
 "encoding/base64"
 "fmt"
 "strings"

 "github.com/tendermint/tendermint/abci/apps/otoncoin/tx"
 "github.com/tendermint/tendermint/abci/types"

 cryptoenc "github.com/tendermint/tendermint/crypto/encoding"
 pc "github.com/tendermint/tendermint/proto/tendermint/crypto"
)

const (

 ValidatorSetChangePrefix string = "val:"


 runCandidat = 0
 stopCandidat = 5

 relaxCandidat = 6
 waitCandidat = 7

)


func isValidatorTx(tx []byte) bool {
 return strings.HasPrefix(string(tx), ValidatorSetChangePrefix)
}


func MakeValSetChangeTx(pubkey pc.PublicKey, power int64) []byte {
 pk, err := cryptoenc.PubKeyFromProto(pubkey)
 if err != nil {
  panic(err)
 }
 pubStr := base64.StdEncoding.EncodeToString(pk.Bytes())
 return []byte(fmt.Sprintf("val:%s!%d", pubStr, power))
}


func (app *Application) Validators() (validators []types.ValidatorUpdate) {

 fmt.Printf("app.Validators()... \n")

 for adr, cv := range app.state.ValProc.ValMap {

  fmt.Printf("adr: %x => pubkey: %x \n", adr, cv.PubKey.Bytes())



  if cv.Step != stopCandidat {
   key := []byte("val:" + string(cv.PubKey.Bytes()))
   vl := app.state.Get(key)

   if len(vl) > 0 {
    validator := new(types.ValidatorUpdate)
    err := types.ReadMessage(bytes.NewBuffer(vl), validator)
    if err != nil {
     panic(err)
    }
    validators = append(validators, *validator)
   }
  }
 }

 fmt.Printf(". \n")
 return
}







func (app *Application) updateValidator(v types.ValidatorUpdate) types.ResponseDeliverTx {
 pubkey, err := cryptoenc.PubKeyFromProto(v.PubKey)
 if err != nil {
  panic(fmt.Errorf("can't decode public key: %w", err))
 }

 fmt.Printf("update val: %x \n", string(pubkey.Bytes()))
 key := []byte("val:" + string(pubkey.Bytes()))

 dubStop := runCandidat

 if v.Power == 0 {
  hasKey := app.state.Get(key)

  if len(hasKey) == 0 {
   pubStr := base64.StdEncoding.EncodeToString(pubkey.Bytes())
   return types.ResponseDeliverTx{
    Code: tx.TypeUnauthorized,
    Log: fmt.Sprintf("Cannot remove non-existent validator %s", pubStr)}
  }





  candidat := app.state.ValProc.ValMap[string(pubkey.Address())]


  if candidat.Step == stopCandidat {
   dubStop = stopCandidat
  }







  candidat.Step = stopCandidat
  app.state.ValProc.ValMap[string(pubkey.Address())] = candidat

  if dubStop != stopCandidat {
   fmt.Printf("del CV: %v \n", candidat)
   app.ValUpdates = append(app.ValUpdates, v)
  } else {
   fmt.Printf("dub CV: %v \n", candidat)
  }

 } else {

  value := bytes.NewBuffer(make([]byte, 0))
  if err := types.WriteMessage(&v, value); err != nil {
   return types.ResponseDeliverTx{
    Code: tx.TypeEncodingError,
    Log: fmt.Sprintf("Error encoding validator: %v", err)}
  }

  app.state.Set(key, value.Bytes())


  candidat := app.state.ValProc.ValMap[string(pubkey.Address())]

  candidat.PubKey = pubkey
  candidat.Power = v.Power







  candidat.Step = runCandidat


  app.state.ValProc.ValMap[string(pubkey.Address())] = candidat
  fmt.Printf("mod CV: %v \n", candidat)


  app.ValUpdates = append(app.ValUpdates, v)
 }

 return types.ResponseDeliverTx{Code: tx.TypeOK}
}

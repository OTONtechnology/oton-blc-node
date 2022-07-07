package tx

import (
 "fmt"
 "reflect"

 "github.com/gogo/protobuf/proto"
 "github.com/tendermint/tendermint/abci/apps/otoncoin/state"
 "github.com/tendermint/tendermint/abci/types"
)

type Tx interface {
 proto.Marshaler
 proto.Unmarshaler
 Exec(state *state.State, hash []byte, check bool) types.ResponseDeliverTx
}

type Signature = []byte



func SignBytes(tx proto.Marshaler, chainID string) []byte {
 signBytes := []byte(chainID)
 txCopy := Copy(tx).(proto.Marshaler)
 data, err := txCopy.Marshal()
 if err != nil {
  panic(fmt.Errorf("marshal transation: %v", err))
 }
 return append(signBytes, data...)
}


func Copy(src interface{}) interface{} {
 newObj := reflect.New(reflect.TypeOf(src).Elem())
 oldVal := reflect.ValueOf(src).Elem()
 newVal := newObj.Elem()
 for i := 0; i < oldVal.NumField(); i++ {
  newValField := newVal.Field(i)
  if newValField.CanSet() {
   field := oldVal.Field(i)
   if ins, ok := field.Interface().([]Input); ok {
    field = reflect.ValueOf(Inputs(ins).UnsignedCopy())
   }
   newValField.Set(field)
  }
 }
 return newObj.Interface()
}

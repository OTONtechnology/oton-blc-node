package state

import (
 "bytes"
 "fmt"
 "testing"

 "github.com/stretchr/testify/assert"
 "github.com/tendermint/tendermint/libs/log"

)

func TestState(t *testing.T) {
 assert := assert.New(t)


 store := NewMemKVStore()
 state := NewState(store)
 state.SetLogger(log.TestingLogger())
 cache := state.CacheWrap()


 fmt.Printf("state.ValProc.ValMap = %v \n", len(state.ValProc.ValMap["dfg"].Contract))


 dumAddr := []byte("dummyAddress")

 acc := new(Account)
 acc.Sequence = 1


 reset := func() {
  store = NewMemKVStore()
  state = NewState(store)
  state.SetLogger(log.TestingLogger())
  cache = state.CacheWrap()
 }
# 48 "state_test.go"
 keyvalue := []struct {
  key string
  value string
 }{
  {"foo", "snake"},
  {"bar", "mouse"},
 }


 setRecords := func(kv KVMem) {
  for _, n := range keyvalue {
   kv.Set([]byte(n.key), []byte(n.value))
  }
 }


 storeHasAll := func(kv KVMem) bool {
  for _, n := range keyvalue {
   if !bytes.Equal(kv.Get([]byte(n.key)), []byte(n.value)) {
    return false
   }
  }
  return true
 }


 state.SetChainID("testchain")
 assert.Equal(state.GetChainID(), "testchain", "ChainID is improperly stored")


 setRecords(state)
 assert.True(storeHasAll(state), "state doesn't retrieve after Set")


 assert.NoError(state.SetAccount(dumAddr, acc))
 acc, err := state.GetAccount(dumAddr)
 assert.NoError(err)
 fmt.Printf("dumAddr.Sequence = %d \n", acc.Sequence)
# 95 "state_test.go"
 reset()
 setRecords(cache)
 assert.False(storeHasAll(store), "store retrieving before CacheSync")
 cache.CacheSync()
 assert.True(storeHasAll(store), "store doesn't retrieve after CacheSync")
# 112 "state_test.go"
}

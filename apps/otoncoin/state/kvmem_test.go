package state

import (
 "bytes"
 "testing"

 "github.com/stretchr/testify/assert"
)

func TestKVStore(t *testing.T) {
 assert := assert.New(t)


 ms := NewMemKVStore()
 store := NewMemKVStore()
 kvc := NewKVCache(store)


 var keyvalue = []struct {
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


 setRecords(ms)
 assert.True(storeHasAll(ms), "MemKVStore doesn't retrieve after Set")


 setRecords(kvc)
 assert.True(storeHasAll(kvc), "KVCache doesn't retrieve after Set")


 kvc.Reset()
 assert.False(storeHasAll(kvc), "KVCache retrieving after reset")


 setRecords(kvc)
 assert.False(storeHasAll(store), "store retrieving before synced")
 kvc.Sync()
 assert.True(storeHasAll(store), "store isn't retrieving after synced")


 assert.Zero(len(kvc.GetLogLines()), "logging events existed before using SetLogging")
 kvc.SetLogging()
 setRecords(kvc)
 assert.Equal(len(kvc.GetLogLines()), 2, "incorrect number of logging events recorded")
 kvc.ClearLogLines()
 assert.Zero(len(kvc.GetLogLines()), "logging events still exists after ClearLogLines")

}

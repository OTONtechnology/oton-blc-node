package state

import (
 "encoding/gob"
 "fmt"
 "testing"

 "github.com/stretchr/testify/assert"
)

func TestNilAccount(t *testing.T) {

 rt, err := NewAMC21()
 assert.NoError(t, err)


 gob.Register(&MakrOne{})
 gob.Register(&MakrTwo{})

 data, err := Serialize(rt)
 assert.NoError(t, err)

 ty := AMC21{}

 err = Deserialize(data, &ty)
 assert.NoError(t, err)


 store := NewMemKVStore()
 state := NewState(store)


 amc2, err := NewAMC2([]byte("dumAddr"))
 err = state.SetAMC2(amc2)
 fmt.Printf("SetAMC2(dumAddr)  %v, %v \n", amc2, err)

 amc2, err = state.GetAMC2([]byte("dumAddr"))
 fmt.Printf("GetAMC2(dumAddr)  %v, %v \n", amc2, err)
}

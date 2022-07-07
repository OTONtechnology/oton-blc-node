package otoncoin

import (
	mrand "math/rand"

	"github.com/tendermint/tendermint/abci/types"
	tmrand "github.com/tendermint/tendermint/libs/rand"
)

func RandVal(i int) types.ValidatorUpdate {
	pubkey := tmrand.Bytes(32)

	power := mrand.Uint32() & (1<<16 - 1)
	v := types.UpdateValidator(pubkey, int64(power), "")
	return v
}

func RandVals(cnt int) []types.ValidatorUpdate {
	res := make([]types.ValidatorUpdate, cnt)
	for i := 0; i < cnt; i++ {
		res[i] = RandVal(i)
	}
	return res
}

func InitKVStore(app *Application) {
	app.InitChain(types.RequestInitChain{
		Validators: RandVals(1),
	})
}

var testGenesisFmtVal = `{
  "accounts":[
      {"pub_key":{
 "type":"tendermint/PubKeyEd25519",
 "value":"AxZUG 4="},
 "sequence":"0",
    "coins":[{
  "Name":"feecoin",
  "Amount":"100",
  "DecimalPoint":0,
  "Movable":true,
  "Hash":null,
  "Delta":"0",
  "CreatorAddr":"9BA140CC905073E01DA1A1B47FD629DD8F24093F",
  "Expired":"0001-01-01T00:00:00Z",
  "Chart":null
 }],
    "up":"",
 "dw":null}
   ]
}`

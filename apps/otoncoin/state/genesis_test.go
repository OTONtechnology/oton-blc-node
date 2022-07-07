package state


import (
 "encoding/base64"
 "fmt"
 "testing"

 "github.com/tendermint/tendermint/crypto/ed25519"
 "github.com/tendermint/tendermint/libs/json"
)

func TestGenesisAccount(t *testing.T) {


 privKey1 := ed25519.GenPrivKey()
 pubKey1 := privKey1.PubKey()


 privKey2 := ed25519.GenPrivKey()
 pubKey2 := privKey2.PubKey()

 acc1 := &Account{
  Address: pubKey1.Address(),
  Balance: Coins{
   Coin{Name: "mycoin", Amount: 10, Movable: true},
   Coin{Name: "testcoin", Amount: 20, Movable: true},
  },
  PubKey: pubKey1,
 }


 bz, _ := json.MarshalIndent(acc1, "", "\t")
 fmt.Printf("\"accounts\":[%s]\n \n", bz)

 fmt.Printf("pubKey1.adr %x \n", pubKey1.Address())
 fmt.Printf("pubKey1.b64 %s \n", base64.StdEncoding.EncodeToString(pubKey1.Bytes()))
 fmt.Printf("pubKey1 %x \n", pubKey1.Bytes())
 fmt.Printf("prvKey1 %x \n\n", privKey1.Bytes())

 fmt.Printf("pubKey2.adr %x \n", pubKey2.Address())
 fmt.Printf("pubKey2 %x \n", pubKey2.Bytes())
 fmt.Printf("prvKey2 %x \n\n", privKey2.Bytes())


 var pv0 ed25519.PrivKey
 pv0 = privKey1.Bytes()
 fmt.Printf("from prv1 %x \n\n", pv0)

}

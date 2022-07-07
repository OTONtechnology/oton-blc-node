package tx

import (
 "encoding/hex"
 "fmt"
 "testing"

 "github.com/tendermint/tendermint/crypto"
 "github.com/tendermint/tendermint/libs/json"

 "github.com/gogo/protobuf/proto"
)

const (
 testHex0 = "0A0A73656E645F636F696E7312DF01120A0A066D79636F696E10021A9A010A149BA140CC905073E01DA1A1B47FD629DD8F24093F120C0A066D79636F696E10C2843D120E0A0874657374636F696E10C0843D18012240261F84D7A9E0F2C934E5F82F0E59F69F944F02842779B2D13F11DFA1B81EE549BC17CAB25612BC52CE3BF37B9B98A06774C99421E6A475D3762C9E4930B00D0F2A200316541BE8E70816CDFF7EE3D3A1B0A597CF66C052AA1F766418E1D91074318E22340A14068F3AAE6C41CF4A177EE19201F8B7195E0DDC30120C0A066D79636F696E10C0843D120E0A0874657374636F696E10C0843D"
 testHex1 = "0a0a6372656174655f616d6312a1030a0f0a0b626974626f6e65436f696e1001128f010a14944940a5d3ab262f939b8be0c7ebf2a52a7ff03812110a0b626974626f6e65436f696e10c0843d180122408d7fec14c10492e932a52bd3151ad8e476326fe785a3fea05e07037e48c3a2ee9b4dd7e1cef912d130fc5d50ce28e5c84ec1b80f583507e3178fbb83c6df18052a20422814c86601b619969de8f8f0f9908c4a7020d7eaedf83746116cfc1984c50a1a1437e637288b59014712e7221dc5a40dd0f32ce097220e626974626f6e65436f6d70616e792a0e0a0c6c696e652e6c6576656c2e302a100a0c6c696e652e6c6576656c2e31100a2a100a0c6c696e652e6c6576656c2e32100a2a100a0c6c696e652e6c6576656c2e33100a2a100a0c6c696e652e6c6576656c2e34100a2a100a0c6c696e652e6c6576656c2e35100a2a100a0c6c696e652e6c6576656c2e36100a2a100a0c6c696e652e6c6576656c2e37100a2a100a0c6c696e652e6c6576656c2e38100a2a100a0c6c696e652e6c6576656c2e39100a2a110a0d6c696e652e7061636b65742e3110082a110a0d6c696e652e7061636b65742e321007"
 testHex2 = "0a0a7365745f696e5f616d6312ca010a140a107461726966426974626f6e65436f696e100112700a141eea9d501e12dac3c2b2d6be514c00216cd531a412140a107461726966426974626f6e65436f696e1001181f2240d81b79149ffbd865ff3dbfb75a565051354b61ba900c4c5784615c87c9f44e55a9b12f0b0096c5dd9f7be40c5be449700b1e3e0f75a03350a62fa148bdc6280d1a14fbbf87478d8107c21a4b41215fc6c5387593750a221457dea6735151bd4b696a8929aed83ecef0de8f5f2a1414092049ddb7e0064085ca7a4a3fcdf27a9b50d7"


 testHex = "0a086275795f6e6f646512d3010a080a046f746f6e100112650a14b6e37b549221fd09a9a87994cd712a881c092be112080a046f746f6e100518c8152240a92a0897720ce38d3bb54bac31739b41d44512d427eaabf6b2b6d15bce29250ffd68100d44c1aa09b48f8605e9df655ea9fbe57c58626d9fd2d15190f7c2cd0e1a200a144055c8c4d9ceefa8cae2f53748d19e46f9a69d5412080a046f746f6e10042a20ca6878518480f37e9ce75363a18f63055f445ae9665ab82e97d7334328f1721630023a100a0c436f6e74726163744e6f6465100442080a046f746f6e1001"
)

func txPrint(tx proto.Unmarshaler, raw []byte) {
 t := Copy(tx).(proto.Unmarshaler)

 if err := t.Unmarshal(raw); err != nil {
  fmt.Printf("Error Unmarshal(Tx): %v", err)
 }

 sTx, _ := json.MarshalIndent(t, "", "\t")
 fmt.Printf("Tx %s \n", sTx)
}


func TestParce_Exec(t *testing.T) {

 t.Run("Unmarshal Tx", func(t *testing.T) {


  var rawTx Raw
  var Tx, _ = hex.DecodeString(testHex)

  hesh := crypto.Sha256(Tx)
  fmt.Printf("hash = sha256(tx): %x \n", hesh)

  err := rawTx.Unmarshal(Tx)
  if err != nil {
   fmt.Printf("Error Unmarshal(Tx): %v \n", err)
  }

  fmt.Printf("Type Tx %v \n", rawTx.Type)

  switch rawTx.Type {
  case TypeSendCoins:
   txPrint(&SendCoins{}, rawTx.Raw)
  case TypeMintCoins:
   txPrint(&MintCoins{}, rawTx.Raw)
  case TypeTest:
   txPrint(&SendCoins{}, rawTx.Raw)


  case TypeCreateAMC:
   txPrint(&CreateAMC{}, rawTx.Raw)
  case TypeSetInAMC:
   txPrint(&SetInAMC{}, rawTx.Raw)
  case TypeBuyInAMC:
   txPrint(&BuyInAMC{}, rawTx.Raw)


  case TypeChangeWant:
   txPrint(&ChangeAddressWant{}, rawTx.Raw)
  case TypeChangeApply:
   txPrint(&ChangeAddressApply{}, rawTx.Raw)


  case TypeInteresNode:
   txPrint(&SetInteresNode{}, rawTx.Raw)

  case TypeNodeReward:
   txPrint(&SetNodeReward{}, rawTx.Raw)

  case TypeBuyNode:
   txPrint(&BuyNode{}, rawTx.Raw)

  default:
   fmt.Printf("Unknown Type Tx %v \n", rawTx.Type)

  }
 })
}

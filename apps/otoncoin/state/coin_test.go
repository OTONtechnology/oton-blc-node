package state

import (
 "fmt"
 "testing"
 "time"

 "github.com/stretchr/testify/assert"
)

func TestCoins(t *testing.T) {
 assert := assert.New(t)


 good := Coins{
  NewCoin("GAS", 1),
  NewCoin("MINERAL", 1),
  NewCoin("TREE", 1),
 }
 neg := good.Negative()
 sum := good.Plus(neg)
 empty := Coins{
  NewCoin("GOLD", 0),
 }
 badSort1 := Coins{
  NewCoin("TREE", 1),
  NewCoin("GAS", 1),
  NewCoin("MINERAL", 1),
 }
 badSort2 := Coins{
  NewCoin("TREE", 1),
  NewCoin("GAS", 1),
  NewCoin("MINERAL", 1),
 }
 badAmt := Coins{
  NewCoin("GAS", 1),
  NewCoin("TREE", 0),
  NewCoin("MINERAL", 1),
 }
 dup := Coins{
  NewCoin("GAS", 1),
  NewCoin("GAS", 1),
  NewCoin("MINERAL", 1),
 }

 mc := Coins{NewCoin("mycoin", 2)}
 gs1 := Coins{NewCoin("GAS", 1)}
 gs2 := Coins{NewCoin("GAS", 2)}


 assert.False(badSort1.IsGTE(gs1), "%v <= %v", badSort1, gs1)
 assert.True(badSort1.IsGTEs(gs1), "%v <= %v", badSort1, gs1)
 assert.True(good.IsGTE(gs1), "%v <= %v", good, gs1)

 assert.False(good.IsGTE(mc), "Expected %v to be >= %v", good, mc)
 assert.False(gs1.IsGTE(gs2), "%v < %v", gs1, gs2)
 assert.True(gs2.IsGTE(gs1), "%v > %v", gs2, gs1)
 assert.True(gs1.IsGTE(gs1), "%v = %v", gs1, gs1)

 assert.True(good.IsValid(), "Coins are valid")
 assert.True(good.IsPositive(), "Expected coins to be positive: %v", good)
 assert.True(good.IsGTE(empty), "Expected %v to be >= %v", good, empty)
 assert.False(neg.IsPositive(), "Expected neg coins to not be positive: %v", neg)
 assert.Zero(len(sum), "Expected 0 coins")
 assert.True(badSort1.IsValid(), "Coins are not sorted")
 assert.False(badSort2.IsValid(), "Coins are not sorted")
 assert.False(badAmt.IsValid(), "Coins cannot include 0 amounts")
 assert.False(dup.IsValid(), "Duplicate coin")


 tr0 := &Coin{Name: AMCCoinTarifBbon, Amount: 3}
 trA := &Coin{Name: AMCCoinTarifBbon, Amount: 4, Expired: time.Now().Add(2 * time.Hour)}
 trN := &Coin{Name: AMCCoinTarifBbon, Amount: 0, Expired: time.Now().Add(24 * time.Hour)}
 trE := &Coin{Name: "Empty", Amount: 2, Expired: time.Now().Add(3 * time.Hour)}


 Balance := Coins{*tr0, *trA, *trE}
 Balance = append(Balance, *trN)
 fmt.Printf("\nAll coins on balance separately... \n")
 fmt.Printf("  %v \n  %v \n  %v \n  %v \n \n", Balance[0], Balance[1], Balance[2], Balance[3])

 tarif := Balance.CoinNoExpired(AMCCoinTarifBbon)
 fmt.Printf("Sum coin no expired %v  \n", tarif)

 tarif = Balance.Coin(AMCCoinTarifBbon)
 fmt.Printf("Balance.Coin(Tarif)  %v  \n", tarif)
 tarif = Balance.Coin("Empty")
 fmt.Printf("Balance.Coin(Empty)  %v  \n \n\n", tarif)


 fmt.Printf("Test Balance.Plus(Coin) \n")
 Balance = Coins{*tr0}
 fmt.Printf("tr0 +3 %v \n", Balance)
 Balance = Balance.Plus(Coins{*trA})
 fmt.Printf("trA +4 %v \n", Balance)
 Balance = Balance.Plus(Coins{*trN})
 fmt.Printf("trN +0 %v \n", Balance)
 Balance = Balance.Plus(Coins{*trE})
 fmt.Printf("trE +2e %v \n \n", Balance)

 tarif = Balance.CoinNoExpired(AMCCoinTarifBbon)
 fmt.Printf("Sum coin no expired %v  \n", tarif)

 tarif = Balance.Coin(AMCCoinTarifBbon)
 fmt.Printf("Balance.Coin(Tarif)  %v  \n", tarif)
 tarif = Balance.Coin("Empty")
 fmt.Printf("Balance.Coin(Empty)  %v  \n \n\n", tarif)

}


func TestParse(t *testing.T) {
 assert := assert.New(t)

 cases := []struct {
  input string
  valid bool
  expected Coins
 }{
  {"", true, nil},
  {"1foo", true, Coins{NewCoin("foo", 1)}},
  {"10bar", true, Coins{NewCoin("bar", 10)}},
  {"99bar,1foo", true, Coins{NewCoin("bar", 99), NewCoin("foo", 1)}},
  {"98 bar , 1 foo  ", true, Coins{NewCoin("bar", 98), NewCoin("foo", 1)}},
  {"  55\t \t bling\n", true, Coins{NewCoin("bling", 55)}},
  {"2foo, 97 bar", true, Coins{NewCoin("bar", 97), NewCoin("foo", 2)}},
  {"5 mycoin,", false, nil},
  {"2 3foo, 97 bar", false, nil},
  {"11me coin, 12you coin", false, nil},
  {"1.2btc", false, nil},
  {"5foo-bar", false, nil},
 }

 for _, tc := range cases {
  res, err := ParseCoins(tc.input)
  if !tc.valid {
   assert.NotNil(err, "%s: %#v", tc.input, res)
  } else if assert.Nil(err, "%s: %+v", tc.input, err) {
   assert.Equal(tc.expected, res)
  }
 }

}

func TestSortCoins(t *testing.T) {
 assert := assert.New(t)

 good := Coins{
  NewCoin("GAS", 1),
  NewCoin("MINERAL", 1),
  NewCoin("TREE", 1),
 }
 empty := Coins{
  NewCoin("GOLD", 0),
 }
 badSort1 := Coins{
  NewCoin("TREE", 1),
  NewCoin("GAS", 1),
  NewCoin("MINERAL", 1),
 }
 badSort2 := Coins{
  NewCoin("GAS", 1),
  NewCoin("TREE", 1),
  NewCoin("MINERAL", 1),
 }
 badAmt := Coins{
  NewCoin("GAS", 1),
  NewCoin("TREE", 0),
  NewCoin("MINERAL", 1),
 }
 dup := Coins{
  NewCoin("GAS", 1),
  NewCoin("GAS", 1),
  NewCoin("MINERAL", 1),
 }

 cases := []struct {
  coins Coins
  before, after bool
 }{
  {good, true, true},
  {empty, false, false},
  {badSort1, false, true},
  {badSort2, false, true},
  {badAmt, false, false},
  {dup, false, false},
 }

 for _, tc := range cases {
  assert.Equal(tc.before, tc.coins.IsValid())
  tc.coins.Sort()
  assert.Equal(tc.after, tc.coins.IsValid())
 }
}

package state

import (
 "fmt"
 "regexp"
 "sort"
 "strconv"
 "strings"
 "time"

 "github.com/pkg/errors"
 "github.com/tendermint/tendermint/crypto"
)

type Coin struct {
 Name string
 Amount int64
 DecimalPoint uint32
 Movable bool
 Hash []byte
 Delta int64
 CreatorAddr crypto.Address
 Expired time.Time




 Chart Coins `json:"Chart"`
}

func NewCoin(name string, amount int64) Coin {
 return Coin{Name: name, Amount: amount}
}

func CoinKey(name string) []byte {
 return append([]byte("base/c/"), name...)
}

func (coin Coin) String() string {

 if coin.Expired.Unix() > 0 {
  return fmt.Sprintf("{%v %v %v %v}", coin.Amount, coin.Name, coin.CreatorAddr, coin.Expired)
 } else {
  return fmt.Sprintf("{%v %v %v }", coin.Amount, coin.Name, coin.CreatorAddr)
 }
}

func (coin Coin) Percent(percent int64) Coin {
 return Coin{
  Name: coin.Name,
  Amount: coin.Amount * percent / 100,
 }
}

func (coin Coin) Fraction(fraction float64) Coin {
 return Coin{
  Name: coin.Name,
  Amount: int64(float64(coin.Amount) * fraction),
 }
}


func (coin Coin) Addon(point int64) Coin {
 return Coin{
  Name: coin.Name,
  Amount: coin.Amount + point,
 }
}


var reName = regexp.MustCompile("")
var reAmt = regexp.MustCompile("(\\d+)")
var reCoin = regexp.MustCompile("^([[:digit:]]+)[[:space:]]*([[:alpha:]]+)$")


func ParseCoin(str string) (Coin, error) {
 var coin Coin

 matches := reCoin.FindStringSubmatch(strings.TrimSpace(str))
 if matches == nil {
  return coin, errors.Errorf("%s is invalid coin definition", str)
 }


 amt, err := strconv.Atoi(matches[1])
 if err != nil {
  return coin, err
 }



 coin.Name = matches[2]
 coin.Amount = int64(amt)

 return coin, nil
}

type Coins []Coin

func (coins Coins) String() string {
 if len(coins) == 0 {
  return ""
 }

 out := ""
 for _, coin := range coins {
  out += fmt.Sprintf("%v,", coin.String())
 }
 return out[:len(out)-1]
}


func (coins Coins) Coin(name string) *Coin {
 for _, coin := range coins {
  if coin.Name == name {
   return &coin
  }
 }
 return nil
}


func (coins Coins) CoinNoExpired(name string) *Coin {
 sum := Coin{Name: name}
 for _, coin := range coins {
  if coin.Name == name && coin.Expired.Unix() > time.Now().Unix() {
   sum.Amount += coin.Amount
   if sum.Expired.Unix() < coin.Expired.Unix() {
    sum.Expired = coin.Expired
   }
  }
 }
 return &sum
}



func (coins Coins) IsMovable(s *State) bool {
 for _, coin := range coins {
  coinfull, err := s.GetCoin(coin.Name)


  if (err != nil) || (coinfull == nil) {
   return false
  }

  if !coinfull.Movable {
   return false
  }
 }
 return true
}

func ParseCoins(str string) (Coins, error) {

 if len(str) == 0 {
  return nil, nil
 }

 split := strings.Split(str, ",")
 var coins Coins

 for _, el := range split {
  coin, err := ParseCoin(el)
  if err != nil {
   return coins, err
  }
  coins = append(coins, coin)
 }


 coins.Sort()
 if !coins.IsValid() {
  return nil, errors.Errorf("ParseCoins invalid: %#v", coins)
 }

 return coins, nil
}


func (coins Coins) IsValid() bool {
 switch len(coins) {
 case 0:
  return true
 case 1:
  return coins[0].Amount != 0
 default:
  lowDenom := coins[0].Name
  for _, coin := range coins[1:] {
   if coin.Name <= lowDenom {
    return false
   }
   if coin.Amount == 0 {
    return false
   }

   lowDenom = coin.Name
  }
  return true
 }
}



func (coinsA Coins) Plus(coinsB Coins) Coins {
 sum := []Coin{}
 indexA, indexB := 0, 0
 lenA, lenB := len(coinsA), len(coinsB)
 for {
  if indexA == lenA {
   if indexB == lenB {
    return sum
   } else {
    return append(sum, coinsB[indexB:]...)
   }
  } else if indexB == lenB {
   return append(sum, coinsA[indexA:]...)
  }
  coinA, coinB := coinsA[indexA], coinsB[indexB]
  switch strings.Compare(coinA.Name, coinB.Name) {
  case -1:
   sum = append(sum, coinA)
   indexA += 1
  case 0:
   if coinA.Amount+coinB.Amount == 0 {

   } else {
    sum = append(sum, Coin{
     Name: coinA.Name,
     Amount: coinA.Amount + coinB.Amount,
     Expired: coinB.Expired,
    })
   }
   indexA += 1
   indexB += 1
  case 1:
   sum = append(sum, coinB)
   indexB += 1
  }
 }
}

func (coins Coins) Percent(percent int64) Coins {
 res := Coins{}
 for _, c := range coins {
  res = append(res, c.Percent(percent))
 }
 return res
}

func (coins Coins) Fraction(fraction float64) Coins {
 res := Coins{}
 for _, c := range coins {
  res = append(res, c.Fraction(fraction))
 }
 return res
}

func (coins Coins) Addon(point int64) Coins {
 res := Coins{}
 for _, c := range coins {
  res = append(res, c.Addon(point))
 }
 return res
}

func (coins Coins) Negative() Coins {
 res := make([]Coin, 0, len(coins))
 for _, coin := range coins {
  res = append(res, Coin{
   Name: coin.Name,
   Amount: -coin.Amount,

  })
 }
 return res
}

func (coinsA Coins) Minus(coinsB Coins) Coins {
 return coinsA.Plus(coinsB.Negative())
}

func (coinsA Coins) IsGTE(coinsB Coins) bool {
 diff := coinsA.Minus(coinsB)
 if len(diff) == 0 {
  return true
 }
 return diff.IsNonnegative()
}


func (coinsA Coins) IsGTEs(coinsB Coins) bool {

 coinsA.Sort()
 coinsB.Sort()

 diff := coinsA.Minus(coinsB)
 if len(diff) == 0 {
  return true
 }
 return diff.IsNonnegative()
}

func (coins Coins) IsZero() bool {
 return len(coins) == 0
}

func (coinsA Coins) IsEqual(coinsB Coins) bool {
 if len(coinsA) != len(coinsB) {
  return false
 }
 for i := 0; i < len(coinsA); i++ {
  if coinsA[i].Amount != coinsB[i].Amount {
   return false
  }

  if coinsA[i].Name != coinsB[i].Name {
   return false
  }

 }
 return true
}

func (coins Coins) IsPositive() bool {
 if len(coins) == 0 {
  return false
 }
 for _, coinAmount := range coins {
  if coinAmount.Amount <= 0 {
   return false
  }
 }
 return true
}

func (coins Coins) IsNonnegative() bool {
 if len(coins) == 0 {
  return true
 }
 for _, coinAmount := range coins {
  if coinAmount.Amount < 0 {
   return false
  }
 }
 return true
}



func (c Coins) Len() int { return len(c) }
func (c Coins) Less(i, j int) bool { return c[i].Name < c[j].Name }
func (c Coins) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Coins) Sort() { sort.Sort(c) }

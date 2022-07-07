package state

import (
 "fmt"
 "math/rand"
 "testing"
 "time"

)

const (
 percent = 5
 num = 15
 devtime = 10
)
# 42 "candidat_test.go"
func TestCoins(t *testing.T) {





 DepoBlk(3, 9, 7)
# 77 "candidat_test.go"
}



func DepoBlk(deposeed, timeseed, blkseed int64) {

 lasttime := time.Now().UnixNano()


 sum := int64(0)
 pw_sum := int64(0)
 min := int64(9999)

 var validator [num]Candidat


 gen_depo := rand.New(rand.NewSource(deposeed))



 for i := 0; i < num; i++ {
  amount := gen_depo.Int63n(9999) + 1
  sum = sum + amount
  validator[i].Contract = Coins{Coin{Name: "Deposit", Amount: amount}}
  validator[i].Power = amount


  if min > amount {
   min = amount
  }
 }


 fmt.Printf("Power: ")
 for i := 0; i < num; i++ {
  validator[i].Power = validator[i].Power / min
  fmt.Printf("%v, ", validator[i].Power)

  validator[i].PowerTest = validator[i].Power
  pw_sum = pw_sum + validator[i].Power
 }
 fmt.Printf("Blk payment in one cycle: %v\n", pw_sum)


 fmt.Printf("\nAll deposit: %v\n", sum)


 prc := int64((sum * percent) / 100)
 fmt.Printf("All prc(%v): %v\n", percent, prc)



 block_in_last_mth := int64(1 + rand.Intn(899))
 fmt.Printf("block_in_last_mth: %v\n", block_in_last_mth)


 cycle_mth := block_in_last_mth / pw_sum
 cycle := cycle_mth



 var priz int64
 priz_out := int64(0)


 priz_forecast := prc / block_in_last_mth
 priz = priz_forecast
 fmt.Printf("priz for one blk: %v\n", priz)
 fmt.Printf("All priz from cycl[%v]: %v\n\n", cycle, priz*block_in_last_mth)


 v := num - 1
 pw := pw_sum

 for {



  if validator[v].PowerTest > 0 {
   validator[v].PowerTest = validator[v].PowerTest - 1
   validator[v].Priz = validator[v].Priz + priz


   priz_out = priz_out + priz



   pw = pw - 1

   if pw < 1 {
    pw = pw_sum

    for i := 0; i < num; i++ {
     validator[i].PowerTest = validator[i].Power
    }
    fmt.Printf("cycle[%v]: %v\n", cycle, priz_out)
    cycle = cycle - 1
   }

   block_in_last_mth = block_in_last_mth - 1

  }


  v = v - 1
  if v < 0 {
   v = num - 1
  }



  if block_in_last_mth < 0 {
   fmt.Printf("priz_out: %v, %v\n", priz_out, pw_sum-pw)
   break
  }

  if time.Now().UnixNano()-lasttime > 5000000 {
   fmt.Printf("time out ?")
   break
  }
 }

 fmt.Printf("Priz_out: %v,  overpayment: %.2f pc \n", priz_out, float64((priz_out-prc)*100.0)/float64(prc))

 fmt.Printf("time all: %v\n", time.Now().UnixNano()-lasttime)

}

func Depo(deposeed, timeseed, blkseed int64) {


 sum := int64(0)
 min := int64(9999)

 var validator [num]Candidat


 gen_depo := rand.New(rand.NewSource(deposeed))
 gen_time := rand.New(rand.NewSource(timeseed))
 gen_blk := rand.New(rand.NewSource(blkseed))

 for i := 0; i < num; i++ {
  amount := gen_depo.Int63n(9999) + 1
  sum = sum + amount
  validator[i].Contract = Coins{Coin{Name: "Deposit", Amount: amount}}
  validator[i].Power = amount


  if min > amount {
   min = amount
  }
 }


 fmt.Printf("Power: ")
 for i := 0; i < num; i++ {
  validator[i].Power = validator[i].Power / min
  fmt.Printf("%v, ", validator[i].Power)
 }
 fmt.Printf("\n")



 fmt.Printf("\nAll deposit: %v\n", sum)


 prc := int64((sum * percent) / 100)
 fmt.Printf("All prc(%v): %v\n", percent, prc)


 var priz int64
 block_in_mth := int64(0)
 priz_out := int64(0)
# 258 "candidat_test.go"
 b_in_clk := int64(0)


 time_in_mth := int(60 * 60 * 4)
 for t := time_in_mth; t > 1; t-- {

  if b_in_clk == 0 {

   b_in_clk = int64(1 + gen_blk.Intn(18))

   block_in_mth = block_in_mth + 1



   k := t + devtime - gen_time.Intn(devtime*2)

   end_blk := int64(k) / b_in_clk


   if end_blk > 0 {
    priz = (prc - priz_out) / end_blk
   }

   if priz < 1 {
    priz = 1
   }
  }


  b_in_clk = b_in_clk - 1

  priz_out = priz_out + priz

  if priz == 0 {
   fmt.Printf("priz[%d]: %v\n", t, priz)
  }

 }

 fmt.Printf("Priz_out: %v,  overpayment: %.2f pc \n", priz_out, float64((priz_out-prc)*100.0)/float64(sum))
}

package state

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	amcLineFree     = 0
	amcLineClient   = 1
	amcLineJunior   = 2
	amcLineManager  = 3
	amcLineSenior   = 4
	amcLineTop      = 5
	amcLineGeneral  = 6
	amcLineDirector = 7
)

type LineMarketing struct {
	Levels []int64

	Packets map[int64]int
}

func (l *LineMarketing) ValidateLevel() error {
	all := int64(0)
	if len(l.Levels) > 0 {
		for _, ls := range l.Levels {
			all = all + ls
		}
	}

	if all > 100 {
		return fmt.Errorf("all pcs more 100")
	}
	return nil
}

func (l *LineMarketing) SetParam(key string, value int64) error {
	p := strings.Split(key, ".")

	if len(p) != 2 {
		return fmt.Errorf("key must have 2 parts")
	}
	switch p[0] {
	case "level":
		lvl, err := strconv.Atoi(p[1])
		if err != nil {
			return err
		}
		if lvl >= len(l.Levels) {
			n := make([]int64, lvl+1)
			copy(n, l.Levels)
			l.Levels = n
		}
		l.Levels[lvl] = value
	case "packet":
		packID, err := strconv.Atoi(p[1])
		if err != nil {
			return err
		}
		if l.Packets == nil {
			l.Packets = map[int64]int{}
		}
		l.Packets[int64(packID)] = int(value)
	default:
		return fmt.Errorf("unknown key part '%s'", p[0])
	}
	return nil
}

func (l *LineMarketing) MakePayment(refChain RefChain, delta Coins) (Coins, RefChain) {
	totalPayments := Coins{}
	event := RefChain{}

	if l == nil {
		return totalPayments, event
	}

	fmt.Printf("Line Marketing processing payment... \n")

	lvl := 0
	for i, acc := range refChain {
		if lvl >= len(l.Levels) {
			break
		}

		if i == 0 {
			lvl++
			continue
		}

		tarif := acc.Balance.Coin(AMCCoinTarifBbon)
		if tarif == nil {

			continue
		}

		tarif = acc.Balance.CoinNoExpired(AMCCoinTarifBbon)
		if tarif.Amount < 1 {
			continue
		}

		pack := acc.Balance.Coin(AMCCoinPacket)
		if pack == nil {
			continue
		}

		packLvl := l.Packets[pack.Amount]
		if packLvl < lvl {
			continue
		}
		percent := l.Levels[lvl]
		if percent <= 0 {
			lvl++
			continue
		}

		payment := delta.Percent(percent)
		acc.Balance = acc.Balance.Plus(payment)
		totalPayments = totalPayments.Plus(payment)
		lvl++

		pAcc := acc.Copy()
		pAcc.Balance = payment
		event = append(event, pAcc)
	}
	return totalPayments, event
}

func (l *LineMarketing) MakeRank(refChain RefChain, refRank RefRank, tarif Coins) (RefChain, RefChain) {

	var TimeTarif time.Time
	var ReferalPaket int64
	var ReferalAddress []byte

	ranks := RefChain{}

	ReferalPaket = 0
	ReferalAddress = []byte("ref")

	if l == nil {
		return refChain, nil
	}

	tarif299 := tarif.Coin(AMCCoinTarifBbon)
	if tarif299 == nil {

		return refChain, nil
	}

	fmt.Printf("Line Marketing TX buy/change rank... \n")

	for i, acc := range refRank {

		fmt.Printf("...Acc  %v \n", acc.Acc)

		if i == 0 {
			ref := acc.Acc.Balance.Coin(AMCCoinTarifBbon)
			if ref != nil {

				fmt.Printf("\nDub buying tarif...%v \n", acc.Acc)

				tarif := acc.Acc.Balance.CoinNoExpired(AMCCoinTarifBbon)
				if tarif.Amount > 0 {
					remainsTimeTarif := tarif.Expired.Sub(TimeWorkingBlock)
					TimeTarif = TimeWorkingBlock.Add(remainsTimeTarif)
				} else {
					TimeTarif = TimeWorkingBlock
				}

				ref = &Coin{Name: AMCCoinTarifBbon, Amount: 0, Expired: TimeTarif.AddDate(1, 0, 0)}

				acc.Acc.Balance = acc.Acc.Balance.Plus(Coins{*ref})

				pAcc := acc.Acc.Copy()
				ref.Amount = 299
				pAcc.Balance = Coins{*ref}
				ranks = append(ranks, pAcc)

				break
			} else {

				ref = &Coin{Name: AMCCoinTarifBbon, Amount: 299, Expired: TimeWorkingBlock.AddDate(1, 0, 0)}

				acc.Acc.Balance = acc.Acc.Balance.Plus(Coins{*ref})

				ref = &Coin{Name: AMCCoinPacket, Amount: amcLineClient}
				acc.Acc.Balance = acc.Acc.Balance.Plus(Coins{*ref})

				ReferalAddress = acc.Acc.Address
				ReferalPaket = ref.Amount

				pAcc := acc.Acc.Copy()
				pAcc.Balance = Coins{*acc.Acc.Balance.Coin(AMCCoinTarifBbon)}
				pAcc.Balance = pAcc.Balance.Plus(Coins{*acc.Acc.Balance.Coin(AMCCoinPacket)})
				ranks = append(ranks, pAcc)
			}
		}

		tarif := acc.Acc.Balance.Coin(AMCCoinTarifBbon)
		if tarif == nil {
			break
		}

		dwrank := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
		for _, dw := range acc.Down {
			nm := dw.Balance.Coin(AMCCoinPacket)
			if nm != nil {
				fmt.Printf("dw.Acc  %v \n", dw)
				if bytes.Compare(dw.Address, ReferalAddress) == 0 {

					dwrank[ReferalPaket] = dwrank[ReferalPaket] + 1
				} else {
					dwrank[nm.Amount] = dwrank[nm.Amount] + 1
				}
			} else {

				if bytes.Compare(dw.Address, ReferalAddress) == 0 {
					dwrank[ReferalPaket] = dwrank[ReferalPaket] + 1
				}
			}
		}

		fmt.Printf("ReferalAddress %x \n", ReferalAddress)
		fmt.Printf("ReferalPaket %v \n", ReferalPaket)
		fmt.Printf("dwrank %v \n \n", dwrank)

		pack := &Coin{Name: AMCCoinPacket, Amount: 0}
		if dwrank[amcLineClient] > 3 {
			pack = &Coin{Name: AMCCoinPacket, Amount: amcLineJunior}
		}
		if dwrank[amcLineJunior] > 1 {
			pack = &Coin{Name: AMCCoinPacket, Amount: amcLineManager}
		}
		if dwrank[amcLineManager] > 1 {
			pack = &Coin{Name: AMCCoinPacket, Amount: amcLineSenior}
		}
		if dwrank[amcLineSenior] > 1 {
			pack = &Coin{Name: AMCCoinPacket, Amount: amcLineTop}
		}
		if dwrank[amcLineTop] > 1 {
			pack = &Coin{Name: AMCCoinPacket, Amount: amcLineGeneral}
		}
		if dwrank[amcLineGeneral] > 1 {
			pack = &Coin{Name: AMCCoinPacket, Amount: amcLineDirector}
		}

		packNow := acc.Acc.Balance.Coin(AMCCoinPacket)

		if packNow.Amount < pack.Amount {
			pack_add := &Coin{Name: AMCCoinPacket, Amount: (pack.Amount - packNow.Amount)}
			acc.Acc.Balance = acc.Acc.Balance.Plus(Coins{*pack_add})

			ReferalAddress = acc.Acc.Address
			ReferalPaket = pack.Amount

			pAcc := acc.Acc.Copy()
			pAcc.Balance = Coins{*acc.Acc.Balance.Coin(AMCCoinTarifBbon)}
			pAcc.Balance = pAcc.Balance.Plus(Coins{*acc.Acc.Balance.Coin(AMCCoinPacket)})
			ranks = append(ranks, pAcc)

		}

	}

	rankUpdate := RefChain{}
	for _, acc := range refRank {
		rankUpdate = append(rankUpdate, &acc.Acc)
	}

	fmt.Printf("make rank end \n \n")
	return rankUpdate, ranks
}

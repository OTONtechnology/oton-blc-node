package state

import (
	"encoding/gob"
	"os/exec"
)

const (
	ageLoadUpdate = "loadupdate"
	ageLoadSHA    = "loadsha"

	ageNoIncentPrize = "offprize"
)

func init() {
	gob.Register(&AgeParam{})
}

type AgeParam struct {
	Key   string
	Value int64
	Text  string

	Height int64
	Used   int64
}

type AgeProcessor struct {
	AgeMap map[string]AgeParam
}

func (a *AgeProcessor) OnFix() int64 {
	p := a.AgeMap["23feb"]
	return p.Height
}

func (a *AgeProcessor) OnIncentPrize() int64 {
	p, ok := a.AgeMap[ageNoIncentPrize]

	if ok {
		return p.Value
	} else {
		return 1
	}
}

func (a *AgeProcessor) InitAge() {

	a.AgeMap[ageNoIncentPrize] = AgeParam{Key: ageNoIncentPrize, Value: 1, Text: ageNoIncentPrize}
}

func (a *AgeProcessor) UpdateAge(in []AgeParam) (out map[string]AgeParam) {
	for _, p := range in {

		old, ok := a.AgeMap[p.Key]
		if ok {
			p.Used = old.Used
		}

		p.Height = HeightLine
		a.AgeMap[p.Key] = p
		out[p.Key] = p
	}
	return out
}

func (a *AgeProcessor) WorkAge(in map[string]AgeParam) error {

	loadAge, ok := in[ageLoadUpdate]
	scriptSHA, _ := in[ageLoadSHA]

	if ok {
		if loadAge.Text != "" {

			cmd := exec.Command("chmod", "+x", "runcmd.sh")
			err := cmd.Start()
			if err != nil {
				return err
			}

			cmd = exec.Command("./runcmd.sh", loadAge.Text, scriptSHA.Text)
			err = cmd.Start()
			if err != nil {
				return err
			}

		}
	}

	return nil
}

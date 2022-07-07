package otoncoin

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"time"

	"fmt"
	"strconv"
	"strings"

	"github.com/tendermint/tendermint/abci/apps/otoncoin/state"
	"github.com/tendermint/tendermint/abci/apps/otoncoin/tx"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/version"

	cryptoenc "github.com/tendermint/tendermint/crypto/encoding"
)

const (
	appVersion = uint64(0x1)
	maxTxSize  = 10240
)

type Application struct {
	types.BaseApplication
	state *state.State

	RetainBlocks int64

	ValUpdates []types.ValidatorUpdate

	logger log.Logger
}

func NewApplication() *Application {
	return &Application{
		state: state.NewState(state.NewMemKVStore()),

		logger: log.NewNopLogger(),
	}
}

func (app *Application) SetLogger(l log.Logger) {
	app.logger = l
}

func (app *Application) Info(req types.RequestInfo) (resInfo types.ResponseInfo) {
	return types.ResponseInfo{
		Data:       fmt.Sprintf("{\"size\":%v}", req.Size()),
		Version:    version.ABCIVersion,
		AppVersion: appVersion,
	}
}

func (app *Application) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	if len(req.Tx) > maxTxSize {
		return types.ResponseDeliverTx{Code: tx.TypeErrTxSize}
	}

	var rawTx tx.Raw
	err := rawTx.Unmarshal(req.Tx)
	if err != nil {
		return types.ResponseDeliverTx{Code: tx.TypeErrTxDecode}
	}

	txExec := func(state *state.State, raw, hash []byte) func(tx.Tx) types.ResponseDeliverTx {
		return func(t tx.Tx) types.ResponseDeliverTx {
			if err := t.Unmarshal(raw); err != nil {
				return types.ResponseDeliverTx{Code: tx.TypeErrTxDecode}
			}
			return t.Exec(state, hash, false)
		}
	}(app.state, rawTx.Raw, crypto.Sha256(req.Tx))

	switch rawTx.Type {
	case tx.TypeSendCoins:
		return txExec(&tx.SendCoins{})
	case tx.TypeMintCoins:
		return txExec(&tx.MintCoins{})

	case tx.TypeTest:
		return txExec(&tx.SendCoins{})

	case tx.TypeCreateAMC:
		return txExec(&tx.CreateAMC{})
	case tx.TypeSetInAMC:
		return txExec(&tx.SetInAMC{})
	case tx.TypeBuyInAMC:
		return txExec(&tx.BuyInAMC{})

	case tx.TypeChangeWant:
		return txExec(&tx.ChangeAddressWant{})
	case tx.TypeChangeApply:
		return txExec(&tx.ChangeAddressApply{})

	case tx.TypeNewAge:
		return txExec(&tx.SetNewAge{})

	case tx.TypeNodeReward:
		return txExec(&tx.SetNodeReward{})
	case tx.TypeBuyNode:
		res := txExec(&tx.BuyNode{})
		if !res.IsErr() {

			app.updateValidator(types.UpdateValidator(res.Data, res.GasWanted, ""))
			res.Data = []byte("buynode")
			res.GasWanted = 0
		}

		return res

	default:
		return types.ResponseDeliverTx{Code: tx.TypeUnknownError}

	}
}

func (app *Application) CheckTx(req types.RequestCheckTx) types.ResponseCheckTx {

	if len(req.Tx) > maxTxSize {
		return types.ResponseCheckTx{Code: tx.TypeErrTxSize}
	}

	var rawTx tx.Raw
	err := rawTx.Unmarshal(req.Tx)
	if err != nil {
		return types.ResponseCheckTx{Code: tx.TypeErrTxDecode}
	}

	txExec := func(state *state.State, raw, hash []byte) func(tx.Tx) types.ResponseCheckTx {
		return func(t tx.Tx) types.ResponseCheckTx {
			if err := t.Unmarshal(raw); err != nil {
				return types.ResponseCheckTx{Code: tx.TypeErrTxDecode}
			}

			res := t.Exec(state, hash, true)
			return types.ResponseCheckTx{

				Code: res.Code,
				Data: res.Data,
				Log:  res.Log,
				Info: res.Info}
		}
	}(app.state, rawTx.Raw, crypto.Sha256(req.Tx))

	switch rawTx.Type {
	case tx.TypeSendCoins:
		return txExec(&tx.SendCoins{})

	case tx.TypeMintCoins:
		return txExec(&tx.MintCoins{})

	case tx.TypeTest:
		return txExec(&tx.SendCoins{})

	case tx.TypeCreateAMC:
		return txExec(&tx.CreateAMC{})
	case tx.TypeSetInAMC:
		return txExec(&tx.SetInAMC{})
	case tx.TypeBuyInAMC:
		return txExec(&tx.BuyInAMC{})

	case tx.TypeChangeWant:
		return txExec(&tx.ChangeAddressWant{})
	case tx.TypeChangeApply:
		return txExec(&tx.ChangeAddressApply{})

	case tx.TypeNewAge:
		return txExec(&tx.SetNewAge{})

	case tx.TypeNodeReward:
		return txExec(&tx.SetNodeReward{})
	case tx.TypeBuyNode:
		return txExec(&tx.BuyNode{})

	default:

		fmt.Println("CheckTx error !!!")
		return types.ResponseCheckTx{Code: tx.TypeErrCheckTx}
	}

	return types.ResponseCheckTx{
		Code: tx.TypeOK,
	}
}

func (app *Application) Commit() types.ResponseCommit {

	fmt.Println("\nCommit")

	appHash := make([]byte, 8)
	binary.PutVarint(appHash, app.state.Size)
	app.state.AppHash = appHash
	app.state.Height++

	resp := types.ResponseCommit{Data: appHash}
	if app.RetainBlocks > 0 && app.state.Height >= app.RetainBlocks {
		resp.RetainHeight = app.state.Height - app.RetainBlocks + 1
	}

	return resp
}

func (app *Application) InitChain(req types.RequestInitChain) types.ResponseInitChain {

	genDoc := new(state.GenesisDoc)
	err := json.Unmarshal(req.GetAppStateBytes(), genDoc)
	if err != nil {
		fmt.Printf("unmarshaling genesis file: %v", err)
		return types.ResponseInitChain{}
	}

	app.state.SetChainID(req.GetChainId())
	fmt.Printf("\nChainID:_%v_\n", app.state.GetChainID())

	bz, _ := json.Marshal(genDoc)
	fmt.Printf("\nGenesis: %s \n \n ", string(bz))

	for _, acc := range genDoc.Accounts {
		app.state.SetAccount(acc.PubKey.Address(), &acc)

		for _, coin := range acc.Balance {
			if err := app.state.SetCoin(&coin); err != nil {
				fmt.Printf("\nErr set genesis coin : %v \n", err)
			}
		}
	}

	app.state.SetAdrOuner(genDoc.Ouner)
	app.state.SetAdrPullFee(genDoc.PullFee)
	app.state.SetAdrMinterOuner(genDoc.Minter)

	fmt.Printf("\ngenDoc.Ouner : %v \n", genDoc.Ouner)
	fmt.Printf("genDoc.Minter : %v \n", genDoc.Minter)
	fmt.Printf("genDoc.PullFee : %v \n", genDoc.PullFee)

	fmt.Printf("\nInitchain validators \n")
	app.state.ValProc.ValMap = make(map[string]state.Candidat)

	app.state.ValProc.Gift = state.Coins{state.NewCoin("testcoin", 1)}

	app.state.Age.InitAge()

	for _, v := range req.Validators {
		r := app.updateValidator(v)
		fmt.Printf("   %v \n", v)
		if r.IsErr() {
			app.logger.Error("Error updating validators", "r", r)
		}
	}
	fmt.Printf("\n \n")

	return types.ResponseInitChain{}
}

func (app *Application) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {

	if len(reqQuery.Data) == 0 {
		resQuery.Log = "Query cannot be zero length"
		resQuery.Code = tx.TypeEncodingError
		return
	}

	if reqQuery.Path == "account" {
		acc, _ := app.state.GetAccount(reqQuery.Data)

		if acc == nil {

			acc = &state.Account{}
			resQuery.Log = "does not exist"
			resQuery.Key = reqQuery.Data

			bz, _ := json.Marshal(acc)
			resQuery.Value = bz
			resQuery.Info = string(bz)
		} else {
			resQuery.Log = "exists"
			resQuery.Key = reqQuery.Data

			bz, _ := json.Marshal(acc)
			resQuery.Value = bz
			resQuery.Info = string(bz)
		}
	}

	if reqQuery.Path == "getcoin" {
		coin, _ := app.state.GetCoin(string(reqQuery.Data))
		if coin == nil {

			coin = &state.Coin{}
			resQuery.Log = "does not exist"
			resQuery.Key = reqQuery.Data
			bz, _ := json.Marshal(coin)
			resQuery.Value = bz
			resQuery.Info = string(bz)
		} else {
			resQuery.Log = "exists"
			resQuery.Key = reqQuery.Data
			bz, _ := json.Marshal(coin)
			resQuery.Value = bz
			resQuery.Info = string(bz)
		}
	}

	if reqQuery.Path == "getamc" {
		amc, _ := app.state.GetAMC(reqQuery.Data)
		if amc == nil {

			amc = &state.AMC{}
			resQuery.Log = "does not exist"
			resQuery.Key = reqQuery.Data
			resQuery.Value = []byte(amc.String())
			resQuery.Info = amc.String()
		} else {
			resQuery.Log = "exists"
			resQuery.Key = reqQuery.Data
			resQuery.Value = []byte(amc.String())
			resQuery.Info = amc.String()
		}
	}

	if reqQuery.Path == "listcoin" {

		resQuery.Log = "this func is not work now"
		resQuery.Key = reqQuery.Data
	}

	if reqQuery.Path == "getval" {
		key := []byte("val:" + string(reqQuery.Data))
		fmt.Printf("abci_query, key val: %x \n", key)
		value := app.state.Get(key)
		if len(value) == 0 {
			resQuery.Log = "does not exist"
			resQuery.Key = reqQuery.Data
			resQuery.Value = value
		} else {
			resQuery.Log = "exists"
			resQuery.Key = reqQuery.Data

			validator := new(types.ValidatorUpdate)
			err := types.ReadMessage(bytes.NewBuffer(value), validator)
			if err != nil {
				panic(err)
			}

			pubkey, _ := cryptoenc.PubKeyFromProto(validator.PubKey)
			candidat := app.state.ValProc.ValMap[string(pubkey.Address())]

			bz, _ := json.Marshal(candidat)
			resQuery.Info = string(bz)
			resQuery.Value = bz
		}
	}

	resQuery.Height = app.state.Height
	return resQuery
}

func (app *Application) BeginBlock(req types.RequestBeginBlock) types.ResponseBeginBlock {
	var out, out2 tx.Outputs
	var allout state.Coins
	var bzs, bze string

	app.ValUpdates = make([]types.ValidatorUpdate, 0)

	for _, ev := range req.ByzantineValidators {
		if ev.Type == types.EvidenceType_DUPLICATE_VOTE {
			addr := string(ev.Validator.Address)

			fmt.Printf("ev.TotalVotingPower: %d \n \n", ev.TotalVotingPower)

			if cv, ok := app.state.ValProc.ValMap[addr]; ok {

				app.updateValidator(types.Ed25519ValidatorUpdate(cv.PubKey.Bytes(), ev.Validator.Power-1))
				app.logger.Info("Decreased val power by 1 because of the equivocation",
					"val", addr)
			} else {
				app.logger.Error("Wanted to punish val, but can't find it",
					"val", addr)
			}
		}
	}

	fmt.Printf("\n")
	fmt.Printf("Block beg %d \n", req.Header.Height)
	fmt.Printf("ProposerAddress %x \n", req.Header.ProposerAddress)
	state.TimeWorkingBlock = req.Header.Time
	state.HeightLine = req.Header.Height

	if proposer, ok := app.state.ValProc.ValMap[string(req.Header.ProposerAddress)]; ok {

		holder := proposer.PubKey.Address()
		if len(proposer.Holder) > 2 {
			holder = proposer.Holder
		}

		tc := tx.FromStateCoins(app.state.ValProc.Gift)
		ot := tx.Output{Address: holder, Coins: tc}
		out = append(out, ot)

		VotesInfo := req.LastCommitInfo
		for _, vi := range VotesInfo.GetVotes() {

			if deputat, ok := app.state.ValProc.ValMap[string(vi.Validator.Address)]; ok {

				if len(deputat.Holder) < 20 {
					deputat.Holder = deputat.PubKey.Address()
				}

				elapsedTime := req.Header.Time.Sub(deputat.TimePay)
				month, _ := time.ParseDuration("720h")

				fractionWork := float64(elapsedTime.Nanoseconds()) / float64(month.Nanoseconds())
				PrizWork := deputat.Priz.Fraction(fractionWork)

				if app.state.Age.OnIncentPrize() < 1 && !vi.SignedLastBlock {
					PrizWork = deputat.Priz.Fraction(0)
					deputat.TimePay = req.Header.Time
				}

				allout = allout.Plus(PrizWork)

				tc2 := tx.FromStateCoins(PrizWork)

				if PrizWork.IsValid() && !PrizWork.IsZero() {
					ot := tx.Output{Address: deputat.Holder, Coins: tc2}

					out2 = append(out2, ot)
					deputat.TimePay = req.Header.Time

				}

				app.state.ValProc.ValMap[string(deputat.PubKey.Address())] = deputat
			}
		}

		txc := tx.SendCoins{
			Fee:     tx.Coin{"", 0},
			Inputs:  tx.Inputs{tx.Input{Address: app.state.ValProc.PullFee, Sequence: 5, Coins: tc}},
			Outputs: out,
		}

		txc2 := tx.SendCoins{
			Fee:     tx.Coin{"", 0},
			Inputs:  tx.Inputs{tx.Input{Address: app.state.ValProc.PullFee, Sequence: 5, Coins: tx.FromStateCoins(allout)}},
			Outputs: out2,
		}

		txcM2 := tx.SendCoins{
			Fee:     tx.Coin{"", 0},
			Inputs:  tx.Inputs{tx.Input{Address: app.state.ValProc.PullFee, Sequence: 5, Coins: tx.FromStateCoins(allout)}},
			Outputs: out2.MergeBalance(),
		}

		events := []types.Event{
			{
				Type: "node_reward",
				Attributes: []types.EventAttribute{
					{Key: "proposer_address", Value: fmt.Sprintf("%x", req.Header.ProposerAddress), Index: true},
				},
			},
		}

		var res types.ResponseDeliverTx
		if req.Header.Height < tx.OnNewPrizHight {
			res = txc.Coinbase(app.state, false)
		}
		if req.Header.Height >= tx.OnNewPrizHight && req.Header.Height <= tx.OnFixDelVal {
			res = txc2.Coinbase(app.state, false)
		}
		if req.Header.Height > tx.OnFixDelVal {
			res = txcM2.Coinbase(app.state, false)
		}

		if !res.IsErr() {

			var bz_coinbase []byte
			if req.Header.Height < tx.OnNewPrizHight {
				txc.Inputs = tx.Inputs{tx.Input{Address: app.state.ValProc.PullFee, Coins: tc}}
				bz_coinbase, _ = json.Marshal(txc)
				fmt.Printf("coinbase: %s \n \n", bz_coinbase)
			} else {
				txc2.Inputs = tx.Inputs{tx.Input{Address: app.state.ValProc.PullFee, Coins: tx.FromStateCoins(allout)}}
				bz_coinbase, _ = json.Marshal(txc2)
				fmt.Printf("allout: %s \n \n", allout)
			}
			bzs = string(bz_coinbase)

			key2 := types.EventAttribute{Key: "coinbase", Value: bzs, Index: true}
			events[0].Attributes = append(events[0].Attributes, key2)
		} else {
			fmt.Printf("error coinbase: %v \n \n", res)
			bze = fmt.Sprintf("code: %v, %s", res.Code, res.Log)

			bz_coinbase, _ := json.Marshal(txc2)
			fmt.Printf("coinbase: %s \n \n", bz_coinbase)

			key2 := types.EventAttribute{Key: "coinbase_error", Value: bze, Index: true}
			events[0].Attributes = append(events[0].Attributes, key2)
		}

		return types.ResponseBeginBlock{Events: events}
	}

	return types.ResponseBeginBlock{}
}

func (app *Application) EndBlock(req types.RequestEndBlock) types.ResponseEndBlock {

	fmt.Printf("Val Updates %s \n", app.ValUpdates)

	fmt.Printf("Block end %v \n", app.ValUpdates)
	return types.ResponseEndBlock{ValidatorUpdates: app.ValUpdates}
}

func (app *Application) ListSnapshots(
	req types.RequestListSnapshots) types.ResponseListSnapshots {
	return types.ResponseListSnapshots{}
}

func (app *Application) LoadSnapshotChunk(
	req types.RequestLoadSnapshotChunk) types.ResponseLoadSnapshotChunk {
	return types.ResponseLoadSnapshotChunk{}
}

func (app *Application) OfferSnapshot(
	req types.RequestOfferSnapshot) types.ResponseOfferSnapshot {
	return types.ResponseOfferSnapshot{Result: types.ResponseOfferSnapshot_ABORT}
}

func (app *Application) ApplySnapshotChunk(
	req types.RequestApplySnapshotChunk) types.ResponseApplySnapshotChunk {
	return types.ResponseApplySnapshotChunk{Result: types.ResponseApplySnapshotChunk_ABORT}
}

func (app *Application) execValidatorTx(txv []byte) types.ResponseDeliverTx {
	txv = txv[len(ValidatorSetChangePrefix):]

	pubKeyAndPower := strings.Split(string(txv), "!")
	if len(pubKeyAndPower) != 2 {
		return types.ResponseDeliverTx{
			Code: tx.TypeEncodingError,
			Log:  fmt.Sprintf("Expected 'pubkey!power'. Got %v", pubKeyAndPower)}
	}
	pubkeyS, powerS := pubKeyAndPower[0], pubKeyAndPower[1]

	pubkey, err := base64.StdEncoding.DecodeString(pubkeyS)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: tx.TypeEncodingError,
			Log:  fmt.Sprintf("Pubkey (%s) is invalid base64", pubkeyS)}
	}

	power, err := strconv.ParseInt(powerS, 10, 64)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: tx.TypeEncodingError,
			Log:  fmt.Sprintf("Power (%s) is not an int", powerS)}
	}

	return app.updateValidator(types.UpdateValidator(pubkey, power, ""))
}

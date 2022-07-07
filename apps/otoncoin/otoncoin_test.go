package otoncoin

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/abci/client"
	"github.com/tendermint/tendermint/abci/example/code"
	abciserver "github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/libs/service"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const (
	testValue  = "def"
	testKeyHex = "0000"
)

var ctx = context.Background()

var testKey2, _ = hex.DecodeString(testKeyHex)
var testKey = string(testKey2)

func testKVStore(t *testing.T, app types.Application, tx []byte, key, value string) {
	req := types.RequestDeliverTx{Tx: tx}
	ar := app.DeliverTx(req)
	require.False(t, ar.IsErr(), ar)

	ar = app.DeliverTx(req)
	require.False(t, ar.IsErr(), ar)

	app.Commit()

	info := app.Info(types.RequestInfo{})
	require.NotZero(t, info.LastBlockHeight)

	resQuery := app.Query(types.RequestQuery{
		Path: "/store",
		Data: []byte(key),
	})
	require.Equal(t, code.CodeTypeOK, resQuery.Code)
	require.Equal(t, key, string(resQuery.Key))
	require.Equal(t, value, string(resQuery.Value))
	require.EqualValues(t, info.LastBlockHeight, resQuery.Height)

	resQuery = app.Query(types.RequestQuery{
		Path:  "/store",
		Data:  []byte(key),
		Prove: true,
	})
	require.EqualValues(t, code.CodeTypeOK, resQuery.Code)
	require.Equal(t, key, string(resQuery.Key))
	require.Equal(t, value, string(resQuery.Value))
	require.EqualValues(t, info.LastBlockHeight, resQuery.Height)
}

func TestKVStoreKV(t *testing.T) {
	kvstore := NewApplication()
	key := testKey
	value := key
	tx := []byte(key)
	testKVStore(t, kvstore, tx, key, value)

	value = testValue
	tx = []byte(key + "=" + value)
	testKVStore(t, kvstore, tx, key, value)
}

func TestQuery(t *testing.T) {
	assert := assert.New(t)

	app := NewApplication()
	key := testKey

	tx := []byte(key)

	req := types.RequestDeliverTx{Tx: tx}
	res := app.DeliverTx(req)

	assert.True(res.IsOK(), "Commit, DeliverTx: Expected OK return from DeliverTx, Error: %v", res)

	resQueryPreCommit := app.Query(types.RequestQuery{
		Path: "/account",
	})

	res2 := app.Commit()
	assert.True(res.IsOK(), res2)

	resQueryPostCommit := app.Query(types.RequestQuery{
		Path: "/account",
	})
	assert.NotEqual(resQueryPreCommit, resQueryPostCommit, "Query should change before/after commit")
}

func TestValUpdates(t *testing.T) {

	kvstore := NewApplication()

	total := 10
	nInit := 5
	vals := RandVals(total)

	kvstore.InitChain(types.RequestInitChain{
		AppStateBytes: []byte(testGenesisFmtVal),
		Validators:    vals[:nInit],
	})

	fmt.Printf("kvstore.vals %x \n \n", kvstore.Validators())

	vals1, vals2 := vals[:nInit], kvstore.Validators()

	fmt.Printf("vals1 %v \n", vals1)
	fmt.Printf("vals2 %v \n", vals2)

	valsEqual(t, vals1, vals2)

	fmt.Println("\n\nAdd some validators")

	var v1, v2, v3 types.ValidatorUpdate
	v1, v2 = vals[nInit], vals[nInit+1]
	diff := []types.ValidatorUpdate{v1, v2}
	tx1 := MakeValSetChangeTx(v1.PubKey, v1.Power)
	fmt.Printf("tx1 %s \n", tx1)
	tx2 := MakeValSetChangeTx(v2.PubKey, v2.Power)
	fmt.Printf("tx2 %s \n", tx2)

	makeApplyBlock(t, kvstore, 1, diff, tx1, tx2)

	vals1, vals2 = vals[:nInit+2], kvstore.Validators()
	valsEqual(t, vals1, vals2)

	fmt.Println("\n\nRemove some validators")
	v1, v2, v3 = vals[nInit-2], vals[nInit-1], vals[nInit]
	v1.Power = 0
	v2.Power = 0
	v3.Power = 0
	diff = []types.ValidatorUpdate{v1, v2, v3}
	tx1 = MakeValSetChangeTx(v1.PubKey, v1.Power)
	tx2 = MakeValSetChangeTx(v2.PubKey, v2.Power)
	tx3 := MakeValSetChangeTx(v3.PubKey, v3.Power)

	makeApplyBlock(t, kvstore, 2, diff, tx1, tx2, tx3)

	vals1 = append(vals[:nInit-2], vals[nInit+1])
	vals2 = kvstore.Validators()
	valsEqual(t, vals1, vals2)

	fmt.Println("\n\nUpdate some validators")
	v1 = vals[0]
	if v1.Power == 5 {
		v1.Power = 6
	} else {
		v1.Power = 5
	}
	diff = []types.ValidatorUpdate{v1}
	tx1 = MakeValSetChangeTx(v1.PubKey, v1.Power)

	makeApplyBlock(t, kvstore, 3, diff, tx1)

	vals1 = append([]types.ValidatorUpdate{v1}, vals1[1:]...)
	vals2 = kvstore.Validators()
	valsEqual(t, vals1, vals2)

}

func makeApplyBlock(
	t *testing.T,
	kvstore types.Application,
	heightInt int,
	diff []types.ValidatorUpdate,
	txs ...[]byte) {

	height := int64(heightInt)
	hash := []byte("foo")
	header := tmproto.Header{
		Height: height,
	}

	kvstore.BeginBlock(types.RequestBeginBlock{Hash: hash, Header: header})
	for _, tx := range txs {
		if r := kvstore.DeliverTx(types.RequestDeliverTx{Tx: tx}); r.IsErr() {
			t.Fatal(r)
		}
	}
	resEndBlock := kvstore.EndBlock(types.RequestEndBlock{Height: header.Height})
	kvstore.Commit()

	valsEqual(t, diff, resEndBlock.ValidatorUpdates)
}

func valsEqual(t *testing.T, vals1, vals2 []types.ValidatorUpdate) {
	if len(vals1) != len(vals2) {
		t.Fatalf("vals dont match in len. got %d, expected %d", len(vals2), len(vals1))
	}
	sort.Sort(types.ValidatorUpdates(vals1))
	sort.Sort(types.ValidatorUpdates(vals2))
	for i, v1 := range vals1 {
		v2 := vals2[i]
		if !v1.PubKey.Equal(v2.PubKey) ||
			v1.Power != v2.Power {
			t.Fatalf("vals dont match at index %d. got %X/%d , expected %X/%d", i, v2.PubKey, v2.Power, v1.PubKey, v1.Power)
		}
	}
}

func makeSocketClientServer(app types.Application, name string) (abcicli.Client, service.Service, error) {

	socket := fmt.Sprintf("unix://%s.sock", name)
	logger := log.TestingLogger()

	server := abciserver.NewSocketServer(socket, app)
	server.SetLogger(logger.With("module", "abci-server"))
	if err := server.Start(); err != nil {
		return nil, nil, err
	}

	client := abcicli.NewSocketClient(socket, false)
	client.SetLogger(logger.With("module", "abci-client"))
	if err := client.Start(); err != nil {
		if err = server.Stop(); err != nil {
			return nil, nil, err
		}
		return nil, nil, err
	}

	return client, server, nil
}

func makeGRPCClientServer(app types.Application, name string) (abcicli.Client, service.Service, error) {

	socket := fmt.Sprintf("unix://%s.sock", name)
	logger := log.TestingLogger()

	gapp := types.NewGRPCApplication(app)
	server := abciserver.NewGRPCServer(socket, gapp)
	server.SetLogger(logger.With("module", "abci-server"))
	if err := server.Start(); err != nil {
		return nil, nil, err
	}

	client := abcicli.NewGRPCClient(socket, true)
	client.SetLogger(logger.With("module", "abci-client"))
	if err := client.Start(); err != nil {
		if err := server.Stop(); err != nil {
			return nil, nil, err
		}
		return nil, nil, err
	}
	return client, server, nil
}

func runClientTests(t *testing.T, client abcicli.Client) {

	key := testKey
	value := key
	tx := []byte(key)
	testClient(t, client, tx, key, value)

	value = testValue
	tx = []byte(key + "=" + value)
	testClient(t, client, tx, key, value)
}

func testClient(t *testing.T, app abcicli.Client, tx []byte, key, value string) {
	ar, err := app.DeliverTxSync(ctx, types.RequestDeliverTx{Tx: tx})
	require.NoError(t, err)
	require.False(t, ar.IsErr(), ar)

	ar, err = app.DeliverTxSync(ctx, types.RequestDeliverTx{Tx: tx})
	require.NoError(t, err)
	require.False(t, ar.IsErr(), ar)

	_, err = app.CommitSync(ctx)
	require.NoError(t, err)

	info, err := app.InfoSync(ctx, types.RequestInfo{})
	require.NoError(t, err)
	require.NotZero(t, info.LastBlockHeight)

	resQuery, err := app.QuerySync(ctx, types.RequestQuery{
		Path: "/store",
		Data: []byte(key),
	})
	require.Nil(t, err)
	require.Equal(t, code.CodeTypeOK, resQuery.Code)
	require.Equal(t, key, string(resQuery.Key))
	require.Equal(t, value, string(resQuery.Value))
	require.EqualValues(t, info.LastBlockHeight, resQuery.Height)

	resQuery, err = app.QuerySync(ctx, types.RequestQuery{
		Path:  "/store",
		Data:  []byte(key),
		Prove: true,
	})
	require.Nil(t, err)
	require.Equal(t, code.CodeTypeOK, resQuery.Code)
	require.Equal(t, key, string(resQuery.Key))
	require.Equal(t, value, string(resQuery.Value))
	require.EqualValues(t, info.LastBlockHeight, resQuery.Height)
}

package client_test

import (
	"os"
	"testing"

	"github.com/danil-lashin/tendermint/abci/example/kvstore"
	nm "github.com/danil-lashin/tendermint/node"
	rpctest "github.com/danil-lashin/tendermint/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a tendermint node (and kvstore) in the background to test against
	app := kvstore.NewKVStoreApplication()
	node = rpctest.StartTendermint(app)
	code := m.Run()

	// and shut down proper at the end
	node.Stop()
	node.Wait()
	os.Exit(code)
}

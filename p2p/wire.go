package p2p

import (
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/danil-lashin/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}

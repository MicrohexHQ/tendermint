package types

import (
	amino "github.com/tendermint/go-amino"
	"github.com/danil-lashin/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}

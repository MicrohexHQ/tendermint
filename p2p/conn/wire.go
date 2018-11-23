package conn

import (
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/danil-lashin/tendermint/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}

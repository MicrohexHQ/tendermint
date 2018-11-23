package evidence

import (
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/danil-lashin/tendermint/crypto/encoding/amino"
	"github.com/danil-lashin/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterEvidenceMessages(cdc)
	cryptoAmino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

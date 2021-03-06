package genutil

import (
	abci "github.com/evdatsion/aphelion-dpos-bft/abci/types"

	"github.com/evdatsion/cusp-sdk/codec"
	sdk "github.com/evdatsion/cusp-sdk/types"
	"github.com/evdatsion/cusp-sdk/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, cdc *codec.Codec, stakingKeeper types.StakingKeeper,
	deliverTx deliverTxfn, genesisState GenesisState,
) []abci.ValidatorUpdate {

	var validators []abci.ValidatorUpdate
	if len(genesisState.GenTxs) > 0 {
		validators = DeliverGenTxs(ctx, cdc, genesisState.GenTxs, stakingKeeper, deliverTx)
	}

	return validators
}

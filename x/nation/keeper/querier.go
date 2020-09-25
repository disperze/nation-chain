package keeper

import (
	"github.com/disperze/nation-chain/x/nation/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for nation clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// case types.QueryParams:
		// 	return queryParams(ctx, k)
		case types.QueryGetPerson:
			return queryNames(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown nation query endpoint")
		}
	}
}

// nolint: unparam
func queryNames(ctx sdk.Context, path []string, keeper Keeper) ([]byte, error) {
	value, err := keeper.Get(ctx, path[0])

	if err != nil {
		return nil, err
	}

	if value.Name == "" {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, value)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
// 	params := k.GetParams(ctx)

// 	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
// 	}

// 	return res, nil
// }

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()

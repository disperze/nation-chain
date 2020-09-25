package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/disperze/nation-chain/x/nation/types"
)

// Keeper of the nation store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
}

// NewKeeper creates a nation keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramspace types.ParamSubspace) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) Get(ctx sdk.Context, key string) (types.Person, error) {
	store := ctx.KVStore(k.storeKey)
	var item types.Person
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// RegisterDni register new DNI
func (k Keeper) RegisterDni(ctx sdk.Context, dni string, value types.Person) error {
	if k.IsDNIPresent(ctx, dni) {
		return sdkerrors.Wrap(types.ErrExistValue, "dni already register")
	}

	k.set(ctx, dni, value)
	return nil
}

// IsDNIPresent Check if the DNI is present in the store or not
func (k Keeper) IsDNIPresent(ctx sdk.Context, dni string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(dni))
}

func (k Keeper) set(ctx sdk.Context, key string, value types.Person) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
	store.Set([]byte(key), bz)
}

func (k Keeper) delete(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

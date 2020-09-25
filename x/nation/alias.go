package nation

import (
	"github.com/disperze/nation-chain/x/nation/keeper"
	"github.com/disperze/nation-chain/x/nation/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	DefaultParamspace   = types.DefaultParamspace
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewMsgRegisterDni   = types.NewMsgRegisterDni
	NewPerson           = types.NewPerson
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState

	MsgRegisterDni = types.MsgRegisterDni
	Person         = types.Person
)

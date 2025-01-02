package strongforce

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

// Keeper is the keeper for strongforce
type Keeper struct {
	cdc          codec.Codec
	storeKey     store.StoreKey
	typeStoreKey store.StoreKey
}

// NewKeeper creates a new keeper for strongforce
func NewKeeper(cdc codec.Codec, storeKey store.StoreKey, typeStoreKey store.StoreKey) Keeper {
	return Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		typeStoreKey: typeStoreKey,
	}
}

// SetState sets the state of a contract
func (k Keeper) SetState(ctx sdk.Context, id []byte, data []byte, typeName []byte) {

	store := ctx.KVStore(k.storeKey)

	store.Set(id, data)

	typeStore := ctx.KVStore(k.typeStoreKey)
	typeStore.Set(id, typeName)

}

// GetState sets the state of a contract
func (k Keeper) GetState(ctx sdk.Context, id []byte) []byte {
	store := ctx.KVStore(k.storeKey)
	return store.Get(id)
}

// GetState sets the state of a contract
func (k Keeper) GetType(ctx sdk.Context, id []byte) []byte {
	store := ctx.KVStore(k.typeStoreKey)
	return store.Get(id)
}

// GetContractsIterator returns an iterator over all stored contracts
func (k Keeper) GetContractsStateIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

func (k Keeper) GetContractsTypeIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.typeStoreKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// // DelegateCoins implements github.com/cosmos/cosmos-sdk/blob/master/x/staking/types BankKeeper
// func (k Keeper) DelegateCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Tags, sdk.Error) {
// 	// return nil, sdk.ErrInternal("Unimplemented")
// 	return sdk.EmptyTags(), nil
// }

// // UndelegateCoins implements github.com/cosmos/cosmos-types/blob/master/x/staking/types BankKeeper
// func (k Keeper) UndelegateCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Tags, sdk.Error) {
// 	// return nil, sdk.ErrInternal("Unimplemented")
// 	return sdk.EmptyTags(), nil
// }

// AddCoins implements github.com/cosmos/cosmos-types/blob/master/x/distribution/types BankKeeper
func (k Keeper) AddCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error) {
	// return nil, sdk.ErrInternal("Unimplemented")
	return sdk.NewCoins(), nil
}

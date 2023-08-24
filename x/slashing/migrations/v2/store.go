package v2

import (
	storetypes "github.com/lightmos/lightmos-sdk/store/types"
	sdk "github.com/lightmos/lightmos-sdk/types"
	v2distribution "github.com/lightmos/lightmos-sdk/x/distribution/migrations/v2"
	v1 "github.com/lightmos/lightmos-sdk/x/slashing/migrations/v1"
)

// MigrateStore performs in-place store migrations from v0.40 to v0.43. The
// migration includes:
//
// - Change addresses to be length-prefixed.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey) error {
	store := ctx.KVStore(storeKey)
	v2distribution.MigratePrefixAddress(store, v1.ValidatorSigningInfoKeyPrefix)
	v2distribution.MigratePrefixAddressBytes(store, v1.ValidatorMissedBlockBitArrayKeyPrefix)
	v2distribution.MigratePrefixAddress(store, v1.AddrPubkeyRelationKeyPrefix)

	return nil
}

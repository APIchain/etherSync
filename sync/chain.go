package sync

import (
	"math/big"
	"github.com/ethereum/go-ethereum"
	"encoding/json"
	"context"
	"github.com/etherSync/metadata"
)

func getBlockInfo(ctx context.Context, number int64) (*metadata.RpcBlock, error) {
	var r json.RawMessage
	err := rpcclient.CallContext(ctx, &r, "eth_getBlockByNumber", toBlockNumArg(big.NewInt(number)), true)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		}
	}
	return metadata.UnMarshalBlock(r)
}
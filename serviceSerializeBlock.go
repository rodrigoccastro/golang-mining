package main

import (
	"bytes"
	"fmt"

	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

func serializeBlock(btcdBlock *wire.MsgBlock) (*btcutil.Block, error) {
	var buf bytes.Buffer
	err := btcdBlock.Serialize(&buf)
	if err != nil {
		return nil, fmt.Errorf("error serializing block: %v", err)
	}

	return btcutil.NewBlock(btcdBlock), nil
}

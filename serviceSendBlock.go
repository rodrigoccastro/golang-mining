package main

// import (
// 	"fmt"

// 	"github.com/btcsuite/btcd/rpcclient"
// 	"github.com/btcsuite/btcutil"
// )

import (
	"fmt"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
)

func sendBlock(client *rpcclient.Client, btcdBlockSerialized *btcutil.Block) error {
	err := client.SubmitBlock(btcdBlockSerialized, nil)
	if err != nil {
		return fmt.Errorf("error sending block to the Bitcoin Core node: %v", err)
	}

	return nil
}
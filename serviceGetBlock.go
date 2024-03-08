package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
)

var lastHashGetBlock *chainhash.Hash

func getBlock(client *rpcclient.Client) *wire.MsgBlock {

	for {
		bestBlockHash, err := client.GetBestBlockHash()
		if err != nil {
			fmt.Println("Error when call GetBestBlockHash: ", err)
			continue
		}
		if bestBlockHash == lastHashGetBlock {
			fmt.Println("Already fetched this block")
			continue
		}
		blockMsg, err := client.GetBlock(bestBlockHash)
		if err != nil {
			fmt.Println("Error when call GetBlock: ", err)
			continue
		}
		fmt.Println(" - bestBlockHash: ", bestBlockHash)
		lastHashGetBlock = bestBlockHash
		return blockMsg
	}

}

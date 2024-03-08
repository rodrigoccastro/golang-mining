package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

func TestMineBlock(t *testing.T) {
	unixTime := int64(1645833559)
	timestamp := time.Unix(unixTime, 0)
	prevBlockBytes := []byte{229, 136, 183, 1, 181, 153, 195, 83, 121, 139, 227, 81, 120, 48, 22, 174, 16, 38, 126, 100, 196, 17, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	merkleRootBytes := []byte{14, 220, 240, 205, 100, 183, 8, 116, 123, 110, 205, 13, 252, 130, 135, 76, 229, 10, 11, 0, 58, 140, 196, 138, 104, 62, 52, 121, 221, 202, 172, 244}

	// Convert byte slices to chainhash.Hash type
	prevBlockHash, _ := chainhash.NewHash(prevBlockBytes)
	merkleRootHash, _ := chainhash.NewHash(merkleRootBytes)

	block := &wire.MsgBlock{
		Header: wire.BlockHeader{
			Version:    567517184,
			PrevBlock:  *prevBlockHash,
			MerkleRoot: *merkleRootHash,
			Timestamp:  timestamp,
			Bits:       386101681,
			Nonce:      3192944194,
		},
		Transactions: []*wire.MsgTx{
			{
				Version: 1,
				TxIn: []*wire.TxIn{
					{
						PreviousOutPoint: wire.OutPoint{
							Hash:  chainhash.Hash{},
							Index: wire.MaxPrevOutIndex,
						},
						SignatureScript: []byte("A0CyDBsvVmlhQlRDL01pbmVkIGJ5IHRvY2hrYXMxOS8s+r5tbXRhbvZ/Ibp+xhZMnRpXnNXSquojvVpkLZzJvE/Y17EtEAAAAAAAAAAQMIXjBSIgOQ6kFMqtkKMBAAAAAAA="),
						Witness:         [][]byte{[]byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=")},
						Sequence:        wire.MaxTxInSequenceNum,
					},
				},
				TxOut: []*wire.TxOut{
					{
						Value:    644058753,
						PkScript: []byte("dqkUU2/6mSSRUI3KA1TlLzKjp6Z5pTqIrA=="),
					},
					{
						Value:    0,
						PkScript: []byte("ailSU0tCTE9DSzqFXPl5qHVw7L0TuO1N+6XgLhx4aaV2uiR6Bw0dAF1QTQ=="),
					},
					{
						Value:    0,
						PkScript: []byte("aiSqIantIbbWoL4mJp8ns0Y5slFu86lcun/A2YTjhrApXEPtZQ8="),
					},
				},
				LockTime: 0,
			},
		},
	}

	count := 0
	hash, nonce := mineBlock(block)
	if hash != "" {
		count++
		fmt.Println("count: ", count, " - nonce:", nonce, " - hash:", hash)
	}
	if count == 0 {
		t.Error("Unexpected error: nonce not founded!")
	}
}

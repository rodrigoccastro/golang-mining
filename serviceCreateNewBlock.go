package main

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

const NullValueIn = ^uint32(0)

func createNewBlock(block *wire.MsgBlock, myAddress btcutil.Address, hashFounded string, nonce uint32) (*wire.MsgBlock, error) {
	
	// Parse the hashFounded string into a chainhash.Hash
	prevHash, err := chainhash.NewHashFromStr(hashFounded)
	if err != nil {
		return nil, fmt.Errorf("error parsing previous block hash: %v", err)
	}
	// Set the PrevBlock and Nonce fields in the block header
	block.Header.PrevBlock = *prevHash
	block.Header.Nonce = nonce

	// Create a coinbase transaction
	cbTx := wire.NewMsgTx(wire.TxVersion)

	// Create a coinbase script that pushes some data onto the stack
	coinbaseScript := []byte{
		txscript.OP_1,      // Pushes the number 1 onto the stack
		txscript.OP_ADD,    // Adds the top two items on the stack
		txscript.OP_RETURN, // Marks the transaction as invalid
	}

	// Create a coinbase input
	var prevOutHash chainhash.Hash // Use zero hash for coinbase transaction
	prevOutHash[0] = 0x00          // Set the first byte to 0x00
	outpoint := wire.NewOutPoint(&prevOutHash, NullValueIn)
	cbTxIn := wire.NewTxIn(outpoint, coinbaseScript, nil)
	cbTx.AddTxIn(cbTxIn)

	// Create the coinbase output
	script, err := txscript.PayToAddrScript(myAddress)
	if err != nil {
		return nil, fmt.Errorf("error creating output script: %v", err)
	}
	txOut := wire.NewTxOut(625_000_000, script) // Reward amount in satoshis
	cbTx.AddTxOut(txOut)

	// Add the coinbase transaction to the block
	block.AddTransaction(cbTx)

	return block, nil
}
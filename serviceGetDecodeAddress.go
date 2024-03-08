package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

var myAddress = "bc1qqvwcf6x6mr6ulljarmy2j7p7s2g3s0fvlcx0nw"

func getDecodeAddress() (btcutil.Address, error) {

	// Decode the Bech32 address
	address, err := btcutil.DecodeAddress(myAddress, &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("error DecodeAddress: %v", err)
	}

	// Ensure that the address is a SegWit address
	_, ok := address.(*btcutil.AddressWitnessPubKeyHash)
	if !ok {
		return nil, fmt.Errorf("Not a SegWit address: %v", err)
	}

	// Create the SegWit P2WPKH address
	addressSegWit, err := btcutil.NewAddressWitnessPubKeyHash(address.ScriptAddress(), &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("Error in NewAddressWitnessPubKeyHash: %v", err)
	}

	return addressSegWit, nil
}

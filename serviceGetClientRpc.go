package main

import (
	"fmt"

	"github.com/btcsuite/btcd/rpcclient"
)

func getClientRpc() (*rpcclient.Client, error) {
	client, err := rpcclient.New(GetConnRpcConfig(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error when call rpcclient.New: %v", err)
	}
	return client, nil
}

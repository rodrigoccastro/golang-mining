package main

import "github.com/btcsuite/btcd/rpcclient"

// ConnConfigProperties holds the properties of rpcclient.ConnConfig
type ConnConfigProperties struct {
	Host         string
	User         string
	Pass         string
	HTTPPostMode bool
	DisableTLS   bool
}

func GetConnConfigProperties() ConnConfigProperties {
	return ConnConfigProperties{
		Host:         "127.0.0.1:8332",    // Bitcoin Core node address and port
		User:         "your_rpc_username", // Bitcoin Core node username
		Pass:         "your_rpc_password", // Bitcoin Core node password
		HTTPPostMode: true,                // Use HTTP POST mode
		DisableTLS:   true,                // Disable TLS (in local environment)
	}
}

func GetConnRpcConfig() *rpcclient.ConnConfig {
	connCfgProps := GetConnConfigProperties()
	return &rpcclient.ConnConfig{
		Host:         connCfgProps.Host,
		User:         connCfgProps.User,
		Pass:         connCfgProps.Pass,
		HTTPPostMode: connCfgProps.HTTPPostMode,
		DisableTLS:   connCfgProps.DisableTLS,
	}
}

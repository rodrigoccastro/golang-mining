package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strings"

	"github.com/btcsuite/btcd/wire"
)

func mineBlock(block *wire.MsgBlock) (string, uint32) {
	difficulty := 7
	maxIterations := 500_000
	nonce := uint32(1_572_866_488 + 25_000)
	iteration := 1

	for iteration <= maxIterations {
		block.Header.Nonce = nonce
		hashStr := getHashByBlock(block)
		if isHashValid(hashStr, difficulty) {
			return hashStr, nonce
		}
		iteration = iteration + 1
		nonce = nonce - 1

	}
	return "", math.MaxUint32
}

func doubleSHA256(data []byte) [32]byte {
	// firstHash := sha256.Sum256(data)
	// return sha256.Sum256(firstHash[:])
	return sha256.Sum256(data)
}

func getHashByBlock(block *wire.MsgBlock) string {
	var buf bytes.Buffer
	if err := block.Serialize(&buf); err != nil {
		panic(err)
	}
	bytesHash := doubleSHA256(buf.Bytes())
	return hex.EncodeToString(bytesHash[:])
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

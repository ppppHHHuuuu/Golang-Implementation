package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Version int8

const (
	CryptoCurrency Version = iota
	SmartContract
	DAPPS
	Industry
)

type Header struct {
	version   Version
	height    int
	timestamp int64
	prevHash  []byte
	hash      []byte
	nonce     int64
}

type Block struct {
	header Header
	data   []byte
}

type Hashing interface {
	Hash() []byte
}

func (block *Block) Hash() {
	timestamp := []byte(strconv.FormatInt(block.header.timestamp, 10))
	headers := bytes.Join([][]byte{block.header.prevHash, block.data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.header.hash = hash[:]
}

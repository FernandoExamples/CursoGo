package models

import (
	"math/big"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

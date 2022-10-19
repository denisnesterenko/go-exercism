package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

var seed = rand.New(rand.NewSource(time.Now().Unix()))

func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)
	privateKey := new(big.Int).Sub(p, two)
	return privateKey.Add(two, privateKey.Rand(seed, privateKey))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (privateKey, public *big.Int) {
	privateKey = PrivateKey(p)
	return privateKey, PublicKey(privateKey, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}

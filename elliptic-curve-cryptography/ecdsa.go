package elliptic_curve_cryptography

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
)

func CreateSignature(privKey *ecdsa.PrivateKey, hash []byte) (r, s *big.Int, err error) {
	sr, ss, serr := ecdsa.Sign(rand.Reader, privKey, hash)
	if serr != nil {
		err = serr
		return
	}
	return sr, ss, nil
}

func VerifySignature(pubKey *ecdsa.PublicKey, hash []byte, r, s *big.Int) bool {
	return ecdsa.Verify(pubKey, hash, r, s)
}
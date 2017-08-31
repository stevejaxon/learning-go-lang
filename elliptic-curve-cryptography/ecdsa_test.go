package elliptic_curve_cryptography

import (
	"testing"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

var privKey *ecdsa.PrivateKey
var pubKey *ecdsa.PublicKey

func generateNewKeyPair(t *testing.T, pk1 *ecdsa.PrivateKey, pk2 *ecdsa.PublicKey) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil || priv == nil {
		t.Errorf("Error generating the private key: %s", err)
		return
	}
	if pk1 != nil && pk2 != nil {
		pk1 = priv
		*pk2 = priv.PublicKey
	} else {
		privKey = priv
		pubKey = &priv.PublicKey
	}
}

func TestCanSignMessage(t *testing.T) {
	// Assign the private and public keys to our global variables
	generateNewKeyPair(t, nil, nil)
	r, s, err := CreateSignature(privKey, []byte("testing"))
	if err != nil {
		t.Errorf("Error generating the signature from the message : %s", err)
		return
	}
	fmt.Printf("The 'R' value as a string %s\n", r)
	fmt.Printf("The 'S' value as a string %s\n", s)
}

func TestCanSignMessageAndVerifySignature(t *testing.T) {
	hashed := []byte("testing")
	r, s, err := CreateSignature(privKey, hashed)
	if err != nil {
		t.Errorf("Error generating the signature from the message : %s", err)
		return
	}
	fmt.Printf("The 'R' value as a string %s\n", r)
	fmt.Printf("The 'S' value as a string %s\n", s)
	if !VerifySignature(pubKey, hashed, r, s) {
		t.Errorf("Verification of the signature against the public key failed")
	}
}

func TestCanNotVerifySignatureAgainstIncorrectHash(t *testing.T) {
	validHashed := []byte("testing")
	incorrectHashed := []byte("Testing")
	r, s, err := CreateSignature(privKey, validHashed)
	if err != nil {
		t.Errorf("Error generating the signature from the message : %s", err)
		return
	}
	if VerifySignature(pubKey, incorrectHashed, r, s) {
		t.Errorf("Verification of the signature against an incorrect hash succeeded")
	}
}

func TestSignatureCanNotBeVerifiedByWrongPublicKey(t *testing.T) {
	var privKey1 *ecdsa.PrivateKey = new(ecdsa.PrivateKey)
	var pubKey1 *ecdsa.PublicKey = new(ecdsa.PublicKey)
	generateNewKeyPair(t, privKey1, pubKey1)
	if privKey1 == nil || pubKey1 == nil {
		t.Error("Unable to generate new Public and Private keys")
		return
	}
	if privKey == privKey1 {
		t.Error("The newly generated new Private key should be different to the original one")
		return
	}
	hashed := []byte("testing")
	r, s, err := CreateSignature(privKey, hashed)
	if err != nil {
		t.Errorf("Error generating the signature from the message : %s", err)
		return
	}
	if VerifySignature(pubKey1, hashed, r, s) {
		t.Errorf("Verification of the signature against the public key failed")
	}
}
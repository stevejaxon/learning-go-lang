package elliptic_curve_cryptography

import (
	"testing"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"fmt"
)

// Private function that is used to generate the private and related public keys
func getKeys(t *testing.T) *ecdsa.PrivateKey {
	priv, err := crypto.GenerateKey()
	if err != nil || priv == nil {
		t.Errorf("Error generating the private key: %s", err)
		return nil
	}
	return priv
}

// A test to verify that the Sign function of the Ethereum library results in the desired outcome - of a digital signature (ECDSA)
func TestSignMessage(t *testing.T) {
	privKey := getKeys(t)
	plaintext := "f2a92a2013ee4ad7"
	hash := crypto.Keccak256([]byte(plaintext))
	fmt.Printf("The hash %s of the plaintext value %s\n", hash, plaintext)
	sig, err := crypto.Sign(hash, privKey)
	if err != nil || sig == nil {
		t.Errorf("Error generating the signature of the hash from the private key: %s", err)
		return
	}
	fmt.Printf("The signature %s of the plaintext value %s\n", sig, plaintext)
}

// A test to verify that given the digital signature and the original hash (that was signed with the private key)
// That it is possible to determine which public key signed the hash.
func TestVerificationOfSignature(t *testing.T) {
	privKey := getKeys(t)
	pubKey := &privKey.PublicKey
	fmt.Printf("The private key is %s and the resulting public key is %s\n", privKey, pubKey)
	plaintext := "870433f01bbee909"
	hash := crypto.Keccak256([]byte(plaintext))
	fmt.Printf("The hash %s of the plaintext value %s\n", hash, plaintext)
	sig, err := crypto.Sign(hash, privKey)
	if err != nil || sig == nil {
		t.Errorf("Error generating the signature of the hash from the private key: %s", err)
		return
	}
	fmt.Printf("The signature %s of the plaintext value %s\n", sig, plaintext)
	retrievedPubKey, err := crypto.SigToPub(hash, sig)
	if err != nil || retrievedPubKey == nil {
		t.Errorf("Error retrieving the public key from the signature and hash: %s", err)
		return
	}
	fmt.Printf("The retrieved public key is %s\n", retrievedPubKey)
	fmt.Printf("The retrieved public key X is %s\n", retrievedPubKey.X)
	fmt.Printf("The original  public key X is %s\n", pubKey.X)
	fmt.Printf("The retrieved public key Y is %s\n", retrievedPubKey.Y)
	fmt.Printf("The original  public key Y is %s\n", pubKey.Y)
	if retrievedPubKey.Curve != pubKey.Curve || retrievedPubKey.Y.Cmp(pubKey.Y) != 0 || retrievedPubKey.X.Cmp(pubKey.X) != 0 {
		t.Errorf("The public key retrieved, from the signature, %s does not match the public key, of the private key used to sign the hash, %s ", *retrievedPubKey, *pubKey)
		return
	}
}

// A test to verify that if you have a valid signature, but don't have the original hash (that was signed with the private key)
// that you are *not* able to retrieve the expected public key
func TestInvalidHashDoesNotRetrieveCorrectPublicKey(t *testing.T) {
	privKey := getKeys(t)
	pubKey := &privKey.PublicKey
	fmt.Printf("The private key is %s and the resulting public key is %s\n", privKey, pubKey)
	plaintext := "870433f01bbee909"
	invalidPlaintext := "770433f01bbee990"
	hash := crypto.Keccak256([]byte(plaintext))
	invalidHash := crypto.Keccak256([]byte(invalidPlaintext))
	fmt.Printf("The hash %s of the plaintext value %s\n", hash, plaintext)
	fmt.Printf("The hash %s of the invalid plaintext value %s\n", invalidHash, invalidPlaintext)
	sig, err := crypto.Sign(hash, privKey)
	if err != nil || sig == nil {
		t.Errorf("Error generating the signature of the hash from the private key: %s", err)
		return
	}
	fmt.Printf("The signature %s of the plaintext value %s\n", sig, plaintext)
	retrievedPubKey, err := crypto.SigToPub(invalidHash, sig)
	if err != nil || retrievedPubKey == nil {
		t.Errorf("Error retrieving the public key from the signature and hash: %s", err)
		return
	}
	fmt.Printf("The retrieved public key is %s\n", retrievedPubKey)
	fmt.Printf("The retrieved public key X is %s\n", retrievedPubKey.X)
	fmt.Printf("The original  public key X is %s\n", pubKey.X)
	fmt.Printf("The retrieved public key Y is %s\n", retrievedPubKey.Y)
	fmt.Printf("The original  public key Y is %s\n", pubKey.Y)
	if retrievedPubKey.Curve != pubKey.Curve || retrievedPubKey.Y.Cmp(pubKey.Y) == 0 || retrievedPubKey.X.Cmp(pubKey.X) == 0 {
		t.Errorf("The public key retrieved, from the signature, %s should not match the public key, of the private key used to sign the hash, %s ", *retrievedPubKey, *pubKey)
		return
	}
}

// A test to verify that if you have the original hash (that was signed with the private key), but the valid signature that you have
// didn't come from the expected private key that you are *not* able to retrieve the expected public key
func TestInvalidSignatureDoesNotRetrieveCorrectPublicKey(t *testing.T) {
	privKey1 := getKeys(t)
	privKey2 := getKeys(t)
	pubKey1 := &privKey1.PublicKey
	pubKey2 := &privKey2.PublicKey
	fmt.Printf("The first private key is %s and the resulting public key is %s\n", privKey1, pubKey1)
	fmt.Printf("The second private key is %s and the resulting public key is %s\n", privKey2, pubKey2)
	plaintext1 := "870433f01bbee909"
	plaintext2 := "670433f01bbee909"
	hash1 := crypto.Keccak256([]byte(plaintext1))
	hash2 := crypto.Keccak256([]byte(plaintext2))
	sig1, err := crypto.Sign(hash1, privKey1)
	if err != nil || sig1 == nil {
		t.Errorf("Error generating the signature of the hash from the private key: %s", err)
		return
	}
	sig2, err := crypto.Sign(hash2, privKey2)
	if err != nil || sig2 == nil {
		t.Errorf("Error generating the signature of the hash from the private key: %s", err)
		return
	}
	retrievedPubKey1, err := crypto.SigToPub(hash1, sig2)
	if err != nil || retrievedPubKey1 == nil {
		t.Errorf("Error retrieving the public key from the signature and hash: %s", err)
		return
	}
	if retrievedPubKey1.Y.Cmp(pubKey1.Y) == 0 || retrievedPubKey1.X.Cmp(pubKey1.X) == 0 ||
		retrievedPubKey1.Y.Cmp(pubKey2.Y) == 0 || retrievedPubKey1.X.Cmp(pubKey2.X) == 0 {
		t.Errorf("The public key retrieved, from the signature, %s should not match the public key, of either of the private key used to sign the hashes, %s ", *retrievedPubKey1, *pubKey)
		return
	}
	retrievedPubKey2, err := crypto.SigToPub(hash2, sig1)
	if err != nil || retrievedPubKey2 == nil {
		t.Errorf("Error retrieving the public key from the signature and hash: %s", err)
		return
	}
	if retrievedPubKey2.Y.Cmp(pubKey1.Y) == 0 || retrievedPubKey2.X.Cmp(pubKey1.X) == 0 ||
		retrievedPubKey2.Y.Cmp(pubKey2.Y) == 0 || retrievedPubKey2.X.Cmp(pubKey2.X) == 0 {
		t.Errorf("The public key retrieved, from the signature, %s should not match the public key, of either of the private key used to sign the hashes, %s ", *retrievedPubKey2, *pubKey)
		return
	}
}
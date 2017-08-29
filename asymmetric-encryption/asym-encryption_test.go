package asymmetric_encryption

import (
	"github.com/ethereum/go-ethereum/crypto"
	mrand "math/rand"
	"time"
	"testing"
)

var seed int64

// InitSingleTest should be called in the beginning of every test, which uses RNG, in order to make the tests
// reproducible independent of their sequence.
func InitSingleTest() {
	seed = time.Now().Unix()
	mrand.Seed(seed)
}


func TestCanSealAndUnsealAsymmetrically(t *testing.T) {
	InitSingleTest()
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed GenerateKey with seed %d: %s.", seed, err)
	}
	pubKey := &privateKey.PublicKey
	if pubKey == nil {
		t.Fatalf("failed to generate a Public Key from the private key %d: %s.", privateKey, err)
	}
	Seal(pubKey)
}

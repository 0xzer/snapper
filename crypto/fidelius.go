package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/0xzer/snapper/data/methods"
	"github.com/0xzer/snapper/protos"
)

// (initializewebkey payload, ecdsa pubkey, ecdsa privkey, error)
func NewFideliusTentativeWebKey() (*protos.FideliusTentativeWebKey, []byte, []byte, error) {
	privKey, pubKey, uncompressedPubKey, err := generateECDSAKeys()
	if err != nil {
		return nil, nil, nil, err
	}

	hmac_sha256_key, keyPairId, err := generateKeyPair(privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	return &protos.FideliusTentativeWebKey{
		UncompressedPubkey: uncompressedPubKey,
		KeyPairId: keyPairId,
		Rwk: hmac_sha256_key,
		Version: 9,
	}, pubKey, privKey, nil
}

func LoadPublicKeyFromBase64(publicKey string) (*protos.FideliusTentativeWebKey, error) {
	privKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKIXPublicKey(privKeyBytes)
	if err != nil {
		return nil, err
	}

	pubKey, ok := key.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid base64 encoded ecdsa publickey string")
	}

	uncompressedKey := convertToUncompressed(pubKey)
	log.Println(key)
	log.Println(uncompressedKey)
	os.Exit(1)

	return nil, nil
}

func generateKeyPair(privKey []byte) ([]byte, []byte, error) {
	hmac_sha256_key, err := methods.GenerateKey(16)
	if err != nil {
		return nil, nil, err
	}
	return hmac_sha256_key, signKeyPairId(hmac_sha256_key, privKey), nil
}

// Generate HMAC using key and ecdsa privkey (keypairId)
func signKeyPairId(key, privKey []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(privKey)
	return mac.Sum(nil)
}

func generateECDSAKeys() ([]byte, []byte, []byte, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, nil, err
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, nil, nil, err
	}

	pubBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, nil, nil, err
	}

	uncompressedPubKey := convertToUncompressed(&priv.PublicKey)

	return pubBytes, privBytes, uncompressedPubKey, nil
}

func convertToUncompressed(pubKey *ecdsa.PublicKey) []byte {
	header := byte(4) // POINT HEADER
	xBytes := pubKey.X.Bytes()
	yBytes := pubKey.Y.Bytes()

	uncompressedKey := append([]byte{header}, append(xBytes, yBytes...)...)
	return uncompressedKey
}
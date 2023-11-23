package rsautil

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func RSAPrivateKeyFromString(privateKeyStr string) (*rsa.PrivateKey, error) {
	var (
		err error
	)

	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		keyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
		if err != nil {
			return nil, err
		}
		private, err := x509.ParsePKCS8PrivateKey(keyBytes)
		if err != nil {
			return nil, err
		}

		switch private := private.(type) {
		case *rsa.PrivateKey:
			return private, nil
		default:
			return nil, errors.New("unknown type of private key")
		}
	}

	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}

	return x509.ParsePKCS1PrivateKey(b)
}

func RSAPKCS1Sign(unsignedContent []byte, key *rsa.PrivateKey) (string, error) {
	var (
		hash = sha256.New()
	)

	hash.Write(unsignedContent)
	digest := hash.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, key, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(s), nil
}

func BuildSign(mac string, privateKey string) (string, error) {
	private, err := RSAPrivateKeyFromString(privateKey)
	if err != nil {
		return "", err
	}
	sig, err := RSAPKCS1Sign([]byte(mac), private)
	if err != nil {
		return "", err
	}

	return sig, nil
}

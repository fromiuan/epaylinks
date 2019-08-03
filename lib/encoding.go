package lib

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"pkcs12"
)

type Encoding struct {
	pfxData  []byte
	password string
}

func NewEncding(pfxData []byte, password string) *Encoding {
	return &Encoding{
		pfxData:  pfxData,
		password: password,
	}
}

// sha256 sign
func (e *Encoding) Sign(byts []byte) (sign string, err error) {
	privateKey, _, err := e.decode()
	if err != nil {
		return "", err
	}

	rng := rand.Reader
	hashed := sha256.Sum256(byts)
	signature, err := rsa.SignPKCS1v15(rng, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return e.base64Encoding(signature), nil
}

// verify
func (e *Encoding) Verify(sign, body string) error {
	_, certificate, err := e.decode()
	if err != nil {
		return err
	}
	byts, err := e.base64Decode(sign)
	if err != nil {
		return err
	}

	public, ok := certificate.PublicKey.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("interface to rsa.publickey error")
	}

	hashed := sha256.Sum256([]byte(body))
	err = rsa.VerifyPKCS1v15(public, crypto.SHA256, hashed[:], byts)
	return err
}

// get certificate
func (e *Encoding) GetCertificate() string {
	_, certificate, err := e.decode()
	if err != nil {
		return ""
	}
	return certificate.SerialNumber.String()
}

// decode pfx(if array for key return first one)
func (e *Encoding) decode() (privateKey *rsa.PrivateKey, certificate *x509.Certificate, err error) {
	var private interface{}
	private, certificate, err = pkcs12.Decode(e.pfxData, e.password)
	if err != nil {
		privates, certificates, err := pkcs12.DecodeAll(e.pfxData, e.password)
		if err != nil {
			return nil, certificate, err
		}
		if len(privates) > 0 && len(certificates) > 0 {
			private = privates[0]
			certificate = certificates[0]
		} else {
			return nil, nil, fmt.Errorf("privateKey and certificate not gt 2")
		}
	}
	if privateVal, ok := private.(*rsa.PrivateKey); ok {
		privateKey = privateVal
	} else {
		return nil, nil, fmt.Errorf("parse rsa.PrivateKey error")
	}
	return privateKey, certificate, nil
}

// pem pfx return first block
func (e *Encoding) Pem() (privateKey *rsa.PrivateKey, err error) {
	block, err := pkcs12.ToPEM(e.pfxData, e.password)
	if err != nil {
		return nil, err
	}

	for i, v := range block {
		private, err := x509.ParsePKCS1PrivateKey(v.Bytes)
		if err == nil && i == 0 {
			return private, nil
		}
	}
	return nil, fmt.Errorf("parse pem error")
}

func (e *Encoding) SignPKCS1v15(src, key []byte, hash crypto.Hash) ([]byte, error) {
	hashed := sha256.Sum256(src)

	var err error
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("private key error")
	}

	var pri *rsa.PrivateKey
	pri, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, pri, hash, hashed[:])
}

func (e *Encoding) base64Encoding(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func (e *Encoding) base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

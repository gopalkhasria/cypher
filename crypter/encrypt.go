package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
)

func Encrypt(filename string, key []byte) error {

	var err error
	plaintext, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	err = ioutil.WriteFile(filename, ciphertext, 0777)
	if err != nil {
		log.Panic(err)
	}
	return err
}

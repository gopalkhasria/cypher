package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
)

func Decrypt(filename string, key []byte) error {
	var err error
	ciphertext, err := ioutil.ReadFile(filename)
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
	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Panic(err)
		return err
	}

	err = ioutil.WriteFile(filename, plaintext, 0777)
	if err != nil {
		log.Panic(err)
		return err
	}
	return err
}

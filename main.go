package main

import (
	"cyper/crypter"
	"cyper/pwd"

	"github.com/gen2brain/dlgs"
)

func main() {
	/*_, err := dlgs.Info("Info", "Lorem ipsum dolor sit amet.")
	if err != nil {
		panic(err)
	}*/
	azione, _, err := dlgs.List("Cosa devo fare?", "Seleziona l'azione", []string{"Cripta", "Decripta"})
	if err != nil {
		panic(err)
	}
	if azione == "" {
		return
	}
	names, _, err := dlgs.FileMulti("Select file", "")
	if err != nil {
		panic(err)
	}
	if len(names) == 0 {
		return
	}
	pwd.GetPass()
	if azione == "Cripta" {
		for i := range names {
			err := crypter.Encrypt(names[i], []byte(pwd.Pass))
			if err != nil {
				panic(err)
			}
		}
	} else {
		for i := range names {
			err := crypter.Decrypt(names[i], []byte(pwd.Pass))
			if err != nil {
				panic(err)
			}
		}
	}
}

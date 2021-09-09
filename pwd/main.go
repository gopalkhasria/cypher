package pwd

import (
	"log"
	"strings"

	"github.com/gen2brain/dlgs"
)

var Pass string

func GetPass() {
	var pwd string
	var err error
	var succ bool

	for len(pwd) != 10 {
		pwd, succ, err = dlgs.Password("Password", "10 lettere/numeri")
		if err != nil {
			log.Println(err)
			panic(err)
		}
		if !succ {
			return
		}
	}
	fullPwd := []string{pwd, pwd, pwd}
	Pass = strings.Join(fullPwd, "-")
}

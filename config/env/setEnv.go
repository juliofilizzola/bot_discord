package env

import (
	"log"
	"os"
)

func SetEnvTerminal() {
	err := os.Setenv("GIN_MODE", SetEnv)
	if err != nil {
		log.Fatal(err)
	}

}

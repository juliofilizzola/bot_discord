package env

import "os"

func SetEnvTerminal() {
	os.Setenv("GIN_MODE", SetEnv)
}

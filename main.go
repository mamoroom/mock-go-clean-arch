package main

import (
	"fmt"

	"github.com/mamoroom/mock-go-clean-arch/config"
	"rsc.io/quote"
)

var envCfg config.EnvConfig

func main() {
	envCfg = config.NewEnvConfig()
	fmt.Println(quote.Hello())
	fmt.Println(envCfg.GetDevEnv())
}

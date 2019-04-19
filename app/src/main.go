package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mamoroom/mock-go-clean-arch/src/config"
)

var envCfg config.EnvConfig

func main() {
	envCfg = config.NewEnvConfig()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": envCfg.GetDevEnv(),
		})
	})
	r.Run(":80")
}

package main

import (
	"fmt"

	"blackhole/pkg/logger"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

type Config struct {
	Port       int    `required:"true" env:"BH_PORT"        yaml:"port"        json:"port"`
	DomainName string `required:"true" env:"BH_DOMAIN_NAME" yaml:"domain_name" json:"domain_name"`
	Email      string `required:"true" env:"BH_EMAIL"       yaml:"email"       json:"email"`
}

func loadConf() *Config {
	cfg := Config{}
	err := configor.
		New(&configor.Config{}).
		Load(&cfg)
	if err != nil {
		err := configor.Load(&cfg, "conf.yaml")
		if err != nil {
			logger.Get().Fatal().Err(err).Msg("Cannot load conf from env or file.")
		}
	}
	return &cfg
}

func main() {
	cfg := loadConf()
	r := gin.New()
	r.Use(
		requestid.New(),
		gin.Recovery(),
		logger.GinRequestLogHandler(),
	)
	r.Any("/*any", func(c *gin.Context) {
		body, err := c.GetRawData()
		if err != nil {
			logger.WithContext(c.Request.Context()).
				Error().
				Err(err).
				Msg("error while getting raw body")
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		logger.WithContext(c.Request.Context()).Info().Bytes("body", body).Msg("request body")
		c.JSON(204, gin.H{})
	})
	logger.Get().Printf("%#v\n", cfg)
	r.Run(fmt.Sprintf(":%d", cfg.Port))
}

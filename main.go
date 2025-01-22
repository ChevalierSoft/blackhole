package main

import (
	"fmt"

	"blackhole/pkg/logger"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

func main() {
	cfg := loadConf()
	gin.SetMode(gin.ReleaseMode)
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

type Config struct {
	Port int `required:"true" env:"BH_PORT" yaml:"port" json:"port"`
}

func (c *Config) Default() {
	c.Port = 80
}

func loadConf() *Config {
	cfg := Config{}
	err := configor.New(&configor.Config{}).Load(&cfg)
	if err != nil {
		err := configor.New(&configor.Config{Silent: true}).Load(&cfg, "conf.yaml")
		if err != nil {
			cfg.Default()
		}
	}
	return &cfg
}

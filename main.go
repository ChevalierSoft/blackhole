package main

import (
	"blackhole/pkg/logger"

	"github.com/caddyserver/certmagic"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

var (
	PORT int
)

type Conf struct {
	Port       int      `required:"true" env:"PORT"        yaml:"port"        json:"port"`
	DomainName []string `required:"true" env:"DOMAIN_NAME" yaml:"domain_name" json:"domain_name"`
}

func loadConf() *Conf {
	conf := Conf{}
	err := configor.Load(&conf, "config.yml")
	if err != nil {
		logger.Get().Fatal().Err(err).Msg("Cannot load conf")
	}
	return &conf
}

func main() {

	conf := loadConf()
	r := gin.New()
	r.Use(
		requestid.New(),
		gin.Recovery(),
		logger.GinRequestLogHandler(),
	)
	r.Any("/*any", func(c *gin.Context) {
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		logger.WithContext(c.Request.Context()).Info().Bytes("body", body).Msg("request body")
		c.JSON(204, gin.H{})
	})
	certmagic.HTTPSPort = conf.Port
	err := certmagic.HTTPS(conf.DomainName, r.Handler())
	logger.Get().Fatal().Err(err).Msg("server closed")
}

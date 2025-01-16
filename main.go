package main

import (
	"blackhole/pkg/logger"

	"github.com/caddyserver/certmagic"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

type Conf struct {
	Port       int    `required:"true" env:"PORT"        yaml:"port"        json:"port"`
	DomainName string `required:"true" env:"DOMAIN_NAME" yaml:"domain_name" json:"domain_name"`
	Email      string `required:"true" env:"EMAIL"       yaml:"email"       json:"email"`
}

func loadConf() *Conf {
	conf := Conf{}
	err := configor.Load(&conf, "conf.yaml")
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
	setCertMagic(conf)
	err := certmagic.HTTPS([]string{conf.DomainName}, r.Handler())
	logger.Get().Fatal().Err(err).Msg("server closed")
}

func setCertMagic(conf *Conf) {
	certmagic.HTTPSPort = conf.Port
	certmagic.DefaultACME.Email = conf.Email
	certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA
	certmagic.DefaultACME.Agreed = true
}

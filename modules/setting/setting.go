package setting

import (
	"os"
	"time"

	"github.com/leminhson2398/zipper/pkg/logger"
	"gopkg.in/ini.v1"
)

var (
	// security
	SecretSalt                 string
	DefaultTokenExpireDuration time.Duration
	Port                       int
)

func init() {
	logger.Logger.Info().Msg("Loading configurations")

	configPath := "conf/conf.ini"
	file, err := ini.Load(configPath)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		os.Exit(1)
	}

	sec := file.Section("jwt-token")
	if sec == nil {
		logger.Logger.Error().Msg("section jwt-token not found")
		os.Exit(1)
	}

	SecretSalt = sec.Key("salt").String()
	if SecretSalt != "" {
		logger.Logger.Info().Msg("Successfully assigned value to salt")
	} else {
		logger.Logger.Error().Msg("key named salt not found")
		os.Exit(1)
	}

	serverSec := file.Section("server")
	if sec == nil {
		logger.Logger.Error().Msg("section server not found")
		os.Exit(1)
	}
	Port = serverSec.Key("port").MustInt(8000)
	logger.Logger.Info().Msg("Successfully assigned value to port")

	DefaultTokenExpireDuration = sec.Key("expire_duration").MustDuration(3 * time.Hour)
}

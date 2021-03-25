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
		logger.Logger.Debug().Msg("section jwt-token not found")
		os.Exit(1)
	}

	salt := sec.Key("salt").String()
	if salt != "" {
		SecretSalt = salt
		logger.Logger.Info().Msg("Successfully assigned value to salt")
	} else {
		logger.Logger.Debug().Msg("key named salt not found")
	}

	dur := sec.Key("expire_duration").MustDuration(3 * time.Hour)
	DefaultTokenExpireDuration = dur
}

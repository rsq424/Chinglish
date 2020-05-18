package global

import (
	"Chinglish/config"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_LOG    *oplogging.Logger
	GVA_VP     *viper.Viper
)
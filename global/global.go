package global

import (
	"Chinglish/config"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	GVA_CONFIG config.Server
	GVA_LOG    *oplogging.Logger
	GVA_VP     *viper.Viper
)
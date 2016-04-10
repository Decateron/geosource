package types

import (
	"github.com/joshheinrichs/geosource/server/config"
	"github.com/joshheinrichs/geosource/server/types/fields"
)

var typesConfig *config.Config

func Init(config *config.Config) {
	typesConfig = config
	fields.Init(config)
}

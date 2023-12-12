package fastly

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type fastlyConfig struct {
	APIKey         *string `hcl:"api_key"`
	BaseURL        *string `hcl:"base_url"`
	ServiceID      *string `hcl:"service_id"`
	ServiceVersion *string `hcl:"service_version"`
}

func ConfigInstance() interface{} {
	return &fastlyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) fastlyConfig {
	if connection == nil || connection.Config == nil {
		return fastlyConfig{}
	}
	config, _ := connection.Config.(fastlyConfig)
	return config
}

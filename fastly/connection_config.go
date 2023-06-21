package fastly

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type fastlyConfig struct {
	APIKey         *string `cty:"api_key"`
	BaseURL        *string `cty:"base_url"`
	ServiceID      *string `cty:"service_id"`
	ServiceVersion *string `cty:"service_version"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"base_url": {
		Type: schema.TypeString,
	},
	"service_id": {
		Type: schema.TypeString,
	},
	"service_version": {
		Type: schema.TypeString,
	},
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

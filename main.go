package main

import (
	"github.com/turbot/steampipe-plugin-fastly/fastly"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: fastly.Plugin})
}

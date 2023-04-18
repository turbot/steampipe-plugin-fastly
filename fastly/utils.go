package fastly

import (
	"context"
	"os"
	"strings"

	"github.com/fastly/go-fastly/v8/fastly"
	"github.com/pkg/errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type serviceClient struct {
	Client    *fastly.Client
	ServiceID string
}

func connect(ctx context.Context, d *plugin.QueryData) (*fastly.Client, string, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "fastly"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		c := cachedData.(serviceClient)
		return c.Client, c.ServiceID, nil
	}

	// Default to the env var settings
	apiKey := os.Getenv("FASTLY_API_KEY")
	serviceID := os.Getenv("FASTLY_SERVICE_ID")

	// Prefer config settings
	fastlyConfig := GetConfig(d.Connection)
	if fastlyConfig.APIKey != nil {
		apiKey = *fastlyConfig.APIKey
	}
	if fastlyConfig.ServiceID != nil {
		serviceID = *fastlyConfig.ServiceID
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, serviceID, errors.New("'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
	}
	if serviceID == "" {
		return nil, serviceID, errors.New("'service_id' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
	}

	conn, err := fastly.NewClient(apiKey)
	if err != nil {
		return nil, serviceID, err
	}

	sc := serviceClient{
		Client:    conn,
		ServiceID: serviceID,
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, sc)

	return conn, serviceID, nil
}

func isNotFoundError(err error) bool {
	return strings.HasPrefix(err.Error(), "404 - Not Found")
}

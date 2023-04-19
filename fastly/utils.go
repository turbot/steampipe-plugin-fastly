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

func connect(ctx context.Context, d *plugin.QueryData) (*serviceClient, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "fastly"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*serviceClient), nil
	}

	// Default to the env var settings
	apiKey := os.Getenv("FASTLY_API_KEY")
	serviceID := os.Getenv("FASTLY_SERVICE_ID")
	baseURL := os.Getenv("FASTLY_API_URL")

	// Prefer config settings
	fastlyConfig := GetConfig(d.Connection)
	if fastlyConfig.APIKey != nil {
		apiKey = *fastlyConfig.APIKey
	}
	if fastlyConfig.ServiceID != nil {
		serviceID = *fastlyConfig.ServiceID
	}
	if fastlyConfig.BaseURL != nil {
		baseURL = *fastlyConfig.BaseURL
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, errors.New("'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
	}
	if serviceID == "" {
		return nil, errors.New("'service_id' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
	}

	sc := &serviceClient{}

	if baseURL != "" {
		conn, err := fastly.NewClientForEndpoint(apiKey, baseURL)
		if err != nil {
			return nil, err
		}
		sc.Client = conn
		sc.ServiceID = serviceID
	} else {
		conn, err := fastly.NewClient(apiKey)
		if err != nil {
			return nil, err
		}
		sc.Client = conn
		sc.ServiceID = serviceID
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, sc)

	return sc, nil
}

func isNotFoundError(err error) bool {
	return strings.HasPrefix(err.Error(), "404 - Not Found")
}

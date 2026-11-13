package fastly

import (
	"context"
	"os"
	"strconv"

	"github.com/fastly/go-fastly/v8/fastly"
	"github.com/pkg/errors"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v6/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type serviceClient struct {
	Client         *fastly.Client
	ServiceID      string
	ServiceVersion int
}

func connect(ctx context.Context, d *plugin.QueryData) (*serviceClient, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "fastly"
	if cachedData, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		return cachedData.(*serviceClient), nil
	}

	// Default to the env var settings
	apiKey := os.Getenv("FASTLY_API_KEY")
	baseURL := os.Getenv("FASTLY_API_URL")
	serviceID := os.Getenv("FASTLY_SERVICE_ID")
	serviceVersion := os.Getenv("FASTLY_SERVICE_VERSION")

	// Prefer config settings
	fastlyConfig := GetConfig(d.Connection)
	if fastlyConfig.APIKey != nil {
		apiKey = *fastlyConfig.APIKey
	}
	if fastlyConfig.BaseURL != nil {
		baseURL = *fastlyConfig.BaseURL
	}
	if fastlyConfig.ServiceID != nil {
		serviceID = *fastlyConfig.ServiceID
	}
	if fastlyConfig.ServiceVersion != nil {
		serviceVersion = *fastlyConfig.ServiceVersion
	}

	// Error if the minimum config is not set
	if apiKey == "" || serviceID == "" {
		return nil, errors.New("'api_key' and 'service_id' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
	}

	sv, _ := strconv.Atoi(serviceVersion)

	sc := &serviceClient{
		ServiceID:      serviceID,
		ServiceVersion: sv,
	}

	// check if a base URL is provided in config
	if baseURL != "" {
		conn, err := fastly.NewClientForEndpoint(apiKey, baseURL)
		if err != nil {
			return nil, err
		}
		sc.Client = conn
	} else {
		conn, err := fastly.NewClient(apiKey)
		if err != nil {
			return nil, err
		}
		sc.Client = conn
	}

	// set active version if version is not provided in config
	if serviceVersion == "" {
		version, err := getActiveVersion(ctx, sc.Client, serviceID, d)
		if err != nil {
			return nil, err
		}
		sc.ServiceVersion = *version

		// set the latest version if there is no active version available for the service
		if sc.ServiceVersion == 0 {
			latestVersion, err := getLatestVersion(ctx, sc.Client, serviceID, d)
			if err != nil {
				return nil, err
			}
			sc.ServiceVersion = latestVersion.Number
		}
	}

	// Save to cache
	_ = d.ConnectionCache.Set(ctx, cacheKey, sc)

	return sc, nil
}

func getActiveVersion(ctx context.Context, client *fastly.Client, serviceID string, d *plugin.QueryData) (*int, error) {
	cacheKey := serviceID + "activeVersion"
	if cachedData, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		return cachedData.(*int), nil
	}

	service, err := client.GetService(&fastly.GetServiceInput{ID: serviceID})
	if err != nil {
		return nil, err
	}

	// Save to cache
	_ = d.ConnectionCache.Set(ctx, cacheKey, &service.ActiveVersion)

	return types.Int(service.ActiveVersion), nil
}

func getLatestVersion(ctx context.Context, client *fastly.Client, serviceID string, d *plugin.QueryData) (*fastly.Version, error) {
	cacheKey := serviceID + "latestVersion"
	if cachedData, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		return cachedData.(*fastly.Version), nil
	}

	version, err := client.LatestVersion(&fastly.LatestVersionInput{ServiceID: serviceID})
	if err != nil {
		return nil, err
	}

	// Save to cache
	_ = d.ConnectionCache.Set(ctx, cacheKey, version)

	return version, nil
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getServiceIdMemoized = plugin.HydrateFunc(getServiceIdUncached).Memoize(memoize.WithCacheKeyFunction(getServiceIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getServiceId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getServiceIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getServiceIdCacheKey.
func getServiceIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getServiceId"
	return key, nil
}

func getServiceIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	serviceClient, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("getServiceId", err)
		return nil, err
	}
	return serviceClient.ServiceID, nil
}

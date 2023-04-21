package fastly

import (
	"context"
	"os"
	"path"
	"strconv"

	"github.com/fastly/go-fastly/v8/fastly"
	"github.com/pkg/errors"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type serviceClient struct {
	Client         *fastly.Client
	ServiceID      string
	ServiceVersion int
}

func connect(ctx context.Context, d *plugin.QueryData) (*serviceClient, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "fastly"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
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

	sc := &serviceClient{}

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
	sc.ServiceID = serviceID

	// set active/latest version if version is not provided in config
	if serviceVersion == "" {
		version, err := getActiveVersion(sc.Client, apiKey, serviceID, d)
		if err != nil {
			return nil, err
		}
		// set the latest version if there is no active version available for the service
		if *version == 0 {
			latestVersion, err := getLatestVersion(sc.Client, apiKey, serviceID, d)
			if err != nil {
				return nil, err
			}
			sc.ServiceVersion = latestVersion.Number
		} else { // set the active version
			sc.ServiceVersion = *version
		}
	} else { // set the version provided in config
		sc.ServiceVersion, _ = strconv.Atoi(serviceVersion)
	}
	plugin.Logger(ctx).Error("version", sc.ServiceVersion)
	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, sc)

	return sc, nil
}

func getActiveVersion(client *fastly.Client, apiKey string, serviceID string, d *plugin.QueryData) (*int, error) {
	cacheKey := serviceID + "activeVersion"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*int), nil
	}

	service, err := client.GetService(&fastly.GetServiceInput{ID: serviceID})
	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, &service.ActiveVersion)

	return types.Int(service.ActiveVersion), nil
}

func getLatestVersion(client *fastly.Client, apiKey string, serviceID string, d *plugin.QueryData) (*fastly.Version, error) {
	cacheKey := serviceID + "latestVersion"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*fastly.Version), nil
	}

	version, err := client.LatestVersion(&fastly.LatestVersionInput{ServiceID: serviceID})
	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, version)

	return version, nil
}

// shouldIgnoreErrors:: function which returns an ErrorPredicate for Aiven API calls
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range notFoundErrors {
			// handle not found error
			if ok, _ := path.Match(pattern, "404"); ok {
				return true
			}
		}
		return false
	}
}
